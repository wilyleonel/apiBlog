package pass

import (
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Token1 struct {
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

var secretkey string = "secretkeyjwt"

type Error struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}

func SetError(err Error, message string) Error {
	err.IsError = true
	err.Message = message
	return err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Token(email string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

// func Authorized(handler http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Header["Token"] == nil {
// 			var err Error
// 			err = SetError(err, "not token found")
// 			json.NewEncoder(w).Encode(err)
// 			return
// 		}
// 		// var mySigningKey = []byte(secretkey)
// 		// _, err := jwt.Parse(r.Header["Token"][1], func(token *jwt.Token) (interface{}, error) {
// 		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 		// 		return nil, fmt.Errorf("There was an error in parsing token.")
// 		// 	}
// 		// 	return mySigningKey, nil
// 		// })
// 		// if err != nil {
// 		// 	var err Error
// 		// 	err = SetError(err, "Your Token has been expired.")
// 		// 	json.NewEncoder(w).Encode(err)
// 		// 	return
// 		// }
// 		// var reserr Error
// 		// reserr = SetError(reserr, "Not Authorized.")
// 		// json.NewEncoder(w).Encode(err)
// 	}
// }
