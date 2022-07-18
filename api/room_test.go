package api

import (
	"encoding/json"
	"github.com/go-playground/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShowRoom(t *testing.T) {
	tests := []struct {
		name   string
		expect string
	}{
		{"good case", "null"},
	}

	r := SetUpRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/show", nil)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var resp = make(map[string]string)
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			if err != nil {
				log.Println(err)
			}
			assert.Equal(t, tt.expect, resp["room_name"])
		})
	}
}

func TestShowChess(t *testing.T) {
	tests := []struct {
		name   string
		param  string
		expect string
	}{
		{"good case", "hardews", "no this room"},
	}

	r := SetUpRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/chess/"+tt.param, nil)

			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code)

			var resp = make(map[string]string)
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			if err != nil {
				log.Println(err)
			}
			assert.Equal(t, tt.expect, resp["msg"])
		})
	}
}
