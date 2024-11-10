package solve

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindSmallestElements(t *testing.T) {
	ts := []struct {
		name string
		vals []int
		res  []int
	}{
		{"1", []int{1, 2, 3, 4, 5}, []int{1, 2}},
		{"2", []int{5, 4, 3, 2, 1}, []int{}},
	}

	for _, v := range ts {
		t.Run(v.name, func(t *testing.T) {
			res := FindTwoSmallestElements(v.vals)
			require.Equal(t, v.res, res)
		})
	}
}
