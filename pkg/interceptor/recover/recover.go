/*
Implement interceptor for recovering after any panic
*/
package recover

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerRecoverInterceptor - Catch panic and return panic-value as error
func UnaryServerRecoverInterceptor(
    ctx context.Context, 
    req any, 
    info *grpc.UnaryServerInfo, 
    handler grpc.UnaryHandler,
) (response any, err error) {
    defer func() {
        if e := recover(); e != nil {
            err = status.Error(codes.Internal, "Internal Error")
            log.Println(info.FullMethod, "Recovered with:", e)
        }
    }()

    response, err = handler(ctx, req)
    return
}

func UnaryClientRecoverInterceptor(
    ctx context.Context, 
    method string, 
    req, reply any, 
    cc *grpc.ClientConn, 
    invoker grpc.UnaryInvoker, 
    opts ...grpc.CallOption,
) (err error) {
    defer func() {
        if e := recover(); e != nil {
            err = fmt.Errorf("[%s]: Recovered with: '%v'", method, e)
        }
    }()

    err = invoker(ctx, method, req, reply, cc, opts...)
    return
}
