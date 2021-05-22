package main

// CountSubSequences counts how many subsequences of length n there are in a string of length m
func CountSubSequences(m, n int) int {
	// if there's as single letter in the sequence, there are m subsequences
	if n == 1 {
		return m
	}

	sum := 0
	// otherwise it's the sum of the count of subsequences for each starting position of the subsequence
	for c := n - 1; c < m; c++ {
		sum += CountSubSequences(c, n-1)
	}
	return sum
}

func fact(x int) int {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}

func Ank(n, k int) int {
	return fact(n) / fact(n-k)
}

func Cnk(n, k int) int {
	return Ank(n, k) / fact(k)
}
