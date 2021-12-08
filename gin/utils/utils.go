package utils

import (
	"github.com/golang-jwt/jwt"
	"hb_gin/config"
)

func GetToken(token string) (*jwt.Token, error) {
	userToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Conf.JWT.SigningKey), nil
	})
	return userToken, err
}
