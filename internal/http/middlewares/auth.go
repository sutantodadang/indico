package middlewares

import (
	"indico/internal/constants"
	"indico/internal/repositories"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Middleware) Auth(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		headerSplit := strings.Split(authHeader, " ")
		if len(headerSplit) != 2 {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		if headerSplit[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		parseToken, err := jwt.Parse(headerSplit[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				c.JSON(401, gin.H{"error": "Invalid token"})
				c.Abort()
				return nil, nil
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !parseToken.Valid {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims, ok := parseToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if len(roles) > 0 {

			user, err := r.repo.SelectOneUserById(c, pgtype.UUID{Bytes: uuid.MustParse(claims["sub"].(string)), Valid: true})
			if err != nil {
				c.JSON(401, gin.H{"error": "Invalid User"})
				c.Abort()
				return
			}

			isAuth := false
			for _, v := range roles {

				if user.UniqueName == repositories.UserRole(v) {
					isAuth = true
					break
				}

			}

			if !isAuth {
				c.JSON(401, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}

		}

		exp := claims["exp"].(float64)
		if exp < float64(time.Now().Unix()) {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		c.Set(constants.USER_ID, claims["sub"])

		c.Next()
	}
}
