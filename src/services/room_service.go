package services

import (
	"context"

	"github.com/Arshia-Izadyar/Jabama-clone/src/api/dto"
	"github.com/Arshia-Izadyar/Jabama-clone/src/config"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/db"
	"github.com/Arshia-Izadyar/Jabama-clone/src/data/models"
	"github.com/Arshia-Izadyar/Jabama-clone/src/pkg/logger"
)

type RoomTypeService struct {
	base *BaseService[models.RoomType, dto.UpdateRoomTypeRequest, dto.CreateRoomTypeRequest, dto.RoomTypeResponse]
}

func NewRoomTypeService(cfg *config.Config) *RoomTypeService {
	base := &BaseService[models.RoomType, dto.UpdateRoomTypeRequest, dto.CreateRoomTypeRequest, dto.RoomTypeResponse]{
		DB:  db.GetDB(),
		Log: logger.NewLogger(cfg),
	}
	return &RoomTypeService{
		base: base,
	}
}

func (s *RoomTypeService) GetByIdRoomType(ctx context.Context, id int) (*dto.RoomTypeResponse, error) {
	return s.base.GetById(&ctx, id)
}

func (s *RoomTypeService) UpdateRoomType(ctx context.Context, req *dto.UpdateRoomTypeRequest, id int) (*dto.RoomTypeResponse, error) {
	return s.base.Update(ctx, req, id)
}

func (s *RoomTypeService) CreateRoomType(ctx context.Context, req *dto.CreateRoomTypeRequest) (*dto.RoomTypeResponse, error) {
	return s.base.Create(ctx, req)
}

func (s *RoomTypeService) DeleteRoomType(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

func (s *RoomTypeService) GetRoomTypeByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PageList[dto.RoomTypeResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
