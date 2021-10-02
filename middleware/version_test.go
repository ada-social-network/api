package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestVersion(t *testing.T) {
	want := "1.0.0"
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	middleware := Version(want)
	middleware(c)

	got := w.Header().Get(versionHeader)
	if got != want {
		t.Errorf("Version got:%s, want:%s", got, want)
	}
}
