package adminrequests

import authrequests "github.com/sherwin-77/golang-basic-auth-api/requests/auth"

type UserRequest struct {
	authrequests.RegisterRequest
}

type UpdateUserRequest struct {
	authrequests.RegisterRequest
	ID string `param:"id" validate:"required"`
}
