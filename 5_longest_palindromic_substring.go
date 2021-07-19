package main

//方法1：动态规划，时间O(n*n), 时间O(n*n)
func longestPalindrome1(s string) string {
	if len(s) < 2 {
		return s
	}
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	res := string(s[0])
	maxLen := 1
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] && (j == i || j == i+1 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
			if dp[i][j] && j-i+1 > maxLen {
				res = s[i : j+1]
				maxLen = j - i + 1
			}
		}
	}
	return res
}

//方法2：从中间往外散发，时间O(n*n)，空间O(1)
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	left := 0
	right := 0
	maxLen := 1
	for i := 0; i < len(s); i++ {
		l, r := getPalindromeEdge(s, i, i)
		if r-l+1 > maxLen {
			left = l
			right = r
			maxLen = r - l + 1
		}
		l, r = getPalindromeEdge(s, i, i+1)
		if r-l+1 > maxLen {
			left = l
			right = r
			maxLen = r - l + 1
		}
	}
	return s[left : right+1]
}

func getPalindromeEdge(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return i + 1, j - 1
}
