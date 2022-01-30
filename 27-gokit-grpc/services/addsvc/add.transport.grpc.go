/*
 Copyright 2022 Google LLC

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/
package addsvc

import (
	"context"
	"errors"
	"log"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	pb "github.com/kevinnguyenai/golang-tutorials/gokitgrpc/pb"
	stdopentracing "github.com/opentracing/opentracing-go"
	stdzipkin "github.com/openzipkin/zipkin-go"
	"google.golang.org/grpc"
)

type grpcServer struct {
	sum    grpctransport.Handler
	concat grpctransport.Handler
}

// NewGRPCServer make a set of endpoints available as a GRPC AddServer
func NewGRPCServer() {}

// NewGRPCClient returns an AddService backed by a gRPC server at the other end
// of the conn. The caller is responsilbe for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func NewGRPCClient(conn *grpc.ClientConn, otTracer stdopentracing.Tracer, zipkinTracer *stdzipkin.Tracer, logger log.Logger) {

}

func encodeGRPCSumRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(SumRequest)
	return &pb.SumRequest{A: int64(req.A), B: int64(req.B)}, nil

}

func encodeGRPCConcatRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(ConcatRequest)
	return &pb.ConcatRequest{A: req.A, B: req.B}, nil
}

// These annoying helper functions are required to translate Go error types to
// and from strings, which is the type we use in our IDLs to present errors.
// There is special casing to treat empty strings as nil errors

func str2err(s string) error {
	if s == "" {
		return nil
	}
	return errors.New(s)
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
