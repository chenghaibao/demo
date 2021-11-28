package testAdmin

import (
	"fmt"
	"hb_gin/test"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hedzr/assert"

	"hb_gin/route"
)

var testGin *gin.Engine

func init() {
	// https://hedzr.com/golang/testing/golang-assert-1/
	testGin = route.Routers()
}

func TestPing(t *testing.T) {
	uri := "/ping"
	body := test.Get(uri, testGin)
	fmt.Println(body.Body.String())
	assert.NotEqual(t, body.Body.String(), "")
}
