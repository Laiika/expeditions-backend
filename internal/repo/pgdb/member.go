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

type MemberRepo struct {
}

func NewMemberRepo() *MemberRepo {
	return &MemberRepo{}
}

func (r *MemberRepo) GetMemberByLogin(ctx context.Context, client any, login string) (*entity.Member, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, phone_number, role, login, password
		FROM members
		WHERE login = $1
	`
	var m entity.Member
	err := pgClient.QueryRow(ctx, q, login).Scan(&m.Id, &m.Name, &m.PhoneNumber, &m.Role, &m.Login, &m.Password)

	if err != nil {
		if pkgErrors.Is(err, pgx.ErrNoRows) {
			return nil, repoerrs.ErrNotFound
		}
		return nil, fmt.Errorf("MemberRepo GetMemberById: %v", err)
	}

	return &m, nil
}

func (r *MemberRepo) GetExpeditionMembers(ctx context.Context, client any, expeditionId int) (entity.Members, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT m.id, m.name, m.phone_number, m.role, m.login, m.password
		FROM members m
		JOIN expeditions_members em ON em.member_id = m.id
		WHERE em.expedition_id = $1
	`
	rows, err := pgClient.Query(ctx, q, expeditionId)
	if err != nil {
		return nil, fmt.Errorf("MemberRepo GetExpeditionMembers: %v", err)
	}

	members := make(entity.Members, 0)
	for rows.Next() {
		var m entity.Member

		err = rows.Scan(&m.Id, &m.Name, &m.PhoneNumber, &m.Role, &m.Login, &m.Password)
		if err != nil {
			return nil, fmt.Errorf("MemberRepo GetExpeditionMembers: %v", err)
		}

		members = append(members, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("MemberRepo GetExpeditionMembers: %v", err)
	}

	return members, nil
}

func (r *MemberRepo) GetAllMembers(ctx context.Context, client any) (entity.Members, error) {
	pgClient := client.(postgres.Client)
	q := `
		SELECT id, name, phone_number, role, login, password
		FROM members
	`
	rows, err := pgClient.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("MemberRepo GetAllMembers: %v", err)
	}

	members := make(entity.Members, 0)
	for rows.Next() {
		var m entity.Member

		err = rows.Scan(&m.Id, &m.Name, &m.PhoneNumber, &m.Role, &m.Login, &m.Password)
		if err != nil {
			return nil, fmt.Errorf("MemberRepo GetAllMembers: %v", err)
		}

		members = append(members, &m)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("MemberRepo GetAllMembers: %v", err)
	}

	return members, nil
}

func (r *MemberRepo) CreateMember(ctx context.Context, client any, member *entity.Member) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO members
		    (name, phone_number, role, login, password) 
		VALUES 
		    ($1, $2, $3, $4, $5) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, member.Name, member.PhoneNumber, member.Role, member.Login, member.Password).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if ok := errors.As(err, &pgErr); ok {
			if pgErr.Code == "23505" {
				return 0, repoerrs.ErrAlreadyExists
			}
		}
		return 0, fmt.Errorf("MemberRepo CreateMember: %v", err)
	}

	return id, nil
}

func (r *MemberRepo) CreateMemberExpedition(ctx context.Context, client any, memberId int, expeditionId int) (int, error) {
	pgClient := client.(postgres.Client)
	q := `
		INSERT INTO expeditions_members
		    (member_id, expedition_id) 
		VALUES 
		    ($1, $2) 
		RETURNING id
	`
	var id int
	err := pgClient.QueryRow(ctx, q, memberId, expeditionId).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("MemberRepo CreateMemberExpedition: %v", err)
	}

	return id, nil
}

func (r *MemberRepo) DeleteMember(ctx context.Context, client any, id int) error {
	pgClient := client.(postgres.Client)
	q := `
		DELETE FROM members
		WHERE id = $1
	`
	commandTag, err := pgClient.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("MemberRepo DeleteMember: %v", err)
	}
	if commandTag.RowsAffected() != 1 {
		return repoerrs.ErrNotFound
	}

	return nil
}
