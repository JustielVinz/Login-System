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

// JWTClaims represents the claims stored in the JWT

type Identification struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SecretKey struct {
	SecretKey string `json:"secret_key"`
}
