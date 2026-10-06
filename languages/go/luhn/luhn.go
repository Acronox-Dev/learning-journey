package luhn

import "strings"

func isDigit(x byte) bool {
	return '0' <= x && x <= '9'
}

func Valid(id string) bool {
	// Formatting input
	id = strings.ReplaceAll(id, " ", "")

	// Validating input
	if len(id) <= 1 {
		return false
	}

	// Doubling the digits and summing digits
	sum := 0
	calculate := false

	for i := len(id) - 1; i >= 0; i-- {
		// Check if id[i] is a digit
		if !isDigit(id[i]) {
			return false
		}
		digit := int(id[i] - '0')

		if calculate {
			// Calculate the new digit
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		calculate = !calculate
		sum += digit
	}

	return sum%10 == 0

}
