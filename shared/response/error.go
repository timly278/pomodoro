package response

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	RESPONSE_CONTEXT_KEYWORD = "http response"
)

func ErrorResponse(ctx *gin.Context, err error) gin.H {
	ctx.Set(RESPONSE_CONTEXT_KEYWORD, err)
	if strings.Contains(err.Error(), "\n") {
		errs := strings.SplitAfter(err.Error(), "\n")
		maping := make(map[string]any)
		for i, e := range errs {
			maping[fmt.Sprintf("error_%d", i+1)] = e
		}
		return maping
	} else {
		return gin.H{"error": err.Error()}
	}
}
