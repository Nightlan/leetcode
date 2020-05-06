package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mincostTickets(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}
	lastDay := days[len(days)-1]
	dp := make([]int, lastDay+1)
	for i, pos := 1, 0; i < len(dp); i++ {
		if i == days[pos] {
			lastDayPos := i - 1
			if lastDayPos < 0 {
				lastDayPos = 0
			}
			lastWeekPos := i - 7
			if lastWeekPos < 0 {
				lastWeekPos = 0
			}
			lastMonthPos := i - 30
			if lastMonthPos < 0 {
				lastMonthPos = 0
			}
			dp[i] = min(dp[lastDayPos]+costs[0], min(dp[lastWeekPos]+costs[1],
				dp[lastMonthPos]+costs[2]))
			pos++
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[lastDay]
}

func main() {

}
