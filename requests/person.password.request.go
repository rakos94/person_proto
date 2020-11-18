package requests

// PersonPasswordRequest ...
type PersonPasswordRequest struct {
	Password string `form:"password" validate:"required"`
}
