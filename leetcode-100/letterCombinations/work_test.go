package letterCombinations

import (
	"fmt"
	"testing"
)

func TestWork(t *testing.T) {
	letterCombinations("23")
}

func stringToInts(s string) {
	fmt.Println([]byte(s))
	//nums := []byte(s)
	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(nums[i] - []byte("0")[0])
	//}
	fmt.Println([]rune(s))
	runes := []rune(s)
	fmt.Println(string(runes))
	fmt.Println("1", string(s[0]), string(s[1]))
}
