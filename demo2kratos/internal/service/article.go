package service

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/yylego/kratos-examples/demo2kratos/api/article"
	"github.com/yylego/kratos-examples/demo2kratos/internal/biz"
	"github.com/yylego/kratos-static-auth/statickratosauth"
	"github.com/yylego/must"
)

type ArticleService struct {
	pb.UnimplementedArticleServiceServer

	uc  *biz.ArticleUsecase
	log *log.Helper
}

func NewArticleService(uc *biz.ArticleUsecase, logger log.Logger) *ArticleService {
	return &ArticleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *ArticleService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	// Extract and validate role name from auth context
	//
	// 从认证上下文中提取并验证角色名
	roleName, ok := statickratosauth.GetUsername(ctx)
	must.True(ok)
	must.Nice(roleName)
	s.log.WithContext(ctx).Infof("CreateArticle roleName=%s", roleName)

	v, ebz := s.uc.CreateArticle(ctx, nil)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.CreateArticleReply{Article: &pb.ArticleInfo{Id: v.ID, Title: fmt.Sprintf("%s (role=%s)", v.Title, roleName), Content: v.Content, StudentId: v.StudentID}}, nil
}

func (s *ArticleService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	v, ebz := s.uc.UpdateArticle(ctx, nil)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.UpdateArticleReply{Article: &pb.ArticleInfo{Id: v.ID, Title: v.Title, Content: v.Content, StudentId: v.StudentID}}, nil
}

func (s *ArticleService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	if ebz := s.uc.DeleteArticle(ctx, req.Id); ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.DeleteArticleReply{Success: true}, nil
}

func (s *ArticleService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	v, ebz := s.uc.GetArticle(ctx, req.Id)
	if ebz != nil {
		return nil, ebz.Erk
	}
	return &pb.GetArticleReply{Article: &pb.ArticleInfo{Id: v.ID, Title: v.Title, Content: v.Content, StudentId: v.StudentID}}, nil
}

func (s *ArticleService) ListArticles(ctx context.Context, req *pb.ListArticlesRequest) (*pb.ListArticlesReply, error) {
	articles, count, ebz := s.uc.ListArticles(ctx, req.Page, req.PageSize)
	if ebz != nil {
		return nil, ebz.Erk
	}
	items := make([]*pb.ArticleInfo, 0, len(articles))
	for _, v := range articles {
		items = append(items, &pb.ArticleInfo{Id: v.ID, Title: v.Title, Content: v.Content, StudentId: v.StudentID})
	}
	return &pb.ListArticlesReply{Articles: items, Count: count}, nil
}
