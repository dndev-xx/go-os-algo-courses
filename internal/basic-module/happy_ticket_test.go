package basicmodule_test

import (
	"testing"

	basicmodule "github.com/dndev-xx/go-os-algo-courses/internal/basic-module" //nolint:depguard
	"github.com/stretchr/testify/assert"                                       //nolint:depguard
	"github.com/stretchr/testify/require"                                      //nolint:depguard
)

func TestIsHappyTicket(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		ticket  int
		want    bool
		wantErr bool
	}{
		{
			name:    "simple succeeded test",
			ticket:  622325,
			want:    true,
			wantErr: false,
		},
		{
			name:    "not happy ticket",
			ticket:  123456,
			want:    false,
			wantErr: false,
		},
		{
			name:    "have error",
			ticket:  1234567,
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := basicmodule.IsHappyTicket(tt.ticket)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("IsHappyTicket() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				require.Error(t, gotErr)
				assert.False(t, got)
			}

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sym     string
		want    int
		wantErr bool
	}{
		{
			name:    "test_0",
			sym:     "1234567890123456789012345678901234567890123456789012345678901234",
			want:    64,
			wantErr: false,
		},
		{
			name:    "test_1",
			sym:     "12345678",
			want:    8,
			wantErr: false,
		},
		{
			name:    "test_2",
			sym:     "",
			want:    0,
			wantErr: false,
		},
		{
			name:    "test_3",
			sym:     "$",
			want:    1,
			wantErr: false,
		},
		{
			name:    "test_4",
			sym:     "abcdefghijklmnoprstuvwxyz",
			want:    25,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.Len(tt.sym)
			if tt.wantErr {
				t.Fatal("Len() succeeded unexpectedly")
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestNHappyTickets(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		digits int
		want   int64
	}{
		{
			name:   "test_0",
			digits: 1,
			want:   10,
		},
		{
			name:   "test_1",
			digits: 2,
			want:   670,
		},
		{
			name:   "test_2",
			digits: 3,
			want:   55252,
		},
		{
			name:   "test_3",
			digits: 4,
			want:   4816030,
		},
		{
			name:   "test_4",
			digits: 5,
			want:   432457640,
		},
		{
			name:   "test_5",
			digits: 6,
			want:   39581170420,
		},
		{
			name:   "test_6",
			digits: 7,
			want:   3671331273480,
		},
		{
			name:   "test_7",
			digits: 8,
			want:   343900019857310,
		},
		{
			name:   "test_8",
			digits: 9,
			want:   32458256583753952,
		},
		{
			name:   "test_9",
			digits: 10,
			want:   3081918923741896840,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basicmodule.NHappyTickets(tt.digits)
			assert.Equal(t, tt.want, got)
		})
	}
}
