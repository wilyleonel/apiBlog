package routes

import (
	"apiBlog/controllers"
	"apiBlog/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/login",middlewares.SetMiddlewareJSON(controllers.Login)).Methods("POST")

	r.HandleFunc("/user",middlewares.SetMiddlewareJSON(controllers.GetUser)).Methods("GET")
	r.HandleFunc("/user/{id}",middlewares.SetMiddlewareJSON(controllers.GetIdUser)).Methods("GET")
	r.HandleFunc("/user",middlewares.SetMiddlewareJSON(controllers.CreateUser)).Methods("POST")
	r.HandleFunc("/user/{id}",middlewares.SetMiddlewareJSON(controllers.UpdateUser)).Methods("PUT")
	r.HandleFunc("/user/{id}",controllers.DeleteUser).Methods("DELETE")

	//route post
	r.HandleFunc("/post",middlewares.SetMiddlewareJSON(controllers.GetAllPost)).Methods("GET")
	r.HandleFunc("/post/{id}",middlewares.SetMiddlewareJSON(controllers.GetPost)).Methods("GET")
	r.HandleFunc("/post",middlewares.SetMiddlewareJSON(controllers.CreatePost)).Methods("POST")
	r.HandleFunc("/post/{id}", controllers.UpdatePost).Methods("PUT")
	r.HandleFunc("/post/{id}", controllers.DeletePost).Methods("DELETE")

	//route comments
	r.HandleFunc("/comment",middlewares.SetMiddlewareJSON(controllers.CreateComment)).Methods("POST")
	r.HandleFunc("/comment/{id}",middlewares.SetMiddlewareJSON(controllers.UpdateComment)).Methods("PUT")
	r.HandleFunc("/comment/{id}",controllers.DeleteComment).Methods("DELETE")

	//route profile
	r.HandleFunc("/profile",middlewares.SetMiddlewareJSON(controllers.CreateProfile)).Methods("POST")
	r.HandleFunc("/profile/{id}",middlewares.SetMiddlewareJSON(controllers.UpdateProfile)).Methods("PUT")

	//route category
	r.HandleFunc("/category",middlewares.SetMiddlewareJSON(controllers.GetAllCategory)).Methods("GET")
	r.HandleFunc("/category",middlewares.SetMiddlewareJSON(controllers.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/{id}",controllers.DeleteCategory).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}

