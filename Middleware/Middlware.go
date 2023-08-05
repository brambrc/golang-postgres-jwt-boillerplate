package Middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func RoleMiddleware(role string) gin.HandlerFunc {
	// return the HandlerFunc
	return func(c *gin.Context) {
		// get the token from the header
		tokenString := c.Request.Header.Get("Authorization")
		// remove the Bearer prefix
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		// parse the token

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

		// if there is an error, the token must have expired

		if err != nil {
			// return a forbidden status
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// validate the token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// check if the role matches
			if claims["role"] == "Admin" {
				// set the user id to the context

				c.Set("user_id", claims["id"])
				// continue with the request
				c.Next()
			} else {
				// return a forbidden status
				c.AbortWithStatus(http.StatusForbidden)

			}
		} else {
			// return a forbidden status
			c.AbortWithStatus(http.StatusForbidden)
		}

	}

}

func AdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("Admin")
}
