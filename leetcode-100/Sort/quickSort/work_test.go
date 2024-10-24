package quickSort

import (
	"testing"
)

func TestSort(t *testing.T) {
	partition([]int{1, 4, 5, 7, 3, 2, 0, 3}, 0, 7)

	quickSort([]int{1, 4, 5, 7, 3, 2, 0, 3}, 0, 7)
}
