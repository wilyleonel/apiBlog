package model

import (
	"apiBlog/database"
	"errors"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserID uint
	Name   string
	Email  string
	*jwt.StandardClaims
}

type User struct {
	Id        uint      `gorm:"primaryKey"`
	Email     string    `gorm:"unique:true not null:true" json:"email"`
	Password  string    `gorm:"not null:true" json:"password"`
	Profile   Profile   `gorm:"constraint:OnDelete:CASCADE;"`
	Post      []Post    `gorm:"constraint:OnDelete:CASCADE;"`
	Comment   []Comment `gorm:"constraint:OnDelete:CASCADE;"`
	Like      []Like    `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdat" `
}

func (user *User) ValidUser(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if user.Email == "" {
			return errors.New("empty Email field")
		}
		if user.Password == "" {
			return errors.New("empty password field")
		}
		return nil
	case "update":
		if user.Password == "" {
			return errors.New("empty password field")
		}
		return nil
	default:
		if user.Email == "" {
			return errors.New("empty Email field")
		}
		if user.Password == "" {
			return errors.New(fmt.Sprintf("empty %s field", user.Password))
		}
		return nil
	}
}

func Delete() {
	database.DB.Exec(`drop table users cascade`)
	database.DB.Exec(`drop table profiles`)
	database.DB.Exec(`drop table posts cascade`)
	database.DB.Exec(`drop table comments`)
	database.DB.Exec(`drop table post_categories`)
	database.DB.Exec(`drop table categories`)
}

func Migration() {
	database.DB.AutoMigrate(&User{}, &Category{},
		&Comment{}, &Post{}, &Profile{},&Like{})
}
