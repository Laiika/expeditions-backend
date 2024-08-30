package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/repo/repoerrs"
	"errors"
	"time"
)

type ExpeditionService struct {
	expeditionRepo repo.ExpeditionRepo
}

func NewExpeditionService(expeditionRepo repo.ExpeditionRepo) *ExpeditionService {
	return &ExpeditionService{
		expeditionRepo: expeditionRepo,
	}
}

func (s *ExpeditionService) GetAllExpeditions(ctx context.Context, client any) (entity.Expeditions, error) {
	return s.expeditionRepo.GetAllExpeditions(ctx, client)
}

func (s *ExpeditionService) GetLocationExpeditions(ctx context.Context, client any, locationId int) (entity.Expeditions, error) {
	return s.expeditionRepo.GetLocationExpeditions(ctx, client, locationId)
}

func (s *ExpeditionService) CreateExpedition(ctx context.Context, client any, input *entity.CreateExpeditionInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	start, _ := time.Parse("2006-01-02", input.StartDate)
	end, _ := time.Parse("2006-01-02", input.EndDate)

	exp := &entity.Expedition{
		LocationId: input.LocationId,
		StartDate:  start,
		EndDate:    end,
	}
	return s.expeditionRepo.CreateExpedition(ctx, client, exp)
}

func (s *ExpeditionService) UpdateExpeditionDates(ctx context.Context, client any, id int, startDate string, endDate string) error {
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)

	err := s.expeditionRepo.UpdateExpeditionDates(ctx, client, id, start, end)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrExpeditionNotFound
		}
		return err
	}

	return nil
}

func (s *ExpeditionService) DeleteExpedition(ctx context.Context, client any, id int) error {
	err := s.expeditionRepo.DeleteExpedition(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrExpeditionNotFound
		}
		return err
	}

	return nil
}

func (s *ExpeditionService) GetLocationExpeditionsTime(ctx context.Context, client any, locationId int) (time.Duration, error) {
	start := time.Now()
	_, err := s.expeditionRepo.GetLocationExpeditions(ctx, client, locationId)
	duration := time.Since(start)

	if err != nil {
		return 0, err
	}

	return duration, nil
}
