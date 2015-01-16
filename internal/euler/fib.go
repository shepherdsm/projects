package euler

// GenNFib generates the first n fibonacci numbers starting at 1.
func GenNFib(n int) []int {
	return fibToN(n)
}

// Returns a fibonacci generator.
func makeFibGen() func() int {
	prev, start := 0, 1

	return func() int {
		prev, start = start, start+prev
		return start
	}
}

func fibToN(n int) []int {
	if n < 0 {
		return nil
	}

	fib := makeFibGen()
	results := make([]int, 0, n)

	for n > 0 {
		results = append(results, fib())
		n--
	}

	return results
}
