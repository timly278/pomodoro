package api

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
	Data    any    `json:"any"`
}

func errorResponse(err error) gin.H {
	errs := strings.SplitAfter(err.Error(), "\n")
	maping := make(map[string]any)
	for i, e := range errs {
		maping[fmt.Sprintf("error_%d", i+1)] = e
	}
	return maping
}
