package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestListPostHandler(t *testing.T) {
	db, res, ctx := InitHTTPTest(&Post{})

	db.Create(&Post{})

	ListPostHandler(db)(ctx)

	got := &[]Post{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListPostHandler response should not be empty")
	}
}

func TestCreatePostHandler(t *testing.T) {
	type args struct {
		post Post
	}

	type want struct {
		count      int64
		statusCode int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "valid post",
			args: args{
				post: Post{Content: "lorem ipsum"},
			},
			want: want{
				count:      1,
				statusCode: 200,
			},
		},
		{
			name: "invalid post",
			args: args{
				post: Post{Content: "l"},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "empty post",
			args: args{
				post: Post{},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, res, ctx := InitHTTPTest(&Post{})

			AddRequestWithBodyToContext(ctx, tt.args.post)

			CreatePostHandler(db)(ctx)

			post := &Post{}
			_ = json.Unmarshal(res.Body.Bytes(), post)

			if res.Code != tt.want.statusCode {
				t.Errorf("CreatePostHandler want:%d, got:%d", tt.want.statusCode, res.Code)
			}

			tx := db.First(&Post{}, post.ID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreatePostHandler want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeletePostHandler(t *testing.T) {
	db, res, ctx := InitHTTPTest(&Post{})

	db.Create(&Post{
		CommonResource: CommonResource{
			ID: 123,
		},
	})

	ctx.Params = gin.Params{
		{
			Key:   "id",
			Value: "123",
		},
	}

	DeletePostHandler(db)(ctx)

	if res.Code != 204 {
		t.Errorf("DeletePostHandler want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&Post{}, 123)
	if tx.RowsAffected != 0 {
		t.Errorf("DeletePostHandler Post should be deleted")
	}
}

func InitDB(i interface{}) *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	_ = db.AutoMigrate(i)
	return db
}

func AddRequestWithBodyToContext(c *gin.Context, i interface{}) {
	marshal, _ := json.Marshal(i)
	body := strings.NewReader(string(marshal))
	c.Request, _ = http.NewRequest("", "", body)
}

func InitHTTPTest(i interface{}) (db *gorm.DB, res *httptest.ResponseRecorder, ctx *gin.Context) {
	db = InitDB(i)
	res = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(res)

	return db, res, ctx
}
