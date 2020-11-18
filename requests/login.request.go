package requests

// Login ...
type Login struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required"`
	Token    string `json:"token" gorm:"-"`
}
