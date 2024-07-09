package auth

import (
	"db_cp_6/internal/service"
	pkgErrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

type AuthService struct {
	member   any
	leader   any
	admin    any
	mx       sync.RWMutex
	sessions map[string]*session
}

func NewAuthService(member any, leader any, admin any) *AuthService {
	return &AuthService{
		member:   member,
		leader:   leader,
		admin:    admin,
		mx:       sync.RWMutex{},
		sessions: make(map[string]*session),
	}
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
		return nil, pkgErrors.WithMessage(service.ErrSessionNotExists, token)
	}

	return ses.GetClient(), nil
}

func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
