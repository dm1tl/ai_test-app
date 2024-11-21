package services

import (
	appmodels "ai_test-app/internal/app_models"
	"ai_test-app/internal/repository"
	"context"
	"fmt"
)

type AuthService struct {
	repo      repository.Authorization
	ssoClient *grpc.Client
}

func NewAuthService(repo repository.Authorization, ssoclient *grpc.Client) *AuthService {
	return &AuthService{
		repo:      repo,
		ssoClient: ssoclient,
	}
}

func (a *AuthService) Create(ctx context.Context, user appmodels.User) error {
	const op = "internal.services.Create()"
	id, err := a.ssoClient.Register(ctx, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = a.repo.Create(ctx, id)
	if err != nil {
		if rollback := a.ssoClient.Delete(ctx, id); rollback != nil {
			return fmt.Errorf("failed to rollback user in grpc after DB error %s: %w", op, rollback)
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *AuthService) Login(ctx context.Context, input appmodels.SignInInput) (string, error) {
	const op = "internal.services.Login()"
	token, err := s.ssoClient.Login(ctx, input.Email, input.Password)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (s *AuthService) Validate(ctx context.Context, token string) (int64, error) {
	const op = "pkg.service.Validate()"
	id, err := s.ssoClient.Validate(ctx, token)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}
