package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	user  = "USER"
	admin = "ADMIN"
)

func (s *Server) AuthMiddleware(scope string, mustHave bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if mustHave && authHeader == "" {
			c.Status(http.StatusUnauthorized)
			return
		}

		if !mustHave && authHeader == "" {
			c.Next()
			return
		}

		auth, err := s.auth.GetAuth(c, authHeader)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Set("userID", auth.ID)

		if !mustHave {
			c.Next()
			return
		}

		for _, s := range auth.Scopes {
			if s == scope {
				c.Next()
				return
			}
		}

		c.Status(http.StatusForbidden)
	}
}
