package entity

import (
	"expert_systems_api/infra/config"
	"expert_systems_api/pkg/exception"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (u *User) ValidateToken(bearerToken string) exception.Exception {

	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return exception.NewUnauthenticatedError("invalid token")
	}

	splitToken := strings.Fields(bearerToken)

	if len(splitToken) != 2 {
		log.Println("salah disini")
		return exception.NewUnauthenticatedError("invalid token")
	}

	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.NewUnauthenticatedError("invalid token")
		}
		return []byte(config.AppConfig().JWTSecretKey), nil
	})

	if err != nil {
		return exception.NewUnauthenticatedError("invalid token")
	}

	mapClaims := jwt.MapClaims{}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return exception.NewUnauthenticatedError("invalid token")
	} else {
		mapClaims = claims
	}

	email, ok := mapClaims["email"].(string)

	if !ok {
		return exception.NewUnauthenticatedError("invalid token")
	}

	u.Email = email

	role, ok := mapClaims["role"].(string)

	if !ok {
		return exception.NewUnauthenticatedError("invalid token")
	}

	u.Role = role

	expiredAt, ok := mapClaims["expired_at"].(float64)

	if !ok {
		return exception.NewUnauthenticatedError("invalid token")
	}

	expirationTime := time.UnixMilli(int64(expiredAt))

	if time.Now().After(expirationTime) {
		return exception.NewUnauthenticatedError("token is expired")
	}

	return nil
}

func (u *User) GenerateTokenString() string {

	claims := jwt.MapClaims{
		"email":      u.Email,
		"role":       u.Role,
		"expired_at": time.Now().Add(8 * time.Hour).UnixMilli(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(config.AppConfig().JWTSecretKey))

	return tokenString
}

func (u *User) CompareHashPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GenerateHashPassword() {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hashPassword)
}
