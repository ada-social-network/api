package testing

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB Initialize in-memory database and auto-migrate model
func InitDB(model interface{}) *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

// AddRequestWithBodyToContext hydrate context with request
func AddRequestWithBodyToContext(c *gin.Context, body interface{}) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	bodyReader := strings.NewReader(string(bodyBytes))

	c.Request, err = http.NewRequest("", "", bodyReader)
	if err != nil {
		log.Fatal(err)
	}
}

// InitHTTPTest init HTTP test context
func InitHTTPTest() (res *httptest.ResponseRecorder, ctx *gin.Context, engine *gin.Engine) {
	gin.SetMode(gin.TestMode)

	res = httptest.NewRecorder()
	ctx, engine = gin.CreateTestContext(res)

	return res, ctx, engine
}
