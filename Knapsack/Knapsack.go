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
	var dp [100][10000]int
	var num [100]int
	var weight [100]int
	var value [100]int

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
