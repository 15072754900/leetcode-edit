package canJump

import "testing"

func TestWork(t *testing.T) {
	canJump([]int{2, 3, 1, 1, 4})
	canJump([]int{3, 2, 1, 0, 4})
}
