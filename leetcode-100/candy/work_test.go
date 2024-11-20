package candy

import "testing"

func TestWork(t *testing.T) {
	candy([]int{1, 0, 2})
	candy([]int{1, 2, 2})
	candy([]int{1, 3, 2, 2, 1})
}
