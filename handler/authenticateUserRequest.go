package handler

import "github.com/admin-srf/go_crud/services"

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthenticateUserRequest) Validate() error {
	if r.Email == "" {
		return services.ErrParamIsRequired("name", "string")
	}

	if r.Password == "" {
		return services.ErrParamIsRequired("password", "string")
	}
	return nil
}
