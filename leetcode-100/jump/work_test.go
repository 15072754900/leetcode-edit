package jump

import "testing"

func TestWork(t *testing.T) {
	jump([]int{2, 3, 1, 1, 4})
	jump([]int{2, 3, 0, 1, 4})
}
