//nolint:depguard,errcheck
package slices_test

import (
	"testing"

	"github.com/dndev-xx/go-os-algo-courses/internal/basic-module/slices"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDArray_Add(t *testing.T) {
	tests := []struct {
		name    string
		item    int
		index   int
		wantErr bool
	}{
		{
			name:    "test01_scenario_add_get_remove",
			item:    1,
			index:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := slices.New[int](0, 5, slices.AddOne)

			gotErr := da.Add(tt.item, tt.index)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Add() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Add() succeeded unexpectedly")
			}

			require.NoError(t, gotErr)
			rsl, err := da.Get(tt.index)
			require.NoError(t, err)
			assert.Equal(t, 1, rsl)
			assert.Equal(t, 1, da.Len())

			err = da.Remove(tt.index)
			require.NoError(t, err)

			rsl, err = da.Get(tt.index)
			require.Error(t, err)
			assert.Equal(t, 0, da.Len())
			assert.Equal(t, 0, rsl)

			require.NoError(t, da.Close())
		})
	}
}

func TestDArray_AddMultipleWithinCapacity(t *testing.T) {
	da := slices.New[int](0, 5, slices.AddOne)
	defer da.Close()

	expectedValues := []int{10, 20, 30, 40, 50}
	for i, val := range expectedValues {
		err := da.Add(val, i)
		require.NoError(t, err, "failed to add element at index %d", i)
	}

	assert.Equal(t, 5, da.Len(), "length should be 5")
	assert.Equal(t, 5, da.Cap(), "capacity should still be 5 (no grow needed)")

	for i, expected := range expectedValues {
		val, err := da.Get(i)
		require.NoError(t, err, "failed to get element at index %d", i)
		assert.Equal(t, expected, val, "value at index %d mismatch", i)
	}

	for i := 4; i >= 0; i-- {
		err := da.Remove(i)
		require.NoError(t, err, "failed to remove element at index %d", i)

		assert.Equal(t, i, da.Len(), "length should be %d after removal", i)

		_, err = da.Get(i)
		require.Error(t, err, "element at index %d should be gone", i)
	}

	assert.Equal(t, 0, da.Len(), "length should be 0")
	assert.Equal(t, 5, da.Cap(), "capacity should remain 5")
}

func TestDArray_GrowWithDifferentStrategies(t *testing.T) {
	testCases := []struct {
		name           string
		growthStrategy slices.LoadType
		initialCap     int
		itemsToAdd     int
		expectedCap    int
	}{
		{
			name:           "AddOne strategy - grow by 1 each time",
			growthStrategy: slices.AddOne,
			initialCap:     3,
			itemsToAdd:     5,
			expectedCap:    5,
		},
		{
			name:           "AddHundred strategy - grow by 100",
			growthStrategy: slices.AddHundred,
			initialCap:     3,
			itemsToAdd:     150,
			expectedCap:    203,
		},
		{
			name:           "MultiplyByTwo strategy - exponential growth",
			growthStrategy: slices.MultiplyByTwo,
			initialCap:     3,
			itemsToAdd:     10,
			expectedCap:    12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			da := slices.New[int](0, tc.initialCap, tc.growthStrategy)
			defer da.Close()

			for i := 0; i < tc.itemsToAdd; i++ {
				err := da.Add(i*10, i)
				require.NoError(t, err, "failed to add element at index %d", i)
			}

			assert.Equal(t, tc.itemsToAdd, da.Len(), "length should match number of items added")
			assert.Equal(t, tc.expectedCap, da.Cap(), "capacity should have grown according to strategy")

			for i := 0; i < tc.itemsToAdd; i++ {
				val, err := da.Get(i)
				require.NoError(t, err, "failed to get element at index %d", i)
				assert.Equal(t, i*10, val, "value at index %d mismatch", i)
			}

			for i := tc.itemsToAdd - 1; i >= 0; i-- {
				err := da.Remove(i)
				require.NoError(t, err, "failed to remove element at index %d", i)
			}

			assert.Equal(t, 0, da.Len(), "length should be 0 after removing all items")
			assert.Equal(t, tc.expectedCap, da.Cap(), "capacity should remain after removals")
		})
	}
}

func TestDArray_GrowBoundaries(t *testing.T) {
	t.Run("Grow from zero capacity", func(t *testing.T) {
		da := slices.New[int](0, 0, slices.MultiplyByTwo)
		defer da.Close()

		err := da.Add(100, 0)
		require.NoError(t, err, "should be able to add to zero capacity array")

		assert.Equal(t, 1, da.Len(), "length should be 1")
		assert.GreaterOrEqual(t, da.Cap(), 1, "capacity should be at least 1 after first add")

		val, err := da.Get(0)
		require.NoError(t, err)
		assert.Equal(t, 100, val)
	})

	t.Run("Multiple grows with AddOne strategy", func(t *testing.T) {
		da := slices.New[int](0, 2, slices.AddOne)
		defer da.Close()

		for i := range 2 {
			err := da.Add(i, i)
			require.NoError(t, err)
		}
		assert.Equal(t, 2, da.Cap(), "initial capacity should be 2")

		err := da.Add(2, 2)
		require.NoError(t, err)
		assert.Equal(t, 3, da.Cap(), "capacity should grow to 3")

		err = da.Add(3, 3)
		require.NoError(t, err)
		assert.Equal(t, 4, da.Cap(), "capacity should grow to 4")

		for i := range 4 {
			val, err := da.Get(i)
			require.NoError(t, err)
			assert.Equal(t, i, val)
		}
	})
}

func TestDArray_ErrorCases(t *testing.T) {
	da := slices.New[int](0, 5, slices.MultiplyByTwo)
	defer da.Close()

	for i := range 3 {
		err := da.Add(i*10, i)
		require.NoError(t, err)
	}

	t.Run("Get with invalid index", func(t *testing.T) {
		_, err := da.Get(-1)
		require.Error(t, err, "should error on negative index")

		_, err = da.Get(5)
		assert.Error(t, err, "should error on index beyond length")
	})

	t.Run("Remove with invalid index", func(t *testing.T) {
		err := da.Remove(-1)
		require.Error(t, err, "should error on negative index")

		err = da.Remove(5)
		assert.Error(t, err, "should error on index beyond length")
	})

	t.Run("Add with invalid index", func(t *testing.T) {
		err := da.Add(100, -1)
		require.Error(t, err, "should error on negative index")

		err = da.Add(100, 10)
		assert.Error(t, err, "should error on index > length")
	})
}
