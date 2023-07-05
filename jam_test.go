package jam

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	jelastic "github.com/wpliam/common-wrap/elastic"
	jgorm "github.com/wpliam/common-wrap/gorm"
	jhttp "github.com/wpliam/common-wrap/http"
	jredis "github.com/wpliam/common-wrap/redis"
)

type testHttp struct {
}

func (t *testHttp) Server() http.Handler {
	return gin.Default()
}

func TestNew(t *testing.T) {
	jhttp.Register("admin.svr", &testHttp{})
	jhttp.Register("front.svr", &testHttp{})

	s := New()
	jgorm.NewClientProxy("mysql")
	jredis.NewClientProxy("redis")
	jelastic.NewClientProxy("elastic")
	s.Run()
}
