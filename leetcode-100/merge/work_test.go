package merge

import "testing"

func TestWork(t *testing.T) {
	merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
}
