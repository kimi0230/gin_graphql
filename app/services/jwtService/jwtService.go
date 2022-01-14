package jwtservice

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtSecret = []byte("secret")

func GenerateToken(claims Claims) (string, error) {
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", nil
	}

	return token, nil
}

func VerifyJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header_auth := c.GetHeader("authorization")
		token := strings.Split(header_auth, "Bearer ")[1]

		tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
			return jwtSecret, nil
		})
		if err != nil {
			var message string
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					message = "token is malformed"
				} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
					message = "token could not be verified because of signing problems"
				} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
					message = "signature validation failed"
				} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
					message = "EXP validation failed"
				} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
					message = "NBF validation failed"
				} else {
					message = fmt.Sprintf("JWT Token ValidationError : %d", ve.Errors)
				}
			}
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": message,
			})
			c.Abort()
			return
		}

		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			fmt.Println("email:", claims.Email)
			fmt.Println("role:", claims.Role)
			c.Set("email", claims.Email)
			c.Set("role", claims.Role)
			c.Next()
		} else {
			c.Abort()
			return
		}

	}
}
