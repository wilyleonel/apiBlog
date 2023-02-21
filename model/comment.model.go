package model

import (
	"errors"
	"strings"
)

type Comment struct {
	Id          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	PostId      uint
	UserId      uint
	Like        []Like `gorm:"constraint:OnDelete:CASCADE;"`
}

func (comment *Comment) ValidComment(action string) error {
	switch strings.ToLower(action) {
	case "create":
		if comment.Description == "" {
			return errors.New("empty Description field")
		}
		return nil
	default:
		if comment.Description == "" {
			return errors.New("se requiere correo electronico")
		}
		return nil
	}
}
