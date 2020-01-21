package main

import (
	"fmt"
	"math/rand"
	"time"
)

func max(before int, after int) int {
	if before > after {
		return before
	} else {
		return after
	}
}

func main() {
	var dp []int
	var num []int

	n := 2
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		number := 10 - rand.Intn(20)
		num = append(num, number)
	}

	dp = append(dp, 0)
	for i := 0; i < n; i++ {
		dp = append(dp, max(dp[i], dp[i]+num[i]))
	}

	fmt.Println(num, dp[n])
}
