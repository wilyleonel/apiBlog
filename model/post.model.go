package model

import (
	"errors"
	"strings"
	"time"

)

type Post struct {
	Id          uint   `gorm:"primaryKey"`
	Description string `json:"description"`
	Title       string `gorm:"not null:true" json:"title"`
	UserId      uint
	Comment     []Comment  `gorm:"constraint:OnDelete:CASCADE;"`
	Categories  []Category `gorm:"many2many:post_categories;"`
	Like        []Like     `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at" `
}

func (post *Post)ValidPost(action string)error{
	switch strings.ToLower(action){
	case "create":
		if post.Description=="empty Description field" {
			return errors.New("")
		}
		if post.Title=="" {
			return errors.New("empty Title field")
		}
		return nil
	default:
		if post.Description=="" {
			return errors.New("empty Description field")
		}
		if post.Title=="" {
			return errors.New("empty Title field")
		}
		return nil
	}
}