package handler

import (
	"encoding/json"
	"testing"

	commonTesting "github.com/ada-social-network/api/testing"
	"github.com/gin-gonic/gin"
)

func TestListPostHandler(t *testing.T) {
	db := commonTesting.InitDB(&Post{})
	res, ctx, _ := commonTesting.InitHTTPTest()

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
			db := commonTesting.InitDB(&Post{})
			res, ctx, _ := commonTesting.InitHTTPTest()

			commonTesting.AddRequestWithBodyToContext(ctx, tt.args.post)

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
	db := commonTesting.InitDB(&Post{})
	res, ctx, _ := commonTesting.InitHTTPTest()

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

func TestGetPostHandler(t *testing.T) {
	db := commonTesting.InitDB(&Post{})

	type args struct {
		post   *Post
		params gin.Params
	}

	type want struct {
		code int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "nominal",
			args: args{
				post: &Post{
					CommonResource: CommonResource{
						ID: 122,
					},
				},
				params: gin.Params{
					{
						Key:   "id",
						Value: "122",
					},
				},
			},
			want: want{
				code: 200,
			},
		},
		{
			name: "not found",
			args: args{
				post: &Post{
					CommonResource: CommonResource{
						ID: 124,
					},
				},
				params: gin.Params{
					{
						Key:   "id",
						Value: "125",
					},
				},
			},
			want: want{
				code: 404,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ctx, _ := commonTesting.InitHTTPTest()

			db.Create(tt.args.post)

			ctx.Params = tt.args.params

			GetPostHandler(db)(ctx)

			if res.Code != tt.want.code {
				t.Errorf("GetPostHandler want:%d, got:%d", tt.want.code, res.Code)
			}
		})
	}
}
