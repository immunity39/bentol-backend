package services

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("e0qAfj3IxEXbHWJRQtzGuvAPm0d+O3rzyaTxFbuKeMg=")

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

type StoreClaims struct {
	StoreID uint `json:"store_id"`
	jwt.StandardClaims
}
