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
