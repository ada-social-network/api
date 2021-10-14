package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestListPostHandler(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	_ = db.AutoMigrate(&Post{})
	db.Create(&Post{})

	ListPostHandler(db)(c)

	got := &[]Post{}
	_ = json.Unmarshal(w.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListPostHandler response should not be empty")
	}
}
