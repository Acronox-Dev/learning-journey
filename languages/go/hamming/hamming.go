package hamming

import (
	"errors"
)

func Max(x, y int) int {
	if x <= y {
		return y
	}
	return x
}

func Distance(a, b string) (int, error) {
	var count int
	if len(a) != len(b) {
		return 0, errors.New("Sequences of different lengths")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}

	return count, nil
}
