// Code generated by streamlogger/generator. DO NOT EDIT.
// source: zsys.pb.go

package zsys

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/ubuntu/zsys/internal/i18n"
	"github.com/ubuntu/zsys/internal/streamlogger"
	"google.golang.org/grpc"
)

/*
  Clients generated code
*/

// ZsysLogClient is a grpc ZsysClient (*zsysClient), augmented by a connexion id in context.
type ZsysLogClient struct {
	ZsysClient
	Ctx context.Context
}

// newZsysClientWithLogs returns a ZsysLogClient, which can send logs at level "level", attached to the given context
func newZsysClientWithLogs(ctx context.Context, cc *grpc.ClientConn, level logrus.Level) *ZsysLogClient {
	return &ZsysLogClient{
		ZsysClient: NewZsysClient(cc),
		Ctx:        streamlogger.NewClientCtx(ctx, level),
	}
}

// Close tear downs the connection under the ZsysLogClient.
func (z *ZsysLogClient) Close() error {
	return z.ZsysClient.(*zsysClient).cc.Close()
}

/*
   Servers
*/

// ZsysLogServer is used to intercept the server and inserting an intermediate log stream.
// This can't be done in interceptor as the creation of each per-function server struct from a grpc.ServerStream
// is only done in the handler() call, which is blocking until the whole handler has ran, once the stream
// has closed.
// It also wraps an idle timeout server, so that the interceptor can stop the idling timeout and reset it after each call.
type ZsysLogServer struct {
	ZsysServerIdleTimeout
}

type ZsysServerIdleTimeout interface {
	ZsysServer
	TrackRequest() func()
}

// registerZsysServerIdleWithLogs wraps the server to an idle timeout server and logged variant intercepting all grpc calls.
func registerZsysServerIdleWithLogs(s *grpc.Server, srv ZsysServerIdleTimeout) {
	RegisterZsysServer(s, &ZsysLogServer{srv})
}

/*
 * Zsys.Version()
 */

// zsysVersionLogStream is a Zsys_VersionServer augmented by its own Context containing the log streamer
type zsysVersionLogStream struct {
	Zsys_VersionServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysVersionLogStream) Context() context.Context {
	return s.ctx
}

// Version overrides ZsysServer Version, installing a logger first
func (z *ZsysLogServer) Version(req *Empty, stream Zsys_VersionServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "Version")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.Version(req, &zsysVersionLogStream{
		Zsys_VersionServer: stream,
		ctx:                ctx,
	})
}

/*
 * Zsys.CreateUserData()
 */

// zsysCreateUserDataLogStream is a Zsys_CreateUserDataServer augmented by its own Context containing the log streamer
type zsysCreateUserDataLogStream struct {
	Zsys_CreateUserDataServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysCreateUserDataLogStream) Context() context.Context {
	return s.ctx
}

// CreateUserData overrides ZsysServer CreateUserData, installing a logger first
func (z *ZsysLogServer) CreateUserData(req *CreateUserDataRequest, stream Zsys_CreateUserDataServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "CreateUserData")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.CreateUserData(req, &zsysCreateUserDataLogStream{
		Zsys_CreateUserDataServer: stream,
		ctx:                       ctx,
	})
}

/*
 * Zsys.ChangeHomeOnUserData()
 */

// zsysChangeHomeOnUserDataLogStream is a Zsys_ChangeHomeOnUserDataServer augmented by its own Context containing the log streamer
type zsysChangeHomeOnUserDataLogStream struct {
	Zsys_ChangeHomeOnUserDataServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysChangeHomeOnUserDataLogStream) Context() context.Context {
	return s.ctx
}

// ChangeHomeOnUserData overrides ZsysServer ChangeHomeOnUserData, installing a logger first
func (z *ZsysLogServer) ChangeHomeOnUserData(req *ChangeHomeOnUserDataRequest, stream Zsys_ChangeHomeOnUserDataServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "ChangeHomeOnUserData")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.ChangeHomeOnUserData(req, &zsysChangeHomeOnUserDataLogStream{
		Zsys_ChangeHomeOnUserDataServer: stream,
		ctx:                             ctx,
	})
}

/*
 * Zsys.PrepareBoot()
 */

// zsysPrepareBootLogStream is a Zsys_PrepareBootServer augmented by its own Context containing the log streamer
type zsysPrepareBootLogStream struct {
	Zsys_PrepareBootServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysPrepareBootLogStream) Context() context.Context {
	return s.ctx
}

