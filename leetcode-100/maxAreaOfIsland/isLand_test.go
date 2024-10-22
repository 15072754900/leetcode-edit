package maxAreaOfIsland

import (
	"fmt"
	"testing"
)

func TestMaxAreaOfIsland(t *testing.T) {
	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 进行搜索的区间
	for _, d := range dirs {
		nr, nc := d[0], d[1]
		fmt.Println(nr, nc)
	}
}
