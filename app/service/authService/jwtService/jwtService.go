package jwtService

import (
	"github.com/MeibisuX673/LessonRedis/app/model/user"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateJwtToken(user user.User) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()
	claims["sub"] = user.Id
	claims["isActivate"] = user.IsActivate

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	return tokenStr, nil

}
