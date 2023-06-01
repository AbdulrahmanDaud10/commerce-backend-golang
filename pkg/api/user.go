package api

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID                string `bson:"_id,omitempty" json:"id"`
	Email             string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"_"`
	IsAdmin           bool   `bson:"isAdmin" json:"isAdmin"`
	Token             string `bson:"token" json:"token"`
}

func NewAdminUser(email, password string) (*User, error) {
	user, err := NewUser(email, password)
	if err != nil {
		return nil, err
	}

	user.IsAdmin = true
	return user, nil
}

func NewUser(email, password string) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:             email,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password))
	return err == nil
}
