package proxy

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// limit simultaneous requests
type limit struct {
	current chan struct{}
	timeout time.Duration
}

func newLimit(max int, timeout time.Duration) *limit {
	l := &limit{
		current: make(chan struct{}, max),
		timeout: timeout,
	}

	for i := 0; i < max; i++ {
		l.current <- struct{}{}
	}
	return l
}

func (l *limit) acquire(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, l.timeout)
	defer cancel()

	select {
	case <-l.current:
		return nil
	case <-ctx.Done():
		if err := ctx.Err(); err != nil {
			return err
		}
		return errors.New("context done")
	}
}

func (l *limit) release() {
	select {
	case l.current <- struct{}{}:
	default:
	}
}
func (l *limit) streamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	if err := l.acquire(ss.Context()); err != nil {
		return status.Errorf(codes.DeadlineExceeded, "request limit failed")
	}
	defer l.release()
	return handler(srv, ss)
}
