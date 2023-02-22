package middlewares

import (
	"net/http"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		err := pass.Authorized()
// 		if err != nil {
// 			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 			return
// 		}
// 		next(w, r)
// 	}
// }
