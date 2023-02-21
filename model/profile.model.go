package model

import (
	"errors"
	"gorm.io/gorm"
	"strings"
)

type Profile struct {
	Id        uint   `gorm:"primaryKey" `
	Name      string `gorm:"not null:true" json:"name"`
	LastName  string `gorm:"not null:true" json:"lastname"`
	Address   string `json:"address"`
	Img       string `json:"img"`
	Age       int    `json:"age"`
	UserId    uint
	DeletedAt gorm.DeletedAt
}

func (profile *Profile) ValidProfile(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if profile.Name == "" {
			return errors.New("empty Name field")
		}
		if profile.LastName == "" {
			return errors.New("empty LastName field")
		}
		if profile.Address == "" {
			return errors.New("empty Address field")
		}
		if profile.Img == "" {
			return errors.New("empty Img field")
		}
		if profile.Age == 0 {
			return errors.New("empty Age field")
		}
		return nil
	default:
		if profile.Name == "" {
			return errors.New("empty Name field")
		}
		if profile.LastName == "" {
			return errors.New("empty LastName field")
		}
		if profile.Address == "" {
			return errors.New("empty Address field")
		}
		if profile.Img == "" {
			return errors.New("empty Img field")
		}
		if profile.Age == 0 {
			return errors.New("empty Age field")
		}
		return nil
	}
}
