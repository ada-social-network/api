package handler

import (
	"encoding/json"
	"testing"

	"github.com/ada-social-network/api/models"
	commonTesting "github.com/ada-social-network/api/testing"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestListCommentHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.Comment{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.Comment{})

	ListComment(db)(ctx)

	got := &[]models.Comment{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListComment response should not be empty")
	}
}

func TestCreateComment(t *testing.T) {
	type args struct {
		comment models.Comment
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
			name: "valid comment",
			args: args{
				comment: models.Comment{Content: "lorem ipsum", UserID: 1},
			},
			want: want{
				count:      1,
				statusCode: 200,
			},
		},
		{
			name: "invalid comment",
			args: args{
				comment: models.Comment{Content: "l"},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "empty comment",
			args: args{
				comment: models.Comment{},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := commonTesting.InitDB(&models.Comment{})
			res, ctx, _ := commonTesting.InitHTTPTest()

			commonTesting.AddRequestWithBodyToContext(ctx, tt.args.comment)

			CreateComment(db)(ctx)

			comment := &models.Comment{}
			_ = json.Unmarshal(res.Body.Bytes(), comment)

			if res.Code != tt.want.statusCode {
				t.Errorf("CreateComment want:%d, got:%d", tt.want.statusCode, res.Code)
			}

			tx := db.First(&models.Comment{}, comment.UserID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreateComment want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeleteComment(t *testing.T) {
	db := commonTesting.InitDB(&models.Comment{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.Comment{
		Model:   gorm.Model{},
		Content: "123",
	})

	ctx.Params = gin.Params{
		{
			Key:   "id",
			Value: "123",
		},
	}

	DeleteComment(db)(ctx)

	if res.Code != 204 {
		t.Errorf("DeleteComment want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&models.Comment{}, 123)
	if tx.RowsAffected != 0 {
		t.Errorf("DeleteComment Comment should be deleted")
	}
}

func TestGetComment(t *testing.T) {
	db := commonTesting.InitDB(&models.Comment{})

	type args struct {
		comment *models.Comment
		params  gin.Params
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
				comment: &models.Comment{
					Model: gorm.Model{
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
				comment: &models.Comment{
					Model:   gorm.Model{},
					Content: "125",
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

			db.Create(tt.args.comment)

			ctx.Params = tt.args.params

			GetComment(db)(ctx)

			if res.Code != tt.want.code {
				t.Errorf("GetCommentHandler want:%d, got:%d", tt.want.code, res.Code)
			}
		})
	}
}
