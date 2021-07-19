package main

//最大O(2n)
func lengthOfLongestSubstring1(s string) int {
	sArr := []byte(s)
	m := make(map[byte]int)
	res := 0
	j := 0
	for i, c := range sArr {
		m[c] ++
		for m[c] > 1 {
			m[sArr[j]]--
			j++
		}
		if res < (i - j + 1) {
			res = i - j + 1
		}
	}
	return res
}

//至多O(n)
func lengthOfLongestSubstring(s string) int {
	sArr := []byte(s)
	m := make(map[byte]int) //用来存(字符, 字符所在索引)
	res := 0
	j := 0
	for i, c := range sArr {
		if _, ok := m[c]; ok {
			if m[c] >= j { //一定要有等号
				j = m[c] + 1
			}
		}
		if res < (i - j + 1) {
			res = i - j + 1
		}
		m[c] = i
	}
	return res
}
