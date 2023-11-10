package security

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   string `json:"status"`
}
