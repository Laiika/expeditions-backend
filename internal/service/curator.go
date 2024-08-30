package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/repo/repoerrs"
	"errors"
)

type CuratorService struct {
	curatorRepo repo.CuratorRepo
}

func NewCuratorService(curatorRepo repo.CuratorRepo) *CuratorService {
	return &CuratorService{
		curatorRepo: curatorRepo,
	}
}

func (s *CuratorService) GetExpeditionCurators(ctx context.Context, client any, expeditionId int) (entity.Curators, error) {
	return s.curatorRepo.GetExpeditionCurators(ctx, client, expeditionId)
}

func (s *CuratorService) GetAllCurators(ctx context.Context, client any) (entity.Curators, error) {
	return s.curatorRepo.GetAllCurators(ctx, client)
}

func (s *CuratorService) CreateCurator(ctx context.Context, client any, input *entity.CreateCuratorInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	m := &entity.Curator{
		Name: input.Name,
	}
	id, err := s.curatorRepo.CreateCurator(ctx, client, m)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, ErrCuratorAlreadyExists
		}
		return 0, err
	}

	return id, nil
}

func (s *CuratorService) CreateCuratorExpedition(ctx context.Context, client any, curatorId int, expeditionId int) (int, error) {
	return s.curatorRepo.CreateCuratorExpedition(ctx, client, curatorId, expeditionId)
}

func (s *CuratorService) DeleteCurator(ctx context.Context, client any, id int) error {
	err := s.curatorRepo.DeleteCurator(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrCuratorNotFound
		}
		return err
	}

	return nil
}
