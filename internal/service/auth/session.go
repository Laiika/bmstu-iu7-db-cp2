package auth

import (
	"db_cp_6/internal/service"
	"github.com/google/uuid"
)

type session struct {
	token  string
	userId int
	role   string
	client service.Client
}

func NewSession(member service.Client, leader service.Client, admin service.Client, id int, role string) *session {
	ses := &session{
		token:  uuid.NewString(),
		userId: id,
		role:   role,
	}

	switch role {
	case "member":
		ses.client = member
	case "leader":
		ses.client = leader
	case "admin":
		ses.client = admin
	}

	return ses
}

func (s *session) GetToken() string {
	return s.token
}

func (s *session) GetRole() string {
	return s.role
}

func (s *session) GetClient() service.Client {
	return s.client
}
