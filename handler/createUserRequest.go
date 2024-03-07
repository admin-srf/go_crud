package handler

import (
	"github.com/admin-srf/go_crud/schemas"
	"github.com/admin-srf/go_crud/services"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return services.ErrParamIsRequired("name", "string")
	}

	if r.Password == "" {
		return services.ErrParamIsRequired("password", "string")
	}
	return nil
}

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := schemas.HashPassword(password)
	if err != nil {
		return hashedPassword, err
	}
	return hashedPassword, nil
}
