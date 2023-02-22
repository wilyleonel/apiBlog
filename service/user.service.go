package service

import (
	"apiBlog/database"
	"apiBlog/model"

	pass "apiBlog/pass"
)

type Result struct {
	Id       int
	Email    string
	Password string
}

	// func Get() ([]Result, error) {
	// 	query := `select id,email,password from users ;`
	// 	// var users []model.User
	// 	rows, err := database.DB.Raw(query).Rows()
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	defer rows.Close()
	// 	var user []Result
	// 	for rows.Next() {
	// 		var user1 Result
	// 		err := rows.Scan(
	// 			&user1.Id,
	// 			&user1.Email,
	// 			&user1.Password,
	// 		)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		user = append(user, user1)
	// 	}
	// 	return user, nil
	// }
func Get() ([]Result, error) {
	var usuario []Result
	database.DB.Offset(0).Limit(10).Table("users").Select("email", "password", "id").Scan(&usuario)
	return usuario, nil
}

func GetIdUser(id int) (model.User, error) {
	var user model.User
	database.DB.Select("email", "password", "id").First(&user, id)
	database.DB.Model(&user).Association("Profile").Find(&user.Profile)
	database.DB.Model(&user).Association("Post").Find(&user.Post)
	return user, nil
}

func Create(user *model.User, err error) error {
	user.Password, err = pass.HashPassword(user.Password)
	if err != nil {
		return err
	}
	database.DB.Create(user)
	return nil
}

func Update(user model.User, id int) (err error) {
	database.DB.Model(&user).Select("password").Updates("password")
	database.DB.Where("id=?", id).Updates(&user)
	return nil
}

func Delete(id int) (model.User, error) {
	var user model.User
	var profile model.Profile
	database.DB.Unscoped().Delete(&profile,id)
	database.DB.Unscoped().Delete(&user,id)
	return user, nil
}
