package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/trace"
	"layout/internal/biz"

	pb "layout/api/quizzes/v1"
)

type QuizzesService struct {
	pb.UnimplementedQuizzesServer
	uc     *biz.QuizUsecase
	log    *log.Helper
	tracer trace.Tracer
}

func NewQuizzesService(uc *biz.QuizUsecase, logger log.Logger, tracer trace.Tracer) *QuizzesService {
	return &QuizzesService{
		uc:     uc,
		log:    log.NewHelper(logger),
		tracer: tracer,
	}
}

func (s *QuizzesService) CreateQuiz(ctx context.Context, req *pb.CreateQuizRequest) (*pb.CreateQuizResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.QuizzesService.CreateQuiz")
	defer span.End()
	s.log.Debug("CreateQuiz")
	quiz := &biz.Quiz{
		UserID:      "01c49030-9c6d-4262-891d-9f4733aa8e31",
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Duration:    req.Duration,
		Thumbnail:   req.Thumbnail,
		Cover:       req.Cover,
		Category:    req.Category,
		Tags:        req.Tags,
		Metadata:    req.Metadata,
	}
	res, err := s.uc.CreateQuiz(ctx, quiz)
	if err != nil {
		s.log.Warn(err)
		return nil, err
	}
	return &pb.CreateQuizResponse{
		Quiz: &pb.Quiz{
			Id:          res.ID,
			Title:       res.Title,
			Description: res.Description,
			Duration:    res.Duration,
			Thumbnail:   res.Thumbnail,
			Cover:       res.Cover,
			Category:    res.Category,
			Tags:        res.Tags,
			Metadata:    res.Metadata,
			Audit: &pb.Audit{
				CreatedAt: res.Audit.CreatedAt,
				UpdatedAt: res.Audit.UpdatedAt,
				DeletedAt: res.Audit.DeletedAt,
			},
		},
	}, nil
}
func (s *QuizzesService) GetQuiz(ctx context.Context, req *pb.GetQuizRequest) (*pb.GetQuizResponse, error) {
	return &pb.GetQuizResponse{}, nil
}
func (s *QuizzesService) ListQuiz(ctx context.Context, req *pb.ListQuizRequest) (*pb.ListQuizResponse, error) {
	ctx, span := s.tracer.Start(ctx, "service.QuizzesService.ListQuiz")
	defer span.End()
	pagination := &biz.Pagination{
		Page: req.Pagination.GetPage(),
		Size: req.Pagination.GetPageSize(),
	}
	res, err := s.uc.ListQuiz(ctx, pagination)
	if err != nil {
		s.log.Warn(err)
		return nil, err
	}
	return &pb.ListQuizResponse{
		Quizzes: make([]*pb.Quiz, len(res)),
	}, nil
}
func (s *QuizzesService) UpdateQuiz(ctx context.Context, req *pb.UpdateQuizRequest) (*pb.UpdateQuizResponse, error) {
	return &pb.UpdateQuizResponse{}, nil
}
func (s *QuizzesService) DeleteQuiz(ctx context.Context, req *pb.DeleteQuizRequest) (*pb.DeleteQuizResponse, error) {
	return &pb.DeleteQuizResponse{}, nil
}
func (s *QuizzesService) SearchQuiz(ctx context.Context, req *pb.SearchQuizRequest) (*pb.SearchQuizResponse, error) {
	return &pb.SearchQuizResponse{}, nil
}
