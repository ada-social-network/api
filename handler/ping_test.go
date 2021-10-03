package handler

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	PingHandler(c)

	if w.Code != 200 {
		t.Errorf("PingHandler want:%d, got:%d", 200, w.Code)
	}

	want := Pong{Message: "pong"}
	got := &Pong{}
	_ = json.Unmarshal(w.Body.Bytes(), got)

	if !reflect.DeepEqual(want, *got) {
		t.Errorf("Pinghandler want:%s, got:%s", want, *got)
	}
}
