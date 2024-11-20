package canCompleteCircuit

import "fmt"

// leetcode 134 加油站

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		cur := 0
		for j := i; j < len(cost)+i; j++ {
			if j >= len(cost) {
				cur += gas[j-len(cost)] - cost[j-len(cost)]
				fmt.Println(cur)
				if cur >= 0 && j == len(cost)+i-1 {
					fmt.Println("res", j-len(cost)+1)
					return j
				} else if cur < 0 {
					break
				}
			} else {
				cur += gas[j] - cost[j]
				fmt.Println(cur)
				if cur >= 0 && j == len(cost)+i-1 {
					fmt.Println("res", j-len(cost)+1)
					return j
				} else if cur < 0 {
					break
				}
			}
		}
	}
	//fmt.Println(-1)
	return -1
}

func notToMuch(gas, cost []int) int {
	// 及其简化的方法，不会发生超时问题
	// 总结上述的超时原因：1.计算数据过大，直接相加会发生超出变量限制的大小。2.计算时间复杂度过大，跑完两重for循环发生的超时。

	// 这次不进行到达和容量满足的同时判断，首先判断能否启动，然后判断能否到达，能启动指的是当前节点能启动，且后面的都能启动
	tank := 0
	have := 0
	lost := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		have += gas[i]
		lost += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if have < lost {
		return -1
	}
	return start
}
