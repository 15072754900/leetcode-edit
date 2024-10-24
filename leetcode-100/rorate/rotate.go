package rorate

import "fmt"

// 旋转图像，leetcode48

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			matrix[i][j] = n*(n-j-1) + i + 1
		}
	}
	fmt.Println(matrix)
}

func rotate1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {

		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] = matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
