package climbing

func climbStairs(n int) int {
	a := 1
	b := 2
	if n <= 2 {
		return n
	}
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
