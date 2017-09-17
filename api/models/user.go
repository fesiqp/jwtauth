package models

import "fmt"
import "golang.org/x/crypto/bcrypt"
import "github.com/jinzhu/gorm"

type UserStorer interface {
	CreateUser(*User) error
	FindUserByEmail(string) (*User, error)
	FindUserByUsername(string) (*User, error)
	FindAllUsers() ([]*User, error)
}

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
}

func NewUser(username, email, password string) (*User, error) {
	if username == "" || email == "" || password == "" {
		return nil, fmt.Errorf("Provided with empty fields")
	}
	return &User{
		Username: username,
		Email:    email,
		Password: EncryptPassword(password),
	}, nil
}

func EncryptPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (db *DB) CreateUser(user *User) error {
	u := User{}

	db.Find(&u, "email = ?", user.Email)
	if u != (User{}) {
		return fmt.Errorf("Email already in use")
	}

	db.Find(&u, "username = ?", user.Username)
	if u != (User{}) {
		return fmt.Errorf("Username already taken")
	}

	db.Create(user)

	return nil
}

func (db *DB) FindUserByEmail(email string) (*User, error) {
	u := User{}

	db.Find(&u, "email = ?", email)
	if u == (User{}) {
		return nil, fmt.Errorf("No user found with email: %s", email)
	}

	return &u, nil
}

func (db *DB) FindUserByUsername(username string) (*User, error) {
	u := User{}

	db.Find(&u, "username = ?", username)
	if u == (User{}) {
		return nil, fmt.Errorf("No user found with this username: %s", username)
	}

	return &u, nil
}

func (db *DB) FindAllUsers() ([]*User, error) {
	users := make([]*User, 0)

	db.Find(&users)

	return users, nil
}
