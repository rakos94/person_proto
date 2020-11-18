package responses

import "person_proto/models"

// LoginResponse ...
type LoginResponse struct {
	Token string `json:"token"`
}

// NewLoginResponse ...
func NewLoginResponse(person *models.Person) *LoginResponse {
	return &LoginResponse{
		Token: person.Token,
	}
}
