package response

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"any"`
}

func ErrorMultiResponse(err error) gin.H {
	errs := strings.SplitAfter(err.Error(), "\n")
	maping := make(map[string]any)
	for i, e := range errs {
		maping[fmt.Sprintf("error_%d", i+1)] = e
	}
	return maping
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
