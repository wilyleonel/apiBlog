package model

type Like struct {
	Id        uint `gorm:"primaryKey"`
	UserId    uint
	PostId    uint
	CommentId uint
}
