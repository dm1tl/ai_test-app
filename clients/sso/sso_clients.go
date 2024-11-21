package sso

import (
	"context"
)

type SSOProvider interface {
	Register(ctx context.Context, email string, password string) (int64, error)
	Login(ctx context.Context, email string, password string) (string, error)
	Validate(ctx context.Context, token string) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type SSOClient struct {
	client string
}

func (s *SSOClient) Register(ctx context.Context, email string, password string) (int64, error) {
	return 0, nil
}
