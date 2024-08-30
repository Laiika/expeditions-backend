package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
)

type ArtifactService struct {
	artifactRepo repo.ArtifactRepo
}

func NewArtifactService(artifactRepo repo.ArtifactRepo) *ArtifactService {
	return &ArtifactService{
		artifactRepo: artifactRepo,
	}
}

func (s *ArtifactService) GetLocationArtifacts(ctx context.Context, client any, locationId int) (entity.Artifacts, error) {
	return s.artifactRepo.GetLocationArtifacts(ctx, client, locationId)
}

func (s *ArtifactService) GetAllArtifacts(ctx context.Context, client any) (entity.Artifacts, error) {
	return s.artifactRepo.GetAllArtifacts(ctx, client)
}

func (s *ArtifactService) CreateArtifact(ctx context.Context, client any, input *entity.CreateArtifactInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	exp := &entity.Artifact{
		LocationId: input.LocationId,
		Name:       input.Name,
		Age:        input.Age,
	}
	return s.artifactRepo.CreateArtifact(ctx, client, exp)
}
