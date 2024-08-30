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

type MemberService struct {
	memberRepo repo.MemberRepo
}

func NewMemberService(memberRepo repo.MemberRepo) *MemberService {
	return &MemberService{
		memberRepo: memberRepo,
	}
}

func (s *MemberService) GetExpeditionMembers(ctx context.Context, client any, expeditionId int) (entity.Members, error) {
	return s.memberRepo.GetExpeditionMembers(ctx, client, expeditionId)
}

func (s *MemberService) GetAllMembers(ctx context.Context, client any) (entity.Members, error) {
	return s.memberRepo.GetAllMembers(ctx, client)
}

func (s *MemberService) CreateMember(ctx context.Context, client any, input *entity.CreateMemberInput) (int, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
	if err != nil {
		return 0, fmt.Errorf("MemberService CreateMember : %v", err)
	}

	m := &entity.Member{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Login:       input.Login,
		Password:    string(bytes),
	}
	id, err := s.memberRepo.CreateMember(ctx, client, m)
	if err != nil {
		if errors.Is(err, repoerrs.ErrAlreadyExists) {
			return 0, ErrMemberAlreadyExists
		}
		return 0, err
	}

	return id, nil
}

func (s *MemberService) CreateMemberExpedition(ctx context.Context, client any, memberId int, expeditionId int) (int, error) {
	return s.memberRepo.CreateMemberExpedition(ctx, client, memberId, expeditionId)
}

func (s *MemberService) DeleteMember(ctx context.Context, client any, id int) error {
	err := s.memberRepo.DeleteMember(ctx, client, id)
	if err != nil {
		if errors.Is(err, repoerrs.ErrNotFound) {
			return ErrMemberNotFound
		}
		return err
	}

	return nil
}
