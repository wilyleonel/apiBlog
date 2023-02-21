package model

import (
	"errors"
	"strings"
)

type Category struct {
	Id       uint   `gorm:"primaryKey"`
	Category string `json:"category"`
	Post     []Post `gorm:"many2many:post_categories;"`
}

func (category *Category) ValidCategory(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if category.Category == "" {
			return errors.New("empty Description field")
		}
		return nil
	default:
		if category.Category == "" {
			return errors.New("se requiere correo electronico")
		}
		return nil
	}
}