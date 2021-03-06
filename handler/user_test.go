package handler

import (
	"encoding/json"
	"testing"

	"github.com/ada-social-network/api/models"
	"github.com/ada-social-network/api/repository"
	commonTesting "github.com/ada-social-network/api/testing"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func TestListUserHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.User{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.User{})

	userRepository := repository.NewCommentRepository(db)
	NewUserHandler((*repository.UserRepository)(userRepository)).ListUser(ctx)

	got := &[]models.User{}
	_ = json.Unmarshal(res.Body.Bytes(), got)

	if len(*got) == 0 {
		t.Error("ListUser response should not be empty")
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

			userRepository := repository.NewCommentRepository(db)
			NewUserHandler((*repository.UserRepository)(userRepository)).CreateUser(ctx)

			user := &models.User{}
			_ = json.Unmarshal(res.Body.Bytes(), user)

			if res.Code != tt.want.statusCode {
				t.Errorf("CreateUser want:%d, got:%d", tt.want.statusCode, res.Code)
			}

			tx := db.First(&models.User{}, "id = ?", user.ID)
			if tx.RowsAffected != tt.want.count {
				t.Errorf("CreateUser want:%d, got:%d", tt.want.count, tx.RowsAffected)
			}
		})
	}
}

func TestDeleteUserHandler(t *testing.T) {
	db := commonTesting.InitDB(&models.User{})
	res, ctx, _ := commonTesting.InitHTTPTest()

	db.Create(&models.User{

		Base: models.Base{ID: uuid.FromStringOrNil("80a08d36-cfea-4898-aee3-6902fa562f0b")},
	})

	ctx.Params = gin.Params{
		{
			Key:   "id",
			Value: "80a08d36-cfea-4898-aee3-6902fa562f0b",
		},
	}

	userRepository := repository.NewCommentRepository(db)
	NewUserHandler((*repository.UserRepository)(userRepository)).DeleteUser(ctx)

	if res.Code != 204 {
		t.Errorf("DeleteUser want:%d, got:%d", 204, res.Code)
	}

	tx := db.First(&models.User{}, "id = ?", "80a08d36-cfea-4898-aee3-6902fa562f0b")
	if tx.RowsAffected != 0 {
		t.Errorf("DeleteUser User should be deleted")
	}
}
