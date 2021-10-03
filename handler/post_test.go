package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPostHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	PostHandler(c)

	got := &[]Post{}
	_ = json.Unmarshal(w.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("PostHandler response should not be empty")
	}
}
