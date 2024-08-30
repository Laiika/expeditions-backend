package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"db_cp_6/internal/repo/repoerrs"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	adminRepo repo.AdminRepo
}

func NewAdminService(adminRepo repo.AdminRepo) *AdminService {
	return &AdminService{
		adminRepo: adminRepo,
	}
}

func (s *AdminService) GetAllAdmins(ctx context.Context, client any) (entity.Admins, error) {
	return s.adminRepo.GetAllAdmins(ctx, client)
}

func (s *AdminService) CreateAdmin(ctx context.Context, client any, input *entity.CreateAdminInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return 0, fmt.Errorf("AdminService CreateAdmin : %v", err)
	}

	l := &entity.Admin{
		Name:     input.Name,
		Login:    input.Login,
		Password: string(bytes),
	}
	id, err := s.adminRepo.CreateAdmin(ctx, client, l)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, ErrAdminAlreadyExists
		}
		return 0, err
	}

	return id, nil
}

func (s *AdminService) DeleteAdmin(ctx context.Context, client any, id int) error {
	err := s.adminRepo.DeleteAdmin(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrAdminNotFound
		}
		return err
	}

	return nil
}
