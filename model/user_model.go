package model

import (
	"errors"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                uint   `jsonapi:"primary,users"`
	Email             string `jsonapi:"attr, email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `jsonapi:"attr, encrypted_password"`
	Role              string `jsonapi:"attr, role"`

	CreatedAt   time.Time `jsonapi:"attr, created_at"`
	UpdatedAt   time.Time `jsonapi:"attr, updated_at"`
	LastLoginAt time.Time `jsonapi:"attr, last_login_at"`
}

func (u *User) SetPassword(p string) {
	h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.Password = ""
	u.EncryptedPassword = string(h)
}

func (u *User) ComparePassword(p string) bool {
	var (
		encryptedPassword = []byte(u.EncryptedPassword)
		password          = []byte(p)
	)

	if err := bcrypt.CompareHashAndPassword(encryptedPassword, password); err != nil {
		return false
	}
	return true
}

var regexpEmail = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

func (u *User) Validate() []error {
	var errs []error

	switch {
	case !regexpEmail.MatchString(u.Email):
		errs = append(errs, errors.New("email is invalid"))
	}
	return errs
}

func (u *User) BeforeSave() error {
	u.SetPassword(u.Password)
	return nil
}
