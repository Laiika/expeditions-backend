package pgdb

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo/repoerrs"
	"db_cp_6/pkg/postgres"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	pkgErrors "github.com/pkg/errors"
)

type LeaderRepo struct {
}

func NewLeaderRepo() *LeaderRepo {
	return &LeaderRepo{}
}

func (r *LeaderRepo) GetLeaderByLogin(ctx context.Context, client any, login string) (*entity.Leader, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, phone_number, login, password
		FROM leaders
		WHERE login = $1
	`
	var l entity.Leader
	err := pgClient.QueryRow(ctx, q, login).Scan(&l.Id, &l.Name, &l.PhoneNumber, &l.Login, &l.Password)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, repoerrs.ErrNotFound
		}
		return nil, fmt.Errorf("LeaderRepo GetLeaderById: %v", err)
	}

	return &l, nil
}

func (r *LeaderRepo) GetExpeditionLeaders(ctx context.Context, client any, expeditionId int) (entity.Leaders, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT l.id, l.name, l.phone_number, l.login, l.password
		FROM leaders l
		JOIN expeditions_leaders el ON el.leader_id = l.id
		WHERE el.expedition_id = $1
	`
	rows, err := pgClient.Query(ctx, q, expeditionId)
	if err != nil {
		return nil, fmt.Errorf("LeaderRepo GetExpeditionLeaders: %v", err)
	}

	leaders := make(entity.Leaders, 0)
	for rows.Next() {
		var l entity.Leader

		err = rows.Scan(&l.Id, &l.Name, &l.PhoneNumber, &l.Login, &l.Password)
		if err != nil {
			return nil, fmt.Errorf("LeaderRepo GetExpeditionLeaders: %v", err)
		}

		leaders = append(leaders, &l)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("LeaderRepo GetExpeditionLeaders: %v", err)
	}

	return leaders, nil
}

func (r *LeaderRepo) GetAllLeaders(ctx context.Context, client any) (entity.Leaders, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, phone_number, login, password
		FROM leaders
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("LeaderRepo GetAllLeaders: %v", err)
	}

	leaders := make(entity.Leaders, 0)
	for rows.Next() {
		var l entity.Leader

		err = rows.Scan(&l.Id, &l.Name, &l.PhoneNumber, &l.Login, &l.Password)
		if err != nil {
			return nil, fmt.Errorf("LeaderRepo GetAllLeaders: %v", err)
		}

		leaders = append(leaders, &l)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("LeaderRepo GetAllLeaders: %v", err)
	}

	return leaders, nil
}

func (r *LeaderRepo) CreateLeader(ctx context.Context, client any, leader *entity.Leader) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO leaders
		    (name, phone_number, login, password) 
		VALUES 
		    ($1, $2, $3, $4) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, leader.Name, leader.PhoneNumber, leader.Login, leader.Password).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("LeaderRepo CreateLeader: %v", err)
	}

	return id, nil
}

func (r *LeaderRepo) CreateLeaderExpedition(ctx context.Context, client any, leaderId int, expeditionId int) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO expeditions_leaders
		    (leader_id, expedition_id) 
		VALUES 
		    ($1, $2) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, leaderId, expeditionId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("LeaderRepo CreateLeaderExpedition: %v", err)
	}

	return id, nil
}

func (r *LeaderRepo) DeleteLeader(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM leaders
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("LeaderRepo DeleteLeader: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
