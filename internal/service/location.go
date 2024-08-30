package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/repo/repoerrs"
	"errors"
)

type LocationService struct {
	locationRepo repo.LocationRepo
}

func NewLocationService(locationRepo repo.LocationRepo) *LocationService {
	return &LocationService{
		locationRepo: locationRepo,
	}
}

func (s *LocationService) GetAllLocations(ctx context.Context, client any) (entity.Locations, error) {
	return s.locationRepo.GetAllLocations(ctx, client)
}

func (s *LocationService) CreateLocation(ctx context.Context, client any, input *entity.CreateLocationInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	exp := &entity.Location{
		Name:        input.Name,
		Country:     input.Country,
		NearestTown: input.NearestTown,
	}
	return s.locationRepo.CreateLocation(ctx, client, exp)
}

func (s *LocationService) DeleteLocation(ctx context.Context, client any, id int) error {
	err := s.locationRepo.DeleteLocation(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrLocationNotFound
		}
		return err
	}

	return nil
}
