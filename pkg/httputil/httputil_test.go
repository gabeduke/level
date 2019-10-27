package httputil

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestNewErrorCode(t *testing.T) {
	const StatusCode = 404

	err := fmt.Errorf("access denied: %v", 400)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	// test function
	NewError(ctx, StatusCode, err)

	//validate status code is loaded into ctx
	assert.Equal(t, ctx.Writer.Status(), StatusCode)
}
