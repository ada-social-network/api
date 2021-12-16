package handler

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/ada-social-network/api/middleware"
	"github.com/ada-social-network/api/models"
)

func TestGetCurrentUser(t *testing.T) {
	type args struct {
		user interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "nominal",
			args: args{user: &models.User{}},
			want: &models.User{},
		},
		{
			name:    "wrong type",
			args:    args{user: struct{}{}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "nil reference",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gin.Context{}
			c.Set(middleware.IdentityKey, tt.args.user)

			got, err := GetCurrentUser(c)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetCurrentUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCurrentUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
