package Utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/taskManagement/Model"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte(GetEnv("SECRET_KEY", "your-secret-key"))

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hash), nil
}

func VerifyPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func GenerateJWT(user Model.User) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	jwtCliams := jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"name":  user.Name,
		"exp":   expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtCliams)
	return token.SignedString(jwtSecretKey)
}
