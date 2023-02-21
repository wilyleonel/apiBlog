package service

import (
	"apiBlog/database"
	"apiBlog/model"
)

func CreateProfile(profile *model.Profile,err error)error{
	database.DB.Create(profile)
	return nil
}

func UpdateProfile(profile model.Profile,id int)( err error){
	database.DB.Where("id=?",id).Updates(&profile)
	return nil
}