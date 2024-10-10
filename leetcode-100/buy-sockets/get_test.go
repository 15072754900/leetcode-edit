package buy

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fmt.Println(Get([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(Get([]int{1, 2, 3, 4, 5}))
	fmt.Println(Get2([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(Get2([]int{1, 2, 3, 4, 5}))
}
