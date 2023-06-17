package jam

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	stdhttp "github.com/wpliam/common-wrap/http"
)

type testImpl struct {
}

func (t *testImpl) Server() http.Handler {
	r := gin.Default()
	r.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "success"})
	})
	return r
}

func TestNew(t *testing.T) {
	stdhttp.Register("admin.svr", &testImpl{})
	s := New()
	s.Run()
}
