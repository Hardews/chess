package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Operation(t *testing.T) {
	client1, client2 := NewClient()
	tests := []struct {
		name       string
		choose     int
		num        int
		direction  int
		step       int
		Client1    *Client
		Client2    *Client
		exceptRes1 error
		exceptRes2 bool
	}{
		{"gun move test", 1, 1, 2, 2, client1, client2, nil, true},
		{"horse move test", 2, 1, 12, 1, client1, client2, nil, true},
		{"guard move test", 3, 1, 344, 2, client1, client2, nil, true},
		{"vehicle move test", 4, 1, 3, 2, client1, client2, nil, true},
		{"minister move test", 5, 1, 12, 2, client1, client2, nil, true},
		{"commander move test", 6, 1, 1, 1, client1, client2, nil, true},
		{"soldier move test", 7, 1, 1, 1, client1, client2, nil, true},
		{"gun move test", 1, 1, 2, 1, client2, client1, nil, true},
		{"horse move test", 2, 1, 12, 1, client2, client1, nil, true},
		{"guard move test", 3, 1, 122, 1, client2, client1, nil, true},
		{"vehicle move test", 4, 1, 1, 2, client2, client1, nil, true},
		{"minister move test", 5, 1, 12, 2, client2, client1, nil, true},
		{"commander move test", 6, 1, 1, 1, client2, client1, nil, true},
		{"soldier move test", 7, 1, 1, 1, client2, client1, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.Client1.Operation(tt.choose, tt.num, tt.direction, tt.step, tt.Client2)
			want1, want2 := tt.exceptRes1, tt.exceptRes2

			assert.Equal(t, want1, got1)
			assert.Equal(t, want2, got2)
		})
	}
}

func TestIsWin(t *testing.T) {
	client1, client2 := NewClient()
	tests := []struct {
		name       string
		Client1    *Client
		Client2    *Client
		exceptRes1 bool
		exceptRes2 int
	}{
		{"is client1 win", client1, client2, false, 2},
		{"is client2 win", client2, client1, false, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := tt.Client1.isWin(tt.Client2)
			want1, want2 := tt.exceptRes1, tt.exceptRes2

			assert.Equal(t, want1, got1)
			assert.Equal(t, want2, got2)
		})
	}
}
