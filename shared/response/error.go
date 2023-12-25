package response

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
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
