package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/surrealdb/surrealdb.go"
	"github.com/surrealdb/surrealdb.go/pkg/models"
	"go.opentelemetry.io/otel/trace"
	"layout/internal/biz"
	"time"
)

type Quiz struct {
	ID          *models.RecordID  `json:"id,omitempty"`
	UserID      string            `json:"user_id,omitempty"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Duration    *uint64           `json:"duration"`
	Thumbnail   *string           `json:"thumbnail"`
	Cover       *string           `json:"cover"`
	Category    *string           `json:"category"`
	Tags        []string          `json:"tags"`
	Metadata    map[string]string `json:"metadata"`
	CreatedBy   string            `json:"created_by"`
	UpdatedBy   string            `json:"updated_by"`
	DeletedBy   string            `json:"deleted_by"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	DeletedAt   string            `json:"deleted_at"`
}

type QuizRepo struct {
	DB     *surrealdb.DB
	log    *log.Helper
	tracer trace.Tracer
	table  string
}

func NewQuizRepo(data *Data, logger log.Logger, tracer trace.Tracer) biz.QuizRepo {
	return &QuizRepo{
		DB:     data.surreal,
		log:    log.NewHelper(logger),
		tracer: tracer,
		table:  "quizzes",
	}
}

func (r *QuizRepo) Save(ctx context.Context, q *biz.Quiz) (*biz.Quiz, error) {
	ctx, span := r.tracer.Start(ctx, "data.QuizRepo.Save")
	defer span.End()
	createdAt := time.Now().String()

	res, err := surrealdb.Create[Quiz](r.DB, r.table, &Quiz{
		UserID:      q.UserID,
		Title:       q.Title,
		Description: q.Description,
		Duration:    q.Duration,
		Thumbnail:   q.Thumbnail,
		Cover:       q.Cover,
		Category:    q.Category,
		Tags:        q.Tags,
		Metadata:    q.Metadata,
		CreatedBy:   q.UserID,
		UpdatedBy:   q.UserID,
		DeletedBy:   "",
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
		DeletedAt:   "",
	})
	if err != nil {
		r.log.Warn(err)
		return nil, err
	}
	r.log.Info("created quiz", res)
	r.log.Infof("created quiz object: %+v", res)
	return &biz.Quiz{
		ID:          res.ID.String(),
		UserID:      res.UserID,
		Title:       res.Title,
		Description: res.Description,
		Duration:    res.Duration,
		Thumbnail:   res.Thumbnail,
		Cover:       res.Cover,
		Category:    res.Category,
		Tags:        res.Tags,
		Metadata:    res.Metadata,
		Audit: &biz.Audit{
			CreatedBy: res.CreatedBy,
			UpdatedBy: res.UpdatedBy,
			DeletedBy: res.DeletedBy,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
		},
	}, nil
}

func (r *QuizRepo) GetByID(ctx context.Context, id string) (*biz.Quiz, error) {
	ctx, span := r.tracer.Start(ctx, "data.QuizRepo.GetByID")
	defer span.End()
	recordID := models.NewRecordID(r.table, id)

	res, err := surrealdb.Select[Quiz, models.RecordID](r.DB, recordID)
	if err != nil {
		r.log.Warn(err)
		return nil, err
	}
	return &biz.Quiz{
		ID:          res.ID.String(),
		Title:       res.Title,
		Description: res.Description,
		Duration:    res.Duration,
		Thumbnail:   res.Thumbnail,
		Cover:       res.Cover,
		Category:    res.Category,
		Tags:        res.Tags,
		Metadata:    res.Metadata,
		Audit: &biz.Audit{
			CreatedBy: res.CreatedBy,
			UpdatedBy: res.UpdatedBy,
			DeletedBy: res.DeletedBy,
			CreatedAt: res.CreatedAt,
			UpdatedAt: res.UpdatedAt,
			DeletedAt: res.DeletedAt,
		},
	}, nil
}

func (r *QuizRepo) List(ctx context.Context, pagination *biz.Pagination) ([]*biz.Quiz, error) {
	ctx, span := r.tracer.Start(ctx, "data.QuizRepo.List")
	defer span.End()

	res, err := surrealdb.Select[[]Quiz, models.Table](r.DB, models.Table(r.table))
	if err != nil {
		r.log.Warn(err)
		return nil, err
	}
	var resList []*biz.Quiz
	for _, res := range *res {
		resList = append(resList, &biz.Quiz{
			ID:          res.ID.String(),
			Title:       res.Title,
			Description: res.Description,
			Duration:    res.Duration,
			Thumbnail:   res.Thumbnail,
			Cover:       res.Cover,
			Category:    res.Category,
			Tags:        res.Tags,
			Metadata:    res.Metadata,
			Audit: &biz.Audit{
				CreatedBy: res.CreatedBy,
				UpdatedBy: res.UpdatedBy,
				DeletedBy: res.DeletedBy,
				CreatedAt: res.CreatedAt,
				UpdatedAt: res.UpdatedAt,
				DeletedAt: res.DeletedAt,
			},
		})
	}
	return resList, nil
}

func (r *QuizRepo) Update(ctx context.Context, q *biz.Quiz) (*biz.Quiz, error) {
	ctx, span := r.tracer.Start(ctx, "data.QuizRepo.Update")
	defer span.End()

	return nil, errors.InternalServer("not implemented", "not implemented")
}
func (r *QuizRepo) Delete(ctx context.Context, id string) (*biz.Quiz, error) {
	ctx, span := r.tracer.Start(ctx, "data.QuizRepo.Delete")
	defer span.End()

	return nil, errors.InternalServer("not implemented", "not implemented")
}
func (r *QuizRepo) Search(ctx context.Context, keyword string, pagination *biz.Pagination) ([]*biz.Quiz, error) {
	ctx, span := r.tracer.Start(ctx, "data.QuizRepo.Search")
	defer span.End()

	return nil, errors.InternalServer("not implemented", "not implemented")
}
