package pgdb

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/pkg/postgres"
	"fmt"
)

type ArtifactRepo struct {
}

func NewArtifactRepo() *ArtifactRepo {
	return &ArtifactRepo{}
}

func (r *ArtifactRepo) GetLocationArtifacts(ctx context.Context, client any, locationId int) (entity.Artifacts, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, location_id, name, age
		FROM artifacts
		WHERE location_id = $1
	`
	rows, err := pgClient.Query(ctx, q, locationId)
	if err != nil {
		return nil, fmt.Errorf("ArtifactRepo GetLocationArtifacts: %v", err)
	}

	artifacts := make(entity.Artifacts, 0)
	for rows.Next() {
		var ar entity.Artifact

		err = rows.Scan(&ar.Id, &ar.LocationId, &ar.Name, &ar.Age)
		if err != nil {
			return nil, fmt.Errorf("ArtifactRepo GetLocationArtifacts: %v", err)
		}

		artifacts = append(artifacts, &ar)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ArtifactRepo GetLocationArtifacts: %v", err)
	}

	return artifacts, nil
}

func (r *ArtifactRepo) GetAllArtifacts(ctx context.Context, client any) (entity.Artifacts, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, location_id, name, age
		FROM artifacts
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("ArtifactRepo GetAllArtifacts: %v", err)
	}

	artifacts := make(entity.Artifacts, 0)
	for rows.Next() {
		var ar entity.Artifact

		err = rows.Scan(&ar.Id, &ar.LocationId, &ar.Name, &ar.Age)
		if err != nil {
			return nil, fmt.Errorf("ArtifactRepo GetAllArtifacts: %v", err)
		}

		artifacts = append(artifacts, &ar)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ArtifactRepo GetAllArtifacts: %v", err)
	}

	return artifacts, nil
}

func (r *ArtifactRepo) CreateArtifact(ctx context.Context, client any, artifact *entity.Artifact) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO artifacts
		    (location_id, name, age) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, artifact.LocationId, artifact.Name, artifact.Age).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ArtifactRepo CreateArtifact: %v", err)
	}

	return id, nil
}
