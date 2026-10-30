package server

import (
	"fmt"

	"buf.build/go/protovalidate"
	pb "github.com/BradleyZhang/ecommerce-review-o/api/gen/go/operation/v1"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func NewGRPCServer(srv pb.OperationServiceServer) (*grpc.Server, error) {
	var opts []grpc.ServerOption
	validator, err := protovalidate.New()
	if err != nil {
		return nil, fmt.Errorf("validator 初始化失败: %w", err)
	}
	interceptor := protovalidate_middleware.UnaryServerInterceptor(validator) // 仅能校验非流式的
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	grpcServer := grpc.NewServer(opts...)

	healthcheck := health.NewServer() // healthcheck 服务器 用于 consul
	healthpb.RegisterHealthServer(grpcServer, healthcheck)

	pb.RegisterOperationServiceServer(grpcServer, srv)
	return grpcServer, nil
}
