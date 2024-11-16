package buildTree

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	fmt.Println(getAnsPost([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	fmt.Println(getAnsPre([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	fmt.Println(getAnsIn([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
