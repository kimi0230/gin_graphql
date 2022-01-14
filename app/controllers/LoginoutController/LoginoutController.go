package LoginoutController

import (
	jwtservice "gin_graphql/app/services/jwtService"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Login(c *gin.Context) {
	// validate request body
	var reqJSON struct {
		Email    string `json:"email" form:"email" binding:"required" `
		Password string `json:"password" form:"password" binding:"required" `
	}
	err := c.ShouldBind(&reqJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: check password
	if reqJSON.Email == "kimi" && reqJSON.Password == "123456" {
		now := time.Now()
		jwtId := reqJSON.Email + strconv.FormatInt(now.Unix(), 10)
		role := "Staff"

		// set claims and sign token
		claims := jwtservice.Claims{
			Email: reqJSON.Email,
			Role:  role,
			StandardClaims: jwt.StandardClaims{
				Audience:  reqJSON.Email,
				ExpiresAt: now.Add(60 * time.Second).Unix(),
				Id:        jwtId,
				IssuedAt:  now.Unix(),
				Issuer:    "KimiGinJWT",
				NotBefore: now.Add(3 * time.Second).Unix(),
				Subject:   reqJSON.Email,
			},
		}

		token, err := jwtservice.GenerateToken(claims)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
		return
	}
	// incorrect account or password
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "authorized ok",
	})
}
