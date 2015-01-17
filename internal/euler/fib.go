package euler

// GenNFib generates the first n fibonacci numbers starting at 1.
func GenNFib(n int) []int {
	return fibToN(n)
}

func GenFibUnder(n int) []int {
	return fibUnderN(n)
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

func fibUnderN(n int) []int {
	if n < 0 {
		return nil
	}

	fib := makeFibGen()
	results := make([]int, 0, 10)

	num := fib()
	for num <= n {
		results = append(results, num)
		num = fib()
	}

	return results
}

// Returns a fibonacci generator.
func makeFibGen() func() int {
	prev, start := 0, 1

	return func() int {
		prev, start = start, start+prev
		return start
	}
}
