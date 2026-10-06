package primefactors

func Factors(n int64) []int64 {
	var primes_factors []int64
	var p int64
	p = 2

	for n > 1 {
		if n%p == 0 {
			n = n / p
			primes_factors = append(primes_factors, p)
		} else {
			p += 1
		}
	}

	return primes_factors
}
