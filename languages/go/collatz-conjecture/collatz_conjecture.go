package collatzconjecture

import (
	"errors"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be a positive integer")
	}

	if n == 1 {
		return 0, nil
	}

	var m int
	if n%2 == 0 {
		m = n / 2
	} else {
		m = 3*n + 1
	}

	steps, _ := CollatzConjecture(m)
	return steps + 1, nil
}
