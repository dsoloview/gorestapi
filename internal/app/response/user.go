package response

import "github.com/dsoloview/gorestapi/internal/app/model"

type CreateUserResponse struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

func NewCreateUserResponse(user *model.User) *CreateUserResponse {
	return &CreateUserResponse{
		Id:    user.ID,
		Email: user.Email,
	}
}
