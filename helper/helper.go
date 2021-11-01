package helper

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func HashPassword(str string) string {
	password, err := bcrypt.GenerateFromPassword([]byte(str), 10)
	if err != nil {
		panic(err)
	}
	return string(password)
}
func IsValidEmail(str string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(str)
}

func ComparePassword(hashpassword, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(str))
	return err == nil
}

func Length(str string) bool {
	return len(strings.Trim(str, " ")) == 0
}
func MiddleWare(c *gin.Context) {

}
