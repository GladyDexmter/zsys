// +build ignore

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

type client struct {
	OrigClient    string
	LogClient     string
	PrivateClient string
}

type server struct {
	OrigServer string
	LogServer  string
}

type call struct {
	Service    string
	OrigServer string
	LogServer  string
	OrigStream string
	LogStream  string
	ReqType    string
	Name       string
}

type streamwriter struct {
	Stream      string
	MessageType string
	MessageResponse
}

type pbData struct {
	Source  string
	Package string
	Clients []client
	Servers []server
	Calls   []call
	Writers []streamwriter
}

type MessageResponse struct {
	OneOfField string
	LogField   string
}

func main() {
	var args []string
	for _, arg := range os.Args {
		if arg == "--" {
			continue
		}
		args = append(args, arg)
	}

	if len(args) != 2 {
		log.Fatalf("expected one argument: go file generated by protoc")
	}

	// Get current directory for finding template
	_, filename, _, _ := runtime.Caller(0)
	curDir := filepath.Dir(filename)
	t, err := template.ParseFiles(filepath.Join(curDir, "streamlogger.template"))
	if err != nil {
		log.Fatalf("couldn't parse template file: %v", err)
	}

	src := args[1]
	dest := fmt.Sprintf("%s.streamlogger.go", strings.TrimSuffix(src, ".pb.go"))

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, src, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("coudln't parse file: %v", err)
	}

	// Package name
	data := pbData{
		Source:  filepath.Base(src),
		Package: f.Name.Name,
	}

	// Key is the messageType
	messageLogs := make(map[string]MessageResponse)

	// We are only interested in top level declarations.
	for _, n := range f.Decls {

		switch m := n.(type) {
		case *ast.GenDecl:
			for _, o := range m.Specs {
				e, ok := o.(*ast.TypeSpec)
				if !ok {
					continue
				}

				switch elem := e.Type.(type) {

				// Messages type responses which has a oneof with a _Log field
				case *ast.StructType:
					messageTypeName := e.Name.Name

					var oneOfField string
					var logField string
					for _, field := range elem.Fields.List {
						if field.Tag == nil || field.Doc == nil || !strings.Contains(field.Tag.Value, "protobuf_oneof") {
							continue
						}

						for _, comment := range field.Doc.List {
							logFieldCandidate := fmt.Sprintf("%s_Log", messageTypeName)
							if strings.HasSuffix(comment.Text, logFieldCandidate) {
								logField = logFieldCandidate
								break
							}
						}

						oneOfField = field.Names[0].Name
						break
					}

					if oneOfField == "" || logField == "" {
						continue
					}

					messageLogs[messageTypeName] = MessageResponse{
						OneOfField: oneOfField,
						LogField:   logField,
					}

				// Server and Client declarations
				case *ast.InterfaceType:
					// Clients
					name := e.Name.Name
					if strings.HasSuffix(name, "Client") {
						// We are not interested in generating call specific call
						if strings.Contains(name, "_") {
							continue
						}

						data.Clients = append(data.Clients, client{
							OrigClient:    name,
							LogClient:     fmt.Sprintf("%sLogClient", strings.TrimSuffix(name, "Client")),
							PrivateClient: firstLetterLowercase(name),
						})
						continue
					}

					// Only look for interface type Server, not per call server
					if !strings.HasSuffix(name, "Server") || strings.Contains(name, "_") {
						continue
					}

					// Connection servers
					data.Servers = append(data.Servers, server{
						OrigServer: name,
						LogServer:  fmt.Sprintf("%sLogServer", strings.TrimSuffix(name, "Server")),
					})

					for _, callFunc := range elem.Methods.List {

						function := callFunc.Names[0].Name

						fields := callFunc.Type.(*ast.FuncType).Params.List
						reqType := fields[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
						origStream := fields[1].Type.(*ast.Ident).Name

						// Per function call server
						r := strings.SplitN(origStream, "_", 2)
						service := r[0]

						data.Calls = append(data.Calls, call{
							Service:    service,
							OrigServer: fmt.Sprintf("%sServer", service),
							LogServer:  fmt.Sprintf("%sLogServer", strings.TrimSuffix(service, "Server")),
							OrigStream: origStream,
							LogStream:  fmt.Sprintf("%s%sLogStream", firstLetterLowercase(service), function),
							ReqType:    reqType,
							Name:       function,
						})
					}
				}
			}

		// Send messages
		case *ast.FuncDecl:
			if m.Name.Name != "Send" {
				continue
			}

			stream := m.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name
			messageType := m.Type.Params.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name

			messageDetails, ok := messageLogs[messageType]
			if !ok {
				log.Fatalf("Found %s.Send(*%s) found, but no %s type declaration beforehand", stream, messageType, messageType)
			}

			data.Writers = append(data.Writers,
				streamwriter{
					Stream:          stream,
					MessageType:     messageType,
					MessageResponse: messageDetails,
				})

		}
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatalf("couldn't execute template: %v", err)
	}

	// Reformat the source (for number of \t in struct to align fields)
	res, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("couldn't format buffered template: %v", err)
	}

	tmpName := dest + ".new"
	out, err := os.Create(tmpName)
	if err != nil {
		log.Fatalf("couldn't create file: %v", err)
	}

	if _, err := out.Write(res); err != nil {
		out.Close()
		os.Remove(tmpName)
		log.Fatalf("couldn't format and save in file: %v", err)
	}
	out.Close()

	if err := os.Rename(tmpName, dest); err != nil {
		log.Fatalf("couldn't rename temporary file to new name: %v", err)
	}
}

func firstLetterLowercase(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
