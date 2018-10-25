// Copyright 2018, Aaron Cannon under the BSD license.
// See the license file for details.

// package choose implements the N choose K formula, or the binomial coefficient
// formula. See https://en.wikipedia.org/wiki/Binomial_coefficient for more.
package choose

// Choose calculates n choose k. Overflows are not detected, and Choose panics
// if n >= k >= 0 is violated.
func Choose(n, k int64) int64 {
	if k > n {
		panic("Choose: k > n")
	}
	if k < 0 {
		panic("Choose: k < 0")
	}
	if n <= 1 || k == 0 || n == k {
		return 1
	}
	if newK := n - k; newK < k {
		k = newK
	}
	if k == 1 {
		return n
	}
	// Our return value, and this allows us to skip the first iteration.
	ret := int64(n - k + 1)
	for i, j := ret+1, int64(2); j <= k; i, j = i+1, j+1 {
		ret = ret * i / j
	}
	return ret
}
