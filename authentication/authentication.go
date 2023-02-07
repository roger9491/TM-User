package authentication

import (
	"tm-user/model/user"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var Secret []byte

// JWT function

type authClaims struct {
	jwt.StandardClaims
	UserID int64 `json:"userId"`
}

// generateToken 產生令牌
func GenerateToken(userInfo user.User) (tokenString string, err error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userInfo.UserName,
			ExpiresAt: expiresAt,
		},
		UserID: userInfo.ID,
	})
	tokenString, err = token.SignedString(Secret)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(tokenString)
	return
}

// 驗證 jwt
func AuthRequired(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	fmt.Println("asda",token)
	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &authClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
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
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if _, ok := tokenClaims.Claims.(*authClaims); ok && tokenClaims.Valid {
		c.Next()
	} else {
		c.Abort()
		return
	}
}
