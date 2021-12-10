package handler

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"

	"github.com/ada-social-network/api/models"
	commonTesting "github.com/ada-social-network/api/testing"
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
				comment: models.Comment{
					Base:      models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
					Content:   "lorem ipsum",
					BdapostID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b"),
					UserID:    uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b"),
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

			tx := db.First(&models.Comment{}, "id = ?", comment.ID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreateComment want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeleteComment(t *testing.T) {
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
			Key:   "id",
			Value: id.String(),
		},
	}

	DeleteComment(db)(ctx)

	if res.Code != 204 {
		t.Errorf("DeleteComment want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&models.Comment{}, "id = ?", 123)
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
					Base: models.Base{
						ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b"),
					},
				},
				params: gin.Params{
					{
						Key:   "id",
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
					Content: "125",
				},
				params: gin.Params{
					{
						Key:   "id",
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

			GetComment(db)(ctx)

			if res.Code != tt.want.code {
				t.Errorf("GetCommentHandler want:%d, got:%d", tt.want.code, res.Code)
			}
		})
	}
}
