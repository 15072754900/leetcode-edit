package lemonadeChange

import "fmt"

// leetcode 860 柠檬水找零

func lemonadeChange(bills []int) bool {
	if len(bills) >= 1 && bills[0] != 5 {
		return false
	}
	hash := make(map[int]int)
	hash[5] = 1
	// 使用面额比当前数字小一点的进行处理
	defer fmt.Println(hash)
	for i := 1; i < len(bills); i++ {
		switch bills[i] {
		case 5:
			hash[5] += 1
		case 10:
			if hash[5] > 0 {
				hash[5] -= 1
				hash[10] += 1
			} else {
				return false
			}
		case 20:
			if hash[10] > 0 && hash[5] > 0 {
				hash[10] -= 1
				hash[5] -= 1
				hash[20] += 1
			} else if hash[10] == 0 && hash[5] > 2 {
				hash[5] -= 3
				hash[20] += 1
			} else {
				return false
			}
		}
	}
	//fmt.Println(hash, true)
	return true
}
