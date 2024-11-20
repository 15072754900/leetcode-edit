package findContentChildren

import (
	"fmt"
	"sort"
)

func finContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := len(s)-1, len(g)-1
	total := 0
	for i >= 0 && j >= 0 {
		if s[i] >= g[j] {
			total++
			i--
			j--
		} else {
			fmt.Println(i, j)
			j--
		}
	}
	fmt.Println(total)
	return total
}
