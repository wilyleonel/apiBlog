package service

import (
	"apiBlog/database"
	"apiBlog/model"

)

func CreateComment(comment *model.Comment,err error)(error){
	database.DB.Create(comment)
	return nil
}

func UpdateComment(comment model.Comment,id int)(err error){
	database.DB.Where("id=?",id).Updates(&comment)
	return nil
}

func DeleteComment(id int)(model.Comment,error){
	var comment model.Comment
	database.DB.First(&comment,id)
	database.DB.Unscoped().Delete(&comment)
	return comment,nil
}