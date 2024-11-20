package subsets

import "testing"

func TestWork(t *testing.T) {
	subsets([]int{1, 2, 3})
	subsets([]int{1, 2, 2})
}
