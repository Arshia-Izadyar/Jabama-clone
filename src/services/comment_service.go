package services

import (
	"context"
	"fmt"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/constants"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type ResidenceCommentService struct {
	base *BaseService[models.ResidenceComment, dto.UpdateResidenceCommentRequest, dto.CreateResidenceCommentRequest, dto.ResidenceCommentResponse]
}

func NewResidenceCommentService(cfg *config.Config) *ResidenceCommentService {
	base := &BaseService[models.ResidenceComment, dto.UpdateResidenceCommentRequest, dto.CreateResidenceCommentRequest, dto.ResidenceCommentResponse]{
		DB:       db.GetDB(),
		Log:      logger.NewLogger(cfg),
		Preloads: []preload{{name: "User"}},
	}
	return &ResidenceCommentService{
		base: base,
	}
}

func (s *ResidenceCommentService) GetByIdResidenceComment(ctx context.Context, id int) (*dto.ResidenceCommentResponse, error) {
	return s.base.GetById(&ctx, id)
}

func (s *ResidenceCommentService) UpdateResidenceComment(ctx context.Context, req *dto.UpdateResidenceCommentRequest, id int) (*dto.ResidenceCommentResponse, error) {
	return s.base.Update(ctx, req, id)
}

func (s *ResidenceCommentService) CreateResidenceComment(ctx context.Context, req *dto.CreateResidenceCommentRequest) (*dto.ResidenceCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	fmt.Println(req.UserId)
	return s.base.Create(ctx, req)
}

func (s *ResidenceCommentService) DeleteResidenceComment(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *ResidenceCommentService) GetResidenceCommentByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.ResidenceCommentResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