// PrepareBoot overrides ZsysServer PrepareBoot, installing a logger first
func (z *ZsysLogServer) PrepareBoot(req *Empty, stream Zsys_PrepareBootServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "PrepareBoot")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.PrepareBoot(req, &zsysPrepareBootLogStream{
		Zsys_PrepareBootServer: stream,
		ctx:                    ctx,
	})
}

/*
 * Zsys.CommitBoot()
 */

// zsysCommitBootLogStream is a Zsys_CommitBootServer augmented by its own Context containing the log streamer
type zsysCommitBootLogStream struct {
	Zsys_CommitBootServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysCommitBootLogStream) Context() context.Context {
	return s.ctx
}

// CommitBoot overrides ZsysServer CommitBoot, installing a logger first
func (z *ZsysLogServer) CommitBoot(req *Empty, stream Zsys_CommitBootServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "CommitBoot")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.CommitBoot(req, &zsysCommitBootLogStream{
		Zsys_CommitBootServer: stream,
		ctx:                   ctx,
	})
}

/*
 * Zsys.SaveSystemState()
 */

// zsysSaveSystemStateLogStream is a Zsys_SaveSystemStateServer augmented by its own Context containing the log streamer
type zsysSaveSystemStateLogStream struct {
	Zsys_SaveSystemStateServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysSaveSystemStateLogStream) Context() context.Context {
	return s.ctx
}

// SaveSystemState overrides ZsysServer SaveSystemState, installing a logger first
func (z *ZsysLogServer) SaveSystemState(req *SaveSystemStateRequest, stream Zsys_SaveSystemStateServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "SaveSystemState")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.SaveSystemState(req, &zsysSaveSystemStateLogStream{
		Zsys_SaveSystemStateServer: stream,
		ctx:                        ctx,
	})
}

/*
 * Zsys.SaveUserState()
 */

// zsysSaveUserStateLogStream is a Zsys_SaveUserStateServer augmented by its own Context containing the log streamer
type zsysSaveUserStateLogStream struct {
	Zsys_SaveUserStateServer
	ctx context.Context
}

// Context access the log streamer context
func (s *zsysSaveUserStateLogStream) Context() context.Context {
	return s.ctx
}

// SaveUserState overrides ZsysServer SaveUserState, installing a logger first
func (z *ZsysLogServer) SaveUserState(req *SaveUserStateRequest, stream Zsys_SaveUserStateServer) error {
	// it's ok to panic in the assertion as we expect to have generated above the Write() function.
	ctx, err := streamlogger.AddLogger(stream.(streamlogger.StreamLogger), "SaveUserState")
	if err != nil {
		return fmt.Errorf(i18n.G("couldn't attach a logger to request: %w"), err)
	}

	// wrap the context to access the context with logger
	return z.ZsysServerIdleTimeout.SaveUserState(req, &zsysSaveUserStateLogStream{
		Zsys_SaveUserStateServer: stream,
		ctx:                      ctx,
	})
}

/*
 * Extend streams to io.Writer
 */

// Write promote zsysVersionServer to an io.Writer
func (s *zsysVersionServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&VersionResponse{
			Reply: &VersionResponse_Log{Log: string(p)},
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// Write promote zsysCreateUserDataServer to an io.Writer
func (s *zsysCreateUserDataServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&LogResponse{
			Log: string(p),
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// Write promote zsysChangeHomeOnUserDataServer to an io.Writer
func (s *zsysChangeHomeOnUserDataServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&LogResponse{
			Log: string(p),
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// Write promote zsysPrepareBootServer to an io.Writer
func (s *zsysPrepareBootServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&PrepareBootResponse{
			Reply: &PrepareBootResponse_Log{Log: string(p)},
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// Write promote zsysCommitBootServer to an io.Writer
func (s *zsysCommitBootServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&CommitBootResponse{
			Reply: &CommitBootResponse_Log{Log: string(p)},
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// Write promote zsysSaveSystemStateServer to an io.Writer
func (s *zsysSaveSystemStateServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&CreateSaveStateResponse{
			Reply: &CreateSaveStateResponse_Log{Log: string(p)},
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}

// Write promote zsysSaveUserStateServer to an io.Writer
func (s *zsysSaveUserStateServer) Write(p []byte) (n int, err error) {
	err = s.Send(
		&CreateSaveStateResponse{
			Reply: &CreateSaveStateResponse_Log{Log: string(p)},
		})
	if err != nil {
		return 0, err
	}

	return len(p), nil
}
