package quickSort

import (
	"testing"
)

func TestSort(t *testing.T) {
	partition([]int{1, 4, 5, 7, 3, 2, 0, 3}, 0, 7)

	quickSort([]int{1, 4, 5, 7, 3, 2, 0, 3}, 0, 7)

	quickSort([]int{70, 746, 198, 121, 303, 53, 376, 62, 63, 64, 300, 304, 91, 96, 309, 377, 279, 357, 322, 132, 338, 486, 877, 55, 647, 5, 712, 72}, 0, 27)
}
