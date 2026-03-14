# Changes

Code differences compared to source project.

## cmd/demo1kratos/wire_gen.go (+1 -1)

```diff
@@ -24,7 +24,7 @@
 		return nil, nil, err
 	}
 	studentUsecase := biz.NewStudentUsecase(dataData, logger)
-	studentService := service.NewStudentService(studentUsecase)
+	studentService := service.NewStudentService(studentUsecase, logger)
 	grpcServer := server.NewGRPCServer(confServer, studentService, logger)
 	httpServer := server.NewHTTPServer(confServer, studentService, logger)
 	app := newApp(logger, grpcServer, httpServer)
```

## configs/config.yaml (+3 -0)

```diff
@@ -5,6 +5,9 @@
   grpc:
     address: 0.0.0.0:9001
     timeout: 1s
+  auth:
+    adminToken: "63a16b29e5bc4a28a880de1b2e707cc6"
+    guestToken: "863676f1118c45c7add65b4adefd94dd"
 data:
   database:
     driver: sqlite3
```

## internal/conf/conf.pb.go (+87 -17)

```diff
@@ -78,6 +78,7 @@
 	state         protoimpl.MessageState `protogen:"open.v1"`
 	Http          *Server_HTTP           `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
 	Grpc          *Server_GRPC           `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
+	Auth          *Server_Auth           `protobuf:"bytes,3,opt,name=auth,proto3" json:"auth,omitempty"`
 	unknownFields protoimpl.UnknownFields
 	sizeCache     protoimpl.SizeCache
 }
@@ -126,6 +127,13 @@
 	return nil
 }
 
+func (x *Server) GetAuth() *Server_Auth {
+	if x != nil {
+		return x.Auth
+	}
+	return nil
+}
+
 type Data struct {
 	state         protoimpl.MessageState `protogen:"open.v1"`
 	Database      *Data_Database         `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
@@ -290,6 +298,58 @@
 	return nil
 }
 
