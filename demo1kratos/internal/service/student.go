package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/yylego/kratos-examples/demo1kratos/api/student"
	"github.com/yylego/kratos-examples/demo1kratos/internal/biz"
	"github.com/yylego/kratos-static-auth/statickratosauth"
	"github.com/yylego/must"
)

type StudentService struct {
	pb.UnimplementedStudentServiceServer

	uc  *biz.StudentUsecase
	log *log.Helper
}

func NewStudentService(uc *biz.StudentUsecase, logger log.Logger) *StudentService {
	return &StudentService{uc: uc, log: log.NewHelper(logger)}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	// Extract and validate role name from auth context
	//
	// 从认证上下文中提取并验证角色名
	roleName, ok := statickratosauth.GetUsername(ctx)
	must.True(ok)
	must.Nice(roleName)
	s.log.WithContext(ctx).Infof("CreateStudent roleName=%s", roleName)

	v, ebz := s.uc.CreateStudent(ctx, nil)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.CreateStudentReply{Student: &pb.StudentInfo{Id: v.ID, Name: fmt.Sprintf("%s (role=%s)", v.Name, roleName), Age: v.Age, ClassName: v.ClassName}}, nil
}

func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	v, ebz := s.uc.UpdateStudent(ctx, nil)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.UpdateStudentReply{Student: &pb.StudentInfo{Id: v.ID, Name: v.Name, Age: v.Age, ClassName: v.ClassName}}, nil
}

func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	if ebz := s.uc.DeleteStudent(ctx, req.Id); ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.DeleteStudentReply{Success: true}, nil
}

func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	v, ebz := s.uc.GetStudent(ctx, req.Id)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.GetStudentReply{Student: &pb.StudentInfo{Id: v.ID, Name: v.Name, Age: v.Age, ClassName: v.ClassName}}, nil
}

func (s *StudentService) ListStudents(ctx context.Context, req *pb.ListStudentsRequest) (*pb.ListStudentsReply, error) {
	students, count, ebz := s.uc.ListStudents(ctx, req.Page, req.PageSize)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.StudentInfo, 0, len(students))
	for _, v := range students {
		items = append(items, &pb.StudentInfo{Id: v.ID, Name: v.Name, Age: v.Age, ClassName: v.ClassName})
	}
	return &pb.ListStudentsReply{Students: items, Count: count}, nil
}
