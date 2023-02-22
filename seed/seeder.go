package seed

import (
	"apiBlog/database"
	"apiBlog/model"
	"fmt"
)

func User() {
	user := model.User{Email: "wilyleodan12hancco@gmail.com", Password: "admin"}
	user1 := model.User{Email: "carlos@gmail.com", Password: "admin"}
	user2 := model.User{Email: "angel@gmail.com", Password: "admin"}
	user3 := model.User{Email: "luis@gmail.com", Password: "admin"}
	database.DB.Create(&user)
	database.DB.Create(&user1)
	database.DB.Create(&user2)
	database.DB.Create(&user3)
}
func Profile() {
	profile := model.Profile{Name: "wily", LastName: "hancco", Address: "las flores", Img: "image.png", Age: 26, UserId: 1}
	profile1 := model.Profile{Name: "carlos", LastName: "larico", Address: "las flores", Img: "image.png", Age: 26, UserId: 2}
	profile2 := model.Profile{Name: "angel", LastName: "caceres", Address: "las flores", Img: "image.png", Age: 26, UserId: 3}
	profile3 := model.Profile{Name: "luis", LastName: "sanchez", Address: "las flores", Img: "image.png", Age: 26, UserId: 4}
	database.DB.Create(&profile)
	database.DB.Create(&profile1)
	database.DB.Create(&profile2)
	database.DB.Create(&profile3)
}
func Post() {
	post := model.Post{Title: "La IA es realmente peligrosa?", Description: "en el siglo XXI cada la ves la tecnologia avanza a pasos agigantados un ejmplo claro es la la inteligencia artificial (IA) ", UserId: 1}
	post1 := model.Post{Title: "La gastronomia en el Peru", Description: "La gastronomia en el Peru es uno de los atrayentes del turismo,la gran varidad de platos tipicos se que puede encontra en el territorio peruano....", UserId: 2}
	database.DB.Create(&post)
	database.DB.Create(&post1)
}

func Comment() {
	comment := model.Comment{Description: "Esto es un comentario del post 1", PostId: 1,UserId: 1}
	comment2 := model.Comment{Description: "Esto es un comentario del post 1", PostId: 1,UserId: 4}
	comment3 := model.Comment{Description: "Esto es un comentario del post 1", PostId: 1,UserId: 3}
	comment1 := model.Comment{Description: "Esto es una publicacion del post 2", PostId: 2,UserId: 2}
	database.DB.Create(&comment)
	database.DB.Create(&comment1)
	database.DB.Create(&comment2)
	database.DB.Create(&comment3)
}
func Category() {
	category := model.Category{Category: "Tecnologia"}
	category1 := model.Category{Category: "Gastronomia"}
	database.DB.Create(&category)
	database.DB.Create(&category1)
}

func Like(){
	like:=model.Like{UserId: 1,PostId: 1,CommentId: 1}
	like1:=model.Like{UserId: 2,PostId: 2,CommentId: 2}
	database.DB.Create(&like)
	database.DB.Create(&like1)
}
func Seed() {
	User()
	Profile()
	Post()
	Comment()
	Category()
	Like()
	fmt.Println("las semillas fueron sembradas correctamente")
}
