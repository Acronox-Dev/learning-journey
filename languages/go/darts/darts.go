package darts

func Score(x, y float64) int {
	if x*x+y*y <= 1 {
		return 10
	}
	if x*x+y*y <= 5 {
		return 5
	}
	if x*x+y*y <= 10 {
		return 1
	}
	return 0
}
