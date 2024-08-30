package service

import (
	"context"
	"db_cp_6/internal/entity"
	"db_cp_6/internal/repo"
	"fmt"
	pkgErrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type AuthService struct {
	member     any
	leader     any
	admin      any
	mx         sync.RWMutex
	sessions   map[string]*session
	memberRepo repo.MemberRepo
	leaderRepo repo.LeaderRepo
	adminRepo  repo.AdminRepo
}

func NewAuthService(member any, leader any, admin any, memberRepo repo.MemberRepo, leaderRepo repo.LeaderRepo, adminRepo repo.AdminRepo) *AuthService {
	return &AuthService{
		member:     member,
		leader:     leader,
		admin:      admin,
		mx:         sync.RWMutex{},
		sessions:   make(map[string]*session),
		memberRepo: memberRepo,
		leaderRepo: leaderRepo,
		adminRepo:  adminRepo,
	}
}

func (s *AuthService) Login(ctx context.Context, data *entity.LoginInput) (string, error) {
	var ses *session
	member, err := s.memberRepo.GetMemberByLogin(ctx, s.admin, data.Login)
	if err != nil || !checkPassword(data.Password, member.Password) {
		leader, err := s.leaderRepo.GetLeaderByLogin(ctx, s.admin, data.Login)
		if err != nil || !checkPassword(data.Password, leader.Password) {
			admin, err := s.adminRepo.GetAdminByLogin(ctx, s.admin, data.Login)
			if err != nil || !checkPassword(data.Password, admin.Password) {
				return "", fmt.Errorf("AuthService Login: incorrect login or password")
			} else {
				ses = NewSession(s.member, s.leader, s.admin, admin.Id, "admin")
			}
		} else {
			ses = NewSession(s.member, s.leader, s.admin, leader.Id, "leader")
		}
	} else {
		ses = NewSession(s.member, s.leader, s.admin, member.Id, "member")
	}

	s.mx.Lock()
	s.sessions[ses.GetToken()] = ses
	s.mx.Unlock()

	return ses.GetToken(), nil
}

func (s *AuthService) Logout(token string) error {
	s.mx.RLock()
	_, ok := s.sessions[token]
	s.mx.RUnlock()
	if !ok {
		return ErrSessionNotExists
	}

	s.mx.Lock()
	delete(s.sessions, token)
	s.mx.Unlock()

	return nil
}

func (s *AuthService) GetSession(token string) bool {
	s.mx.RLock()
	_, ok := s.sessions[token]
	s.mx.RUnlock()

	return ok
}

func (s *AuthService) GetClient(token string) (any, error) {
	s.mx.RLock()
	ses, ok := s.sessions[token]
	s.mx.RUnlock()
	if !ok {
		return nil, pkgErrors.WithMessage(ErrSessionNotExists, token)
	}

	return ses.GetClient(), nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
