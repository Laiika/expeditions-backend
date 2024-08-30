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

type AdminRepo struct {
}

func NewAdminRepo() *AdminRepo {
	return &AdminRepo{}
}

func (r *AdminRepo) GetAdminByLogin(ctx context.Context, client any, login string) (*entity.Admin, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, login, password
		FROM admins
		WHERE login = $1
	`
	var ad entity.Admin
	err := pgClient.QueryRow(ctx, q, login).Scan(&ad.Id, &ad.Name, &ad.Login, &ad.Password)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, repoerrs.ErrNotFound
		}
		return nil, fmt.Errorf("AdminRepo GetAdminById: %v", err)
	}

	return &ad, nil
}

func (r *AdminRepo) GetAllAdmins(ctx context.Context, client any) (entity.Admins, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, login, password
		FROM admins
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("AdminRepo GetAllAdmins: %v", err)
	}

	admins := make(entity.Admins, 0)
	for rows.Next() {
		var ad entity.Admin

		err = rows.Scan(&ad.Id, &ad.Name, &ad.Login, &ad.Password)
		if err != nil {
			return nil, fmt.Errorf("AdminRepo GetAllAdmins: %v", err)
		}

		admins = append(admins, &ad)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("AdminRepo GetAllAdmins: %v", err)
	}

	return admins, nil
}

func (r *AdminRepo) CreateAdmin(ctx context.Context, client any, admin *entity.Admin) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO admins
		    (name, login, password) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, admin.Name, admin.Login, admin.Password).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("AdminRepo CreateAdmin: %v", err)
	}

	return id, nil
}

func (r *AdminRepo) DeleteAdmin(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM admins
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("AdminRepo DeleteAdmin: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
