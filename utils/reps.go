package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Request          *Request
	TLS              *tls.ConnectionState
	Trailer          Header
}

func ErrorResponse(c *gin.Context, code int, msg string, err error) {

	statusCode := code

	// check if dedicated status code
	errorResponseCode := c.PostForm("errorResponseCode")
	if errorResponseCode != "" && errorResponseCode == "200" {
		s, e := strconv.Atoi(errorResponseCode)
		if e == nil {
			statusCode = s
		}
	}

	errorMsg := msg
	if err != nil {
		errorMsg = fmt.Sprintf("%s: %v", msg, err)
	}
	c.Set("ErrorMsg", errorMsg)
	c.JSON(statusCode, gin.H{"error": msg, "code": code})
}
