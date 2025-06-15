package rule

import (
	"errors"
	"strings"
)

type RegisterInput struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func (r *RegisterInput) Validate() error {
	r.Name = strings.TrimSpace(r.Name)
	r.Email = strings.TrimSpace(strings.ToLower(r.Email))
	r.Username = strings.TrimSpace(r.Username)
	r.Password = strings.TrimSpace(r.Password)

	if r.Name == "" || r.Email == "" || r.Username == "" || r.Password == "" {
		return errors.New("all fields are required")
	}

	if len(r.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if !strings.Contains(r.Email, "@") {
		return errors.New("invalid email format")
	}

	return nil
}
