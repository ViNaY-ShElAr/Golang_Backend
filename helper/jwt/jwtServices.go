package jwt

import (
	"GO_PROJECT/config"
	"GO_PROJECT/model"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(tokenData model.JwtData) (string, error) {
	tokenData.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.Config.JWT_TOKEN.EXPIRY_TIME)).Unix(), // Example: token expires in 24 hours
		Issuer:    config.Config.JWT_TOKEN.TOKEN_ISSUER,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenData)

	jwtToken, err := token.SignedString([]byte(config.Config.JWT_TOKEN.TOKEN_SECRET))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method in jwt token")
		}
		return []byte(config.Config.JWT_TOKEN.TOKEN_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return token, nil
}
