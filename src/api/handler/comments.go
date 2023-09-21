package handler

import (
	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/services"
	"github.com/gin-gonic/gin"
)

type ResidenceCommentHandler struct {
	service *services.ResidenceCommentService
}

func NewResidenceCommentHandler(cfg *config.Config) *ResidenceCommentHandler {
	s := services.NewResidenceCommentService(cfg)
	return &ResidenceCommentHandler{
		service: s,
	}
}

func (ch *ResidenceCommentHandler) CreateResidenceComment(ctx *gin.Context) {

	Create[dto.CreateResidenceCommentRequest, dto.ResidenceCommentResponse](ctx, ch.service.CreateResidenceComment)
}

func (ch *ResidenceCommentHandler) GetById(ctx *gin.Context) {

	GetById[dto.ResidenceCommentResponse](ctx, ch.service.GetByIdResidenceComment)
}

func (ch *ResidenceCommentHandler) UpdateResidenceComment(ctx *gin.Context) {

	Update[dto.UpdateResidenceCommentRequest, dto.ResidenceCommentResponse](ctx, ch.service.UpdateResidenceComment)

}

func (ch *ResidenceCommentHandler) DeleteResidenceComment(ctx *gin.Context) {

	Delete(ctx, ch.service.DeleteResidenceComment)
}

func (ch *ResidenceCommentHandler) GetByFilter(ctx *gin.Context) {
	GetByFilter[dto.PaginationInputWithFilter, dto.ResidenceCommentResponse](ctx, ch.service.GetResidenceCommentByFilter)
}
