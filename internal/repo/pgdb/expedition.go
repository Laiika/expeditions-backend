package pgdb

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/repoerrs"
	"db_cp_6/pkg/postgres"
	"fmt"
	"time"
)

type ExpeditionRepo struct {
}

func NewExpeditionRepo() *ExpeditionRepo {
	return &ExpeditionRepo{}
}

func (r *ExpeditionRepo) GetAllExpeditions(ctx context.Context, client any) (entity.Expeditions, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, location_id, start_date, end_date
		FROM expeditions
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("ExpeditionRepo GetAllExpeditions: %v", err)
	}

	expeditions := make(entity.Expeditions, 0)
	for rows.Next() {
		var exp entity.Expedition

		err = rows.Scan(&exp.Id, &exp.LocationId, &exp.StartDate, &exp.EndDate)
		if err != nil {
			return nil, fmt.Errorf("ExpeditionRepo GetAllExpeditions: %v", err)
		}

		expeditions = append(expeditions, &exp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ExpeditionRepo GetAllExpeditions: %v", err)
	}

	return expeditions, nil
}

func (r *ExpeditionRepo) GetLocationExpeditions(ctx context.Context, client any, locationId int) (entity.Expeditions, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, location_id, start_date, end_date
		FROM expeditions
		WHERE location_id = $1
	`
	rows, err := pgClient.Query(ctx, q, locationId)
	if err != nil {
		return nil, fmt.Errorf("ExpeditionRepo GetAllExpeditions: %v", err)
	}

	expeditions := make(entity.Expeditions, 0)
	for rows.Next() {
		var exp entity.Expedition

		err = rows.Scan(&exp.Id, &exp.LocationId, &exp.StartDate, &exp.EndDate)
		if err != nil {
			return nil, fmt.Errorf("ExpeditionRepo GetAllExpeditions: %v", err)
		}

		expeditions = append(expeditions, &exp)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ExpeditionRepo GetAllExpeditions: %v", err)
	}

	return expeditions, nil
}

func (r *ExpeditionRepo) CreateExpedition(ctx context.Context, client any, expedition *entity.Expedition) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO expeditions
		    (location_id, start_date, end_date) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, expedition.LocationId, expedition.StartDate, expedition.EndDate).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("ExpeditionRepo CreateExpedition: %v", err)
	}

	return id, nil
}

func (r *ExpeditionRepo) UpdateExpeditionDates(ctx context.Context, client any, id int, start time.Time, end time.Time) error {
	pgClient := client.(postgres.Client)
	q := `
		UPDATE expeditions
		SET
			start_date = $1, end_date = $2
		WHERE id = $3
	`
	commandTag, err := pgClient.Exec(ctx, q, start, end, id)
	if err != nil {
		return fmt.Errorf("ExpeditionRepo UpdateExpedition: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}

func (r *ExpeditionRepo) DeleteExpedition(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM expeditions
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("ExpeditionRepo DeleteExpedition: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
