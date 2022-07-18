package api

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		except string
	}{
		{"good case 1", `{"username":"1225101127","password":"12345678"}`, "successful"},
		{"good case 2", `{"username":"1225101127","password":"123456"}`, "密码错误"},
		{"good case 3", `{"username":"1225101","password":"123456"}`, "无此账号"},
		{"bad  case 1", `{"username":"","password":"123456"}`, "输入的账号为空"},
		{"bad  case 2", `{"username":"1225101","password":""}`, "输入的密码为空"},
	}

	r := SetUpRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(tt.param))

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var resp = make(map[string]string)
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.except, resp["msg"])
		})
	}
}

func TestRegister(t *testing.T) {

}
