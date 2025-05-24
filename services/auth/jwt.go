package auth

import (
	// "context"
	// "fmt"
	// "log"
	// "net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luisfucros/go-api-tutorial/configs"
	// "github.com/luisfucros/go-api-tutorial/types"
	// "github.com/luisfucros/go-api-tutorial/utils"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}
