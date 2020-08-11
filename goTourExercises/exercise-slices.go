package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	dp := make([][]uint8, dy)
	for i, _ := range dp {
		dp[i] = make([]uint8, dx)
	}

	for i, s := range dp {
		for j, _ := range s {
			dp[i][j] = uint8((i + j) / 2)
		}
	}

	return dp
}

func main1() {
	pic.Show(Pic)
}
