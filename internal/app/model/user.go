package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User...
type User struct {
	ID                int    `param:"id" query:"id" form:"id" json:"id" xml:"id"`
	Email             string `param:"email" query:"email" form:"email" json:"email" xml:"email"`
	Password          string `param:"password" query:"password" form:"password" json:"password" xml:"password"`
	EncryptedPassword string `param:"-" query:"-" form:"-" json:"-" xml:"-"`
}

// Validate...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

// Sanitize ...
func (u *User) Sanitize() {
	u.Password = ""
}

// ComparePassword...
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

// BeforeCreate...
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}

		u.EncryptedPassword = enc
	}

	return nil
}

// encryptString...
func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
