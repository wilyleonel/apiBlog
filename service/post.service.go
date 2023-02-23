package service

import (
	"apiBlog/database"
	"apiBlog/model"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func GetAllPost(limit1,offset1 int) ([]model.Post, error) {
	var post []model.Post
	database.DB.Offset(offset1).Limit(limit1).Order("id desc").Table("posts").Scan(&post)
	return post, nil
}
// func GetAllPost() ([]model.Post, error) {
// 	var post []model.Post
// 	database.DB.Order("id desc").Table("posts").Scan(&post)
// 	return post, nil
// }

func GetPost(id int) (model.Post, error) {
	var post model.Post
	database.DB.Table("posts").First(&post, id)
	database.DB.Model(&post).Association("Comment").Find(&post.Comment)
	database.DB.Model(&post).Association("Categories").Find(&post.Categories)
	database.DB.Model(&post).Association("Like").Find(&post.Like)
	return post, nil
}

func CreatePost(post *model.Post,err error) error {
	database.DB.Create(post)
	return nil
}

func UpdatePost(post model.Post, id int) (err error) {
	database.DB.Where("id=?", id).Updates(&post)
	return nil
}

func DeletePost(id int)(model.Post,error){
	var post model.Post
	database.DB.First(&post,id)
	database.DB.Unscoped().Delete(&post)
	return post,nil
}

func Pagination(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}