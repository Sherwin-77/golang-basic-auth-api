package authrequests

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	LoginRequest
	Username string `json:"username" validate:"required,gte=3,lte=20"`
}
