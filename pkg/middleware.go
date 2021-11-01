package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := c.Cookie("session")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Set("userId", session)
		c.Set("is_Logged_in", true)
		c.Next()
	}
}