+type Server_Auth struct {
+	state         protoimpl.MessageState `protogen:"open.v1"`
+	AdminToken    string                 `protobuf:"bytes,1,opt,name=adminToken,proto3" json:"adminToken,omitempty"`
+	GuestToken    string                 `protobuf:"bytes,2,opt,name=guestToken,proto3" json:"guestToken,omitempty"`
+	unknownFields protoimpl.UnknownFields
+	sizeCache     protoimpl.SizeCache
+}
+
+func (x *Server_Auth) Reset() {
+	*x = Server_Auth{}
+	mi := &file_conf_conf_proto_msgTypes[5]
+	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
+	ms.StoreMessageInfo(mi)
+}
+
+func (x *Server_Auth) String() string {
+	return protoimpl.X.MessageStringOf(x)
+}
+
+func (*Server_Auth) ProtoMessage() {}
+
+func (x *Server_Auth) ProtoReflect() protoreflect.Message {
+	mi := &file_conf_conf_proto_msgTypes[5]
+	if x != nil {
+		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
+		if ms.LoadMessageInfo() == nil {
+			ms.StoreMessageInfo(mi)
+		}
+		return ms
+	}
+	return mi.MessageOf(x)
+}
+
+// Deprecated: Use Server_Auth.ProtoReflect.Descriptor instead.
+func (*Server_Auth) Descriptor() ([]byte, []int) {
+	return file_conf_conf_proto_rawDescGZIP(), []int{1, 2}
+}
+
+func (x *Server_Auth) GetAdminToken() string {
+	if x != nil {
+		return x.AdminToken
+	}
+	return ""
+}
+
+func (x *Server_Auth) GetGuestToken() string {
+	if x != nil {
+		return x.GuestToken
+	}
+	return ""
+}
+
 type Data_Database struct {
 	state         protoimpl.MessageState `protogen:"open.v1"`
 	Driver        string                 `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
@@ -300,7 +360,7 @@
 
 func (x *Data_Database) Reset() {
 	*x = Data_Database{}
-	mi := &file_conf_conf_proto_msgTypes[5]
+	mi := &file_conf_conf_proto_msgTypes[6]
 	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
 	ms.StoreMessageInfo(mi)
 }
@@ -312,7 +372,7 @@
 func (*Data_Database) ProtoMessage() {}
 
 func (x *Data_Database) ProtoReflect() protoreflect.Message {
-	mi := &file_conf_conf_proto_msgTypes[5]
+	mi := &file_conf_conf_proto_msgTypes[6]
 	if x != nil {
 		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
 		if ms.LoadMessageInfo() == nil {
@@ -350,10 +410,11 @@
 	"kratos.api\x1a\x1egoogle/protobuf/duration.proto\"]\n" +
 	"\tBootstrap\x12*\n" +
 	"\x06server\x18\x01 \x01(\v2\x12.kratos.api.ServerR\x06server\x12$\n" +
-	"\x04data\x18\x02 \x01(\v2\x10.kratos.api.DataR\x04data\"\xc4\x02\n" +
+	"\x04data\x18\x02 \x01(\v2\x10.kratos.api.DataR\x04data\"\xb9\x03\n" +
 	"\x06Server\x12+\n" +
 	"\x04http\x18\x01 \x01(\v2\x17.kratos.api.Server.HTTPR\x04http\x12+\n" +
-	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x1ao\n" +
+	"\x04grpc\x18\x02 \x01(\v2\x17.kratos.api.Server.GRPCR\x04grpc\x12+\n" +
+	"\x04auth\x18\x03 \x01(\v2\x17.kratos.api.Server.AuthR\x04auth\x1ao\n" +
 	"\x04HTTP\x12\x18\n" +
 	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x18\n" +
 	"\aaddress\x18\x02 \x01(\tR\aaddress\x123\n" +
@@ -361,7 +422,14 @@
 	"\x04GRPC\x12\x18\n" +
 	"\anetwork\x18\x01 \x01(\tR\anetwork\x12\x18\n" +
 	"\aaddress\x18\x02 \x01(\tR\aaddress\x123\n" +
-	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\"y\n" +
+	"\atimeout\x18\x03 \x01(\v2\x19.google.protobuf.DurationR\atimeout\x1aF\n" +
+	"\x04Auth\x12\x1e\n" +
+	"\n" +
+	"adminToken\x18\x01 \x01(\tR\n" +
+	"adminToken\x12\x1e\n" +
+	"\n" +
+	"guestToken\x18\x02 \x01(\tR\n" +
+	"guestToken\"y\n" +
 	"\x04Data\x125\n" +
 	"\bdatabase\x18\x01 \x01(\v2\x19.kratos.api.Data.DatabaseR\bdatabase\x1a:\n" +
 	"\bDatabase\x12\x16\n" +
@@ -380,29 +448,31 @@
 	return file_conf_conf_proto_rawDescData
 }
 
-var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
+var file_conf_conf_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
 var file_conf_conf_proto_goTypes = []any{
 	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
 	(*Server)(nil),              // 1: kratos.api.Server
 	(*Data)(nil),                // 2: kratos.api.Data
 	(*Server_HTTP)(nil),         // 3: kratos.api.Server.HTTP
 	(*Server_GRPC)(nil),         // 4: kratos.api.Server.GRPC
-	(*Data_Database)(nil),       // 5: kratos.api.Data.Database
-	(*durationpb.Duration)(nil), // 6: google.protobuf.Duration
+	(*Server_Auth)(nil),         // 5: kratos.api.Server.Auth
+	(*Data_Database)(nil),       // 6: kratos.api.Data.Database
+	(*durationpb.Duration)(nil), // 7: google.protobuf.Duration
 }
 var file_conf_conf_proto_depIdxs = []int32{
 	1, // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
 	2, // 1: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
 	3, // 2: kratos.api.Server.http:type_name -> kratos.api.Server.HTTP
 	4, // 3: kratos.api.Server.grpc:type_name -> kratos.api.Server.GRPC
-	5, // 4: kratos.api.Data.database:type_name -> kratos.api.Data.Database
-	6, // 5: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
-	6, // 6: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
-	7, // [7:7] is the sub-list for method output_type
-	7, // [7:7] is the sub-list for method input_type
-	7, // [7:7] is the sub-list for extension type_name
-	7, // [7:7] is the sub-list for extension extendee
-	0, // [0:7] is the sub-list for field type_name
+	5, // 4: kratos.api.Server.auth:type_name -> kratos.api.Server.Auth
+	6, // 5: kratos.api.Data.database:type_name -> kratos.api.Data.Database
+	7, // 6: kratos.api.Server.HTTP.timeout:type_name -> google.protobuf.Duration
+	7, // 7: kratos.api.Server.GRPC.timeout:type_name -> google.protobuf.Duration
+	8, // [8:8] is the sub-list for method output_type
+	8, // [8:8] is the sub-list for method input_type
+	8, // [8:8] is the sub-list for extension type_name
+	8, // [8:8] is the sub-list for extension extendee
+	0, // [0:8] is the sub-list for field type_name
 }
 
 func init() { file_conf_conf_proto_init() }
@@ -416,7 +486,7 @@
 			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
 			RawDescriptor: unsafe.Slice(unsafe.StringData(file_conf_conf_proto_rawDesc), len(file_conf_conf_proto_rawDesc)),
 			NumEnums:      0,
-			NumMessages:   6,
+			NumMessages:   7,
 			NumExtensions: 0,
 			NumServices:   0,
 		},
```

## internal/conf/conf.proto (+5 -0)

```diff
@@ -21,8 +21,13 @@
     string address = 2;
     google.protobuf.Duration timeout = 3;
   }
+  message Auth {
+    string adminToken = 1;
+    string guestToken = 2;
+  }
   HTTP http = 1;
   GRPC grpc = 2;
+  Auth auth = 3;
 }
 
 message Data {
```

## internal/server/grpc.go (+9 -0)

```diff
@@ -1,3 +1,6 @@
+// Package server provides HTTP and gRPC with auth middleware
+//
+// Package server 提供带认证中间件的 HTTP 和 gRPC 服务
 package server
 
 import (
@@ -9,10 +12,16 @@
 	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
 )
 
+// NewGRPCServer creates gRPC with role-based auth middleware
+// Uses same NewRoleMiddleware as HTTP to keep auth consistent
+//
+// NewGRPCServer 创建带角色认证中间件的 gRPC 服务器
+// 与 HTTP 服务器共享相同的 NewRoleMiddleware 以保持认证一致性
 func NewGRPCServer(c *conf.Server, student *service.StudentService, logger log.Logger) *grpc.Server {
 	var opts = []grpc.ServerOption{
 		grpc.Middleware(
 			recovery.Recovery(),
+			NewRoleMiddleware(c, logger), // Role-based auth from config // 基于配置的角色认证
 		),
 	}
 	if c.Grpc.Network != "" {
```

## internal/server/http.go (+42 -0)

```diff
@@ -1,18 +1,26 @@
+// Package server provides HTTP and gRPC with auth middleware
+//
+// Package server 提供带认证中间件的 HTTP 和 gRPC 服务
 package server
 
 import (
 	"github.com/go-kratos/kratos/v2/log"
+	"github.com/go-kratos/kratos/v2/middleware"
 	"github.com/go-kratos/kratos/v2/middleware/recovery"
 	"github.com/go-kratos/kratos/v2/transport/http"
+	"github.com/yylego/kratos-auth/authkratos"
 	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
 	"github.com/yylego/kratos-examples/demo1kratos/internal/conf"
 	"github.com/yylego/kratos-examples/demo1kratos/internal/service"
+	"github.com/yylego/kratos-static-auth/statickratosauth"
+	"github.com/yylego/must"
 )
 
 func NewHTTPServer(c *conf.Server, student *service.StudentService, logger log.Logger) *http.Server {
 	var opts = []http.ServerOption{
 		http.Middleware(
 			recovery.Recovery(),
+			NewRoleMiddleware(c, logger), // Role-based auth from config // 基于配置的角色认证
 		),
 	}
 	if c.Http.Network != "" {
@@ -27,4 +35,38 @@
 	srv := http.NewServer(opts...)
 	pb.RegisterStudentServiceHTTPServer(srv, student)
 	return srv
+}
+
+// Requires Authorization header with role token from config file
+// Token values: admin or guest role token
+//
+// 需要提供 Authorization 请求头，使用配置文件中的角色令牌
+// 令牌值：admin 或 guest 角色令牌
+/*
+curl --location 'http://127.0.0.1:8001/v1/students' --header 'Authorization: 63a16b29e5bc4a28a880de1b2e707cc6'
+curl --location 'http://127.0.0.1:8001/v1/students' --header 'Authorization: 863676f1118c45c7add65b4adefd94dd'
+*/
+
+// NewRoleMiddleware creates auth middleware with token validation and route scope
+// Configure which routes need auth and setup valid tokens
+//
+// NewRoleMiddleware 创建认证中间件，进行令牌验证和路由范围控制
+// 配置需要认证的路由并设置有效令牌
+func NewRoleMiddleware(c *conf.Server, logger log.Logger) middleware.Middleware {
+	routeScope := authkratos.NewInclude( // Create INCLUDE mode route scope // 创建 INCLUDE 模式的路由范围
+		pb.OperationStudentServiceCreateStudent,
+		pb.OperationStudentServiceUpdateStudent,
+		pb.OperationStudentServiceDeleteStudent,
+		pb.OperationStudentServiceGetStudent,
+		pb.OperationStudentServiceListStudents,
+	)
+	authTokens := map[string]string{ // Setup valid tokens map // 设置有效令牌映射表
+		"admin": must.Nice(c.Auth.AdminToken),
+		"guest": must.Nice(c.Auth.GuestToken),
+	}
+	authConfig := statickratosauth.NewConfig(routeScope, authTokens).
+		WithFieldName("Authorization").
+		WithSimpleEnable(). // Enable simple token type // 启用简单令牌类型
+		WithDebugMode(true) // Enable debug mode to log auth process // 启用调试模式记录认证过程
+	return statickratosauth.NewMiddleware(authConfig, logger)
 }
```

## internal/service/student.go (+17 -4)

```diff
@@ -2,27 +2,40 @@
 
 import (
 	"context"
+	"fmt"
 
+	"github.com/go-kratos/kratos/v2/log"
 	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
 	"github.com/yylego/kratos-examples/demo1kratos/internal/biz"
+	"github.com/yylego/kratos-static-auth/statickratosauth"
+	"github.com/yylego/must"
 )
 
 type StudentService struct {
 	pb.UnimplementedStudentServiceServer
 
-	uc *biz.StudentUsecase
+	uc  *biz.StudentUsecase
+	log *log.Helper
 }
 
-func NewStudentService(uc *biz.StudentUsecase) *StudentService {
-	return &StudentService{uc: uc}
+func NewStudentService(uc *biz.StudentUsecase, logger log.Logger) *StudentService {
+	return &StudentService{uc: uc, log: log.NewHelper(logger)}
 }
 
 func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
+	// Extract and validate role name from auth context
+	//
+	// 从认证上下文中提取并验证角色名
+	roleName, ok := statickratosauth.GetUsername(ctx)
+	must.True(ok)
+	must.Nice(roleName)
+	s.log.WithContext(ctx).Infof("CreateStudent roleName=%s", roleName)
+
 	v, ebz := s.uc.CreateStudent(ctx, nil)
 	if ebz != nil {
 		return nil, ebz.Erk
 	}
-	return &pb.CreateStudentReply{Student: &pb.StudentInfo{Id: v.ID, Name: v.Name, Age: v.Age, ClassName: v.ClassName}}, nil
+	return &pb.CreateStudentReply{Student: &pb.StudentInfo{Id: v.ID, Name: fmt.Sprintf("%s (role=%s)", v.Name, roleName), Age: v.Age, ClassName: v.ClassName}}, nil
 }
 
 func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
```

