package handler

import (
	"encoding/json"
	"testing"

	"github.com/ada-social-network/api/models"
	commonTesting "github.com/ada-social-network/api/testing"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TestListUserHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.User{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.User{})

	ListUserHandler(db)(ctx)

	got := &[]models.User{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListUserHandler response should not be empty")
	}
}

func TestCreateUserHandler(t *testing.T) {
	type args struct {
		user models.User
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
			name: "valid user",
			args: args{
				user: models.User{LastName: "Vedrenne", FirstName: "Alice", Email: "alice@gmail.com", DateOfBirth: "01/01/2020"},
			},
			want: want{
				count:      1,
				statusCode: 200,
			},
		},
		{
			name: "invalid user",
			args: args{
				user: models.User{LastName: "A", FirstName: "Alice", Email: "alice@gmail.com", DateOfBirth: "01/01/2020"},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
		{
			name: "empty user",
			args: args{
				user: models.User{},
			},
			want: want{
				count:      0,
				statusCode: 500,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := commonTesting.InitDB(&models.User{})
			res, ctx, _ := commonTesting.InitHTTPTest()

			commonTesting.AddRequestWithBodyToContext(ctx, tt.args.user)

			CreateUserHandler(db)(ctx)

			user := &models.User{}
			_ = json.Unmarshal(res.Body.Bytes(), user)

			if res.Code != tt.want.statusCode {
				t.Errorf("CreateUserHandler want:%d, got:%d", tt.want.statusCode, res.Code)
			}

			tx := db.First(&models.User{}, user.ID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreateUserHandler want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeleteUserHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.User{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.User{
		Model: gorm.Model{
			ID: 123,
		},
	})

	ctx.Params = gin.Params{
		{
			Key:   "id",
			Value: "123",
		},
	}

	DeleteUserHandler(db)(ctx)

	if res.Code != 204 {
		t.Errorf("DeleteUserHandler want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&models.User{}, 123)
	if tx.RowsAffected != 0 {
		t.Errorf("DeleteUserHandler User should be deleted")
	}
}
