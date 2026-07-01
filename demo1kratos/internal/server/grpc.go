// Package server provides HTTP and gRPC with auth middleware
//
// Package server 提供带认证中间件的 HTTP 和 gRPC 服务
package server

import (
	"log/slog"

	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/grpc"
	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
)

// NewGRPCServer creates gRPC with role-based auth middleware
// Uses same NewRoleMiddleware as HTTP to keep auth consistent
//
// NewGRPCServer 创建带角色认证中间件的 gRPC 服务器
// 与 HTTP 服务器共享相同的 NewRoleMiddleware 以保持认证一致性
func NewGRPCServer(c *conf.Server, student *service.StudentService, logger *slog.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			NewRoleMiddleware(c, logger), // Role-based auth from config // 基于配置的角色认证
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Address != "" {
		opts = append(opts, grpc.Address(c.Grpc.Address))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	pb.RegisterStudentServiceServer(srv, student)
	return srv
}
