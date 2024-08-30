package pgdb

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/repoerrs"
	"db_cp_6/pkg/postgres"
	"fmt"
)

type LocationRepo struct {
}

func NewLocationRepo() *LocationRepo {
	return &LocationRepo{}
}

func (r *LocationRepo) GetAllLocations(ctx context.Context, client any) (entity.Locations, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, country, nearest_town
		FROM locations
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("LocationRepo GetAllLocations: %v", err)
	}

	locations := make(entity.Locations, 0)
	for rows.Next() {
		var l entity.Location

		err = rows.Scan(&l.Id, &l.Name, &l.Country, &l.NearestTown)
		if err != nil {
			return nil, fmt.Errorf("LocationRepo GetAllLocations: %v", err)
		}

		locations = append(locations, &l)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("LocationRepo GetAllLocations: %v", err)
	}

	return locations, nil
}

func (r *LocationRepo) CreateLocation(ctx context.Context, client any, location *entity.Location) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO locations
		    (name, country, nearest_town) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, location.Name, location.Country, location.NearestTown).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("LocationRepo CreateLocation: %v", err)
	}

	return id, nil
}

func (r *LocationRepo) DeleteLocation(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM locations
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("LocationRepo DeleteLocation: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
