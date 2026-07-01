// Package server provides HTTP and gRPC with auth middleware
//
// Package server 提供带认证中间件的 HTTP 和 gRPC 服务
package server

import (
	"log/slog"

	"github.com/go-kratos/kratos/v3/middleware"
	"github.com/go-kratos/kratos/v3/middleware/recovery"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/yylego/kratos-auth/authkratos"
	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
	"github.com/yylego/kratos-static-auth/statickratosauth"
	"github.com/yylego/must"
)

func NewHTTPServer(c *conf.Server, student *service.StudentService, logger *slog.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			NewRoleMiddleware(c, logger), // Role-based auth from config // 基于配置的角色认证
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Address != "" {
		opts = append(opts, http.Address(c.Http.Address))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	pb.RegisterStudentServiceHTTPServer(srv, student)
	return srv
}

// Requires Authorization header with role token from config file
// Token values: admin or guest role token
//
// 需要提供 Authorization 请求头，使用配置文件中的角色令牌
// 令牌值：admin 或 guest 角色令牌
/*
curl --location 'http://127.0.0.1:8001/v1/students' --header 'Authorization: 63a16b29e5bc4a28a880de1b2e707cc6'
curl --location 'http://127.0.0.1:8001/v1/students' --header 'Authorization: 863676f1118c45c7add65b4adefd94dd'
*/

// NewRoleMiddleware creates auth middleware with token validation and route scope
// Configure which routes need auth and setup valid tokens
//
// NewRoleMiddleware 创建认证中间件，进行令牌验证和路由范围控制
// 配置需要认证的路由并设置有效令牌
func NewRoleMiddleware(c *conf.Server, logger *slog.Logger) middleware.Middleware {
	routeScope := authkratos.NewInclude( // Create INCLUDE mode route scope // 创建 INCLUDE 模式的路由范围
		pb.OperationStudentServiceCreateStudent,
		pb.OperationStudentServiceUpdateStudent,
		pb.OperationStudentServiceDeleteStudent,
		pb.OperationStudentServiceGetStudent,
		pb.OperationStudentServiceListStudents,
	)
	authTokens := map[string]string{ // Setup valid tokens map // 设置有效令牌映射表
		"admin": must.Nice(c.Auth.AdminToken),
		"guest": must.Nice(c.Auth.GuestToken),
	}
	authConfig := statickratosauth.NewConfig(routeScope, authTokens).
		WithFieldName("Authorization").
		WithSimpleEnable(). // Enable simple token type // 启用简单令牌类型
		WithDebugMode(true) // Enable debug mode to log auth process // 启用调试模式记录认证过程
	return statickratosauth.NewMiddleware(authConfig, logger)
}
