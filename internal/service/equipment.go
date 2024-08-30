package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/repo/repoerrs"
	"errors"
)

type EquipmentService struct {
	equipmentRepo repo.EquipmentRepo
}

func NewEquipmentService(equipmentRepo repo.EquipmentRepo) *EquipmentService {
	return &EquipmentService{
		equipmentRepo: equipmentRepo,
	}
}

func (s *EquipmentService) GetExpeditionEquipments(ctx context.Context, client any, expeditionId int) (entity.Equipments, error) {
	return s.equipmentRepo.GetExpeditionEquipments(ctx, client, expeditionId)
}

func (s *EquipmentService) GetAllEquipments(ctx context.Context, client any) (entity.Equipments, error) {
	return s.equipmentRepo.GetAllEquipments(ctx, client)
}

func (s *EquipmentService) CreateEquipment(ctx context.Context, client any, input *entity.CreateEquipmentInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	exp := &entity.Equipment{
		ExpeditionId: input.ExpeditionId,
		Name:         input.Name,
		Amount:       input.Amount,
	}
	return s.equipmentRepo.CreateEquipment(ctx, client, exp)
}

func (s *EquipmentService) DeleteEquipment(ctx context.Context, client any, id int) error {
	err := s.equipmentRepo.DeleteEquipment(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrEquipmentNotFound
		}
		return err
	}

	return nil
}
