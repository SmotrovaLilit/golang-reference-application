package slices

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	var result = Convert(
		func(i int) string { return strconv.Itoa(i) },
		[]int{1, 2, 3, 4, 5},
	)
	require.Equal(t, []string{"1", "2", "3", "4", "5"}, result)
}
