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

type LeaderService struct {
	leaderRepo repo.LeaderRepo
}

func NewLeaderService(leaderRepo repo.LeaderRepo) *LeaderService {
	return &LeaderService{
		leaderRepo: leaderRepo,
	}
}

func (s *LeaderService) GetExpeditionLeaders(ctx context.Context, client any, expeditionId int) (entity.Leaders, error) {
	return s.leaderRepo.GetExpeditionLeaders(ctx, client, expeditionId)
}

func (s *LeaderService) GetAllLeaders(ctx context.Context, client any) (entity.Leaders, error) {
	return s.leaderRepo.GetAllLeaders(ctx, client)
}

func (s *LeaderService) CreateLeader(ctx context.Context, client any, input *entity.CreateLeaderInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return 0, fmt.Errorf("LeaderService CreateLeader : %v", err)
	}

	l := &entity.Leader{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Login:       input.Login,
		Password:    string(bytes),
	}
	id, err := s.leaderRepo.CreateLeader(ctx, client, l)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, ErrLeaderAlreadyExists
		}
		return 0, err
	}

	return id, nil
}

func (s *LeaderService) CreateLeaderExpedition(ctx context.Context, client any, leaderId int, expeditionId int) (int, error) {
	return s.leaderRepo.CreateLeaderExpedition(ctx, client, leaderId, expeditionId)
}

func (s *LeaderService) DeleteLeader(ctx context.Context, client any, id int) error {
	err := s.leaderRepo.DeleteLeader(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrLeaderNotFound
		}
		return err
	}

	return nil
}
