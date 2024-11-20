package canCompleteCircuit

import "testing"

func TestComplete(t *testing.T) {
	// canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	notToMuch([]int{2, 3, 4}, []int{3, 4, 3})
}
