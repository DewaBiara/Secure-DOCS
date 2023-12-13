package service

import (
	"context"

	"github.com/DewaBiara/Secure-DOCS/internal/user/dto"
)

type UserService interface {
	SignUpUser(ctx context.Context, user *dto.UserSignUpRequest) error
	LogInUser(ctx context.Context, user *dto.UserLoginRequest) (string, error)
	GetBriefUsers(ctx context.Context, page int, limit int) (*dto.BriefUsersResponse, error)
	UpdateUser(ctx context.Context, userID string, request *dto.UserUpdateRequest) error
}
