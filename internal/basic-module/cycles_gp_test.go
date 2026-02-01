package basicmodule_test

import (
	"strings"
	"testing"

	basicmodule "github.com/dndev-xx/go-os-algo-courses/internal/basic-module" //nolint:depguard
	"github.com/stretchr/testify/assert"                                       //nolint:depguard
	"github.com/stretchr/testify/require"                                      //nolint:depguard
)

func TestGpmagic(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		condFunc func(i, j int) bool
		err      bool
	}{
		{
			name:     "diagonal_condition",
			condFunc: basicmodule.IsDiagonalCond,
			expected: `
.########################
..#######################
...######################
....#####################
.....####################
......###################
.......##################
........#################
.........################
..........###############
...........##############
............#############
.............############
..............###########
...............##########
................#########
.................########
..................#######
...................######
....................#####
.....................####
......................###
.......................##
........................#
.........................
`,
			err: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := basicmodule.Gpmagic(tt.condFunc)

			if tt.err {
				require.Error(t, err)
				return
			}
			result = strings.TrimSpace(result)
			tt.expected = strings.TrimSpace(tt.expected)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
