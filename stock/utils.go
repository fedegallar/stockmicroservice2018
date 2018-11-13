package stock

import (
	"strings"

	"github.com/fedegallar/stockmicroservice2018/config/errors"
	"github.com/fedegallar/stockmicroservice2018/security"
	"github.com/gin-gonic/gin"
)

func getTokenHeader(c *gin.Context) (string, error) {
	tokenString := c.GetHeader("Authorization")
	if strings.Index(tokenString, "bearer") != 0 {
		return "", errors.Unauthorized
	}
	return tokenString[:7], nil
}

func validateAuthentication(c *gin.Context) error {
	tokenString, err := getTokenHeader(c)
	if err != nil {
		return err
	}
	if _, err = security.Validate(tokenString); err != nil {
		return err
	}
	return nil
}
