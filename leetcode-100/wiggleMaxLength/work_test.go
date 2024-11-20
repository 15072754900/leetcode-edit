package wiggleMaxLength

import "testing"

func TestWork(t *testing.T) {
	wiggleMaxLength([]int{1, 7, 4, 9, 2, 5})
	wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8})
	wiggleMaxLength([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
