package differenceofsquares

func SquareOfSum(n int) int {
	return (n * n * (n + 1) * (n + 1)) / 4
}

func SumOfSquares(n int) int {
	if n == 0 {
		return 0
	}
	return n*n + SumOfSquares(n-1)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
