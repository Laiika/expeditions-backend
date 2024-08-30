package pgdb

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/repoerrs"
	"db_cp_6/pkg/postgres"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
)

type CuratorRepo struct {
}

func NewCuratorRepo() *CuratorRepo {
	return &CuratorRepo{}
}

func (r *CuratorRepo) GetExpeditionCurators(ctx context.Context, client any, expeditionId int) (entity.Curators, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT c.id, c.name
		FROM curators c
		JOIN expeditions_curators ec ON ec.curator_id = c.id
		WHERE ec.expedition_id = $1
	`
	rows, err := pgClient.Query(ctx, q, expeditionId)
	if err != nil {
		return nil, fmt.Errorf("CuratorRepo GetExpeditionCurators: %v", err)
	}

	curators := make(entity.Curators, 0)
	for rows.Next() {
		var c entity.Curator

		err = rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, fmt.Errorf("CuratorRepo GetExpeditionCurators: %v", err)
		}

		curators = append(curators, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("CuratorRepo GetExpeditionCurators: %v", err)
	}

	return curators, nil
}

func (r *CuratorRepo) GetAllCurators(ctx context.Context, client any) (entity.Curators, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name
		FROM curators
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("CuratorRepo GetAllCurators: %v", err)
	}

	curators := make(entity.Curators, 0)
	for rows.Next() {
		var c entity.Curator

		err = rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, fmt.Errorf("CuratorRepo GetAllCurators: %v", err)
		}

		curators = append(curators, &c)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("CuratorRepo GetAllCurators: %v", err)
	}

	return curators, nil
}

func (r *CuratorRepo) CreateCurator(ctx context.Context, client any, curator *entity.Curator) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO curators
		    (name) 
		VALUES 
		    ($1) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, curator.Name).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("CuratorRepo CreateCurator: %v", err)
	}

	return id, nil
}

func (r *CuratorRepo) CreateCuratorExpedition(ctx context.Context, client any, curatorId int, expeditionId int) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO expeditions_curators
		    (curator_id, expedition_id) 
		VALUES 
		    ($1, $2) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, curatorId, expeditionId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("CuratorRepo CreateCuratorExpedition: %v", err)
	}

	return id, nil
}

func (r *CuratorRepo) DeleteCurator(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM curators
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("CuratorRepo DeleteCurator: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
