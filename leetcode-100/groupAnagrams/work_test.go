package groupAnagrams

import (
	"testing"
)

func TestWork(t *testing.T) {
	//s1 := []byte("hufeng")
	//fmt.Println(s1)
	//n1 := len(s1)
	//quickSort(s1, 0, n1-1)
	//fmt.Println(s1, string(s1))
	//
	//s2 := []rune("hufeng")
	//fmt.Println(s2)
	//n2 := len(s2)
	//quickSortRune(s2, 0, n2-1)
	//fmt.Println(s2, string(s2))

	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

}
