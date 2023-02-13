package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetIP(c *gin.Context) string {

	h := c.Request.Header.Get("X-Forwarded-For")

	// logrus.Printf("header: %v ip: %v", c.Request.Header, strings.Split(h, ""))

	if h != "" {
		hh := strings.Split(h, ",")
		if len(hh) > 0 {
			return hh[0]
		}
	}

	return "0.0.0.0"

}
