package buy

import "fmt"

func Get(price []int) int {
	// 进行买入和卖出的最佳时机的选取，进行数据分析，遇到拐点进行继续分析
	max := 0
	i, j := 0, 1
	for j < len(price)-1 {
		if price[i] < price[j] && price[j] > price[j+1] {
			max += price[j] - price[i]
			fmt.Println("price[i],price[j],price[j+1],max ", i, j, j+1, price[i], price[j], price[j+1], max)
		} else if price[i] < price[j] && price[j] < price[j+1] && j == len(price)-2 {
			fmt.Println("price[i],price[j],price[j+1],max ", i, j, j+1, price[i], price[j], price[j+1], max)
			max += price[j+1] - price[i]
		}
		i = j
		j++
	}
	fmt.Println(i, j, max)
	return max
}

func Get2(prices []int) int {
	if len(prices) < 2 {
		return 0 // 如果价格数组长度小于2，无法进行买卖操作
	}

	maxProfit := 0
	buyIndex := 0
	sellIndex := 1

	for sellIndex < len(prices)-1 {
		// 如果当前价格低于下一个价格，并且下一个价格高于再下一个价格（拐点）
		if prices[buyIndex] < prices[sellIndex] && prices[sellIndex] > prices[sellIndex+1] {
			maxProfit += prices[sellIndex] - prices[buyIndex]
			fmt.Printf("Buy at %d, Sell at %d, Profit: %d, Total Profit: %d\n", buyIndex, sellIndex, prices[sellIndex]-prices[buyIndex], maxProfit)
		} else if prices[buyIndex] < prices[sellIndex] && prices[sellIndex] < prices[sellIndex+1] && sellIndex == len(prices)-2 {
			// 如果当前价格低于下一个价格，并且已经是倒数第二个元素
			maxProfit += prices[sellIndex+1] - prices[buyIndex]
			fmt.Printf("Buy at %d, Sell at %d, Profit: %d, Total Profit: %d\n", buyIndex, sellIndex+1, prices[sellIndex+1]-prices[buyIndex], maxProfit)
		}
		buyIndex = sellIndex
		sellIndex++
	}

	// 处理最后一个元素的情况
	if prices[buyIndex] < prices[sellIndex] {
		maxProfit += prices[sellIndex] - prices[buyIndex]
		fmt.Printf("Buy at %d, Sell at %d, Profit: %d, Total Profit: %d\n", buyIndex, sellIndex, prices[sellIndex]-prices[buyIndex], maxProfit)
	}

	fmt.Println("Final Max Profit:", maxProfit)
	return maxProfit
}
