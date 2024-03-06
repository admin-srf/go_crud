package handler

type AuthenticateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthenticateUserRequest) Validate() error {
	if r.Email == "" {
		return errParamIsRequired("name", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}
