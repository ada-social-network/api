package error

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHttpError(t *testing.T) {
	type args struct {
		msg string
		err error
	}

	type fields struct {
		mode string
	}

	tests := []struct {
		name   string
		args   args
		fields fields
		want   Content
	}{
		{
			name: "mode release",
			args: args{
				msg: "internal error",
				err: errors.New("cannot connect to database"),
			},
			want: Content{
				Message: "internal error",
			},
			fields: fields{
				mode: gin.ReleaseMode,
			},
		},
		{
			name: "mode debug",
			args: args{
				msg: "internal error",
				err: errors.New("cannot connect to database"),
			},
			want: Content{
				Message: "internal error",
				Error:   "cannot connect to database",
			},
			fields: fields{
				mode: gin.DebugMode,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(tt.fields.mode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			code := http.StatusInternalServerError

			HTTPError(c, code, tt.args.msg, tt.args.err)

			var got Content

			_ = json.Unmarshal(w.Body.Bytes(), &got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPError content got:%s, want:%s", got, tt.want)
			}

			if code != w.Code {
				t.Errorf("HTTPError code got:%d, want:%d", w.Code, code)
			}
		})
	}
}

func TestInternalError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	code := http.StatusInternalServerError
	want := Content{
		Message: "Internal error",
	}

	Internal(c, errors.New("some error"))

	var got Content

	_ = json.Unmarshal(w.Body.Bytes(), &got)

	if got.Message != want.Message {
		t.Errorf("HTTPError content got:%s, want:%s", got.Message, want.Message)
	}

	if code != w.Code {
		t.Errorf("HTTPError code got:%d, want:%d", w.Code, code)
	}
}
