package service

import (
	"apiBlog/database"
	"apiBlog/model"
)

func GetCategory() ([]model.Category, error) {
	var Category []model.Category
	// var count int64
	database.DB.Find(&Category)
	// database.DB.Model("posts").Association()
	return Category, nil
}

func CreateCategory(category *model.Category,err error)error{
	database.DB.Create(category)
	return nil
}

func DeleteCategory(id int)(model.Category,error){
	var category model.Category
	database.DB.First(&category,id)
	database.DB.Scopes().Delete(&category,id)
	return category,nil
}