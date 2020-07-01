package leetcode

/*
You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins.

Given n, find the total number of full staircase rows that can be formed.

n is a non-negative integer and fits within the range of a 32-bit signed integer.

Example 1:

n = 5

The coins can form the following rows:
¤
¤ ¤
¤ ¤

Because the 3rd row is incomplete, we return 2.
Example 2:

n = 8

The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤

Because the 4th row is incomplete, we return 3.
*/
func arrangeCoins(n int) int {
	var low, high int = 1, n

	for low < high {
		mid := low + (high-low+1)/2
		if (mid+1)*mid/2.0 <= n {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return high
}
