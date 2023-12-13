package dto

import "github.com/DewaBiara/Secure-DOCS/pkg/entity"

type UserSignUpRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Telp     string `json:"telp" validate:"required"`
}

func (u *UserSignUpRequest) ToEntity() *entity.User {
	return &entity.User{
		Username: u.Username,
		Password: u.Password,
		Name:     u.Name,
		Telp:     u.Telp,
	}
}

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Telp     string `json:"telp"`
}

func (u *UserUpdateRequest) ToEntity() *entity.User {
	return &entity.User{
		Username: u.Username,
		Password: u.Password,
		Name:     u.Name,
		Telp:     u.Telp,
	}
}

type BriefUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

func NewBriefUserResponse(user *entity.User) *BriefUserResponse {
	return &BriefUserResponse{
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
	}
}

type BriefUsersResponse []BriefUserResponse

func NewBriefUsersResponse(users *entity.Users) *BriefUsersResponse {
	var briefUsersResponse BriefUsersResponse
	for _, user := range *users {
		briefUsersResponse = append(briefUsersResponse, *NewBriefUserResponse(&user))
	}
	return &briefUsersResponse
}
