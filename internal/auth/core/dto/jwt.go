package dto

import "github.com/golang-jwt/jwt/v4"

// JWTのClaims
type AccountClaims struct {
	ID        string `json:"id"`
	SessionID string `json:"sessionID"`
	jwt.RegisteredClaims
}
