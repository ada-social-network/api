package handler

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/ada-social-network/api/middleware"
	"github.com/ada-social-network/api/models"
	commonTesting "github.com/ada-social-network/api/testing"
)

func TestListBdaPostComments(t *testing.T) {
	db := commonTesting.InitDB(&models.Comment{})
	res, ctx, _ := commonTesting.InitHTTPTest()
	ctx.Params = []gin.Param{{
		Key:   "id",
		Value: "80a08d36-cfea-4898-aee3-6902fa562f1d"},
	}

	tx := db.Create(&models.Comment{BdaPostID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")})
	if tx.Error != nil {
		t.Error("ListBdaPostComments response should not have an error: %w", tx.Error)
	}

	NewComment(db).ListBdaPostComments(ctx)

	got := &[]models.Comment{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListBdaPostComments response should not be empty")
	}
}

func TestCreateBdaPostComment(t *testing.T) {
	type args struct {
		comment models.Comment
		user    *models.User
		bdaPost *models.BdaPost
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
				user:    &models.User{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				bdaPost: &models.BdaPost{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				comment: models.Comment{
					Base:    models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
					Content: "lorem ipsum",
				},
			},
			want: want{
				count:      1,
				statusCode: 200,
			},
		},
		{
			name: "invalid comment",
			args: args{
				user:    &models.User{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				bdaPost: &models.BdaPost{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
				comment: models.Comment{Content: "l"},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "missing user",
			args: args{
				comment: models.Comment{Content: "lorem ipsum"},
				bdaPost: &models.BdaPost{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "empty comment",
			args: args{
				user:    &models.User{Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f1d")}},
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

			ctx.Set(middleware.IdentityKey, tt.args.user)
			ctx.Params = []gin.Param{{
				Key:   "id",
				Value: "80a08d36-cfea-4898-aee3-6902fa562f1d"},
			}

			NewComment(db).CreateBdaPostComment(ctx)

			comment := &models.Comment{}
			_ = json.Unmarshal(res.Body.Bytes(), comment)

			if res.Code != tt.want.statusCode {
				t.Errorf("CreateComment want:%d, got:%d", tt.want.statusCode, res.Code)
			}

			tx := db.First(&models.Comment{}, "id = ?", comment.ID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreateComment want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeleteBdaPostComment(t *testing.T) {
	db := commonTesting.InitDB(&models.Comment{})
	res, ctx, _ := commonTesting.InitHTTPTest()
	id := uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")

	comment := &models.Comment{
		Base:    models.Base{ID: id},
		Content: "lorem ipsum",
	}

	db.Create(comment)

	ctx.Params = gin.Params{
		{
			Key:   "commentId",
			Value: id.String(),
		},
	}

	NewComment(db).DeleteBdaPostComment(ctx)

	if res.Code != 204 {
		t.Errorf("DeleteComment want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&models.Comment{}, "id = ?", id)
	if tx.RowsAffected != 0 {
		t.Errorf("DeleteComment Comment should be deleted")
	}
}

func TestGetBdaPostComment(t *testing.T) {
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
					Base: models.Base{
						ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b"),
					},
					Content: "Lorem ipsum",
				},
				params: gin.Params{
					{
						Key:   "commentId",
						Value: "80a08d36-cfea-4898-aee3-6902fa562f0b",
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
					Base:    models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
					Content: "Lorem ipsum",
				},
				params: gin.Params{
					{
						Key:   "commentId",
						Value: "99999999-9999-9999-9999-999999999999",
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

			NewComment(db).GetBdaPostComment(ctx)

			if res.Code != tt.want.code {
				t.Errorf("GetCommentHandler want:%d, got:%d", tt.want.code, res.Code)
			}
		})
	}
}
