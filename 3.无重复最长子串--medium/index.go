package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	len := len(s)
	if len <= 1 {
		return len
	}
	maxLength := 0
	start, end := 0, 0
	maps := make(map[byte]int)
	for end < len && start < len {
		if maps[s[end]] == 0 {
			maps[s[end]] = 1
			end = end + 1
			if end-start > maxLength {
				maxLength = end - start
			}
		} else {
			maps = make(map[byte]int)
			start = start + 1
			for start < len-1 && s[start] == s[start+1] {
				start = start + 1
			}
			end = start
		}
	}
	return maxLength
}

func lengthOfLongestSubstring2(s string) int {
	len := len(s)
	if len <= 1 {
		return len
	}
	maxLength := 1
	left, right, maps := 0, 0, make(map[byte]int)
	for right < len {
		rightChar := s[right]
		rightCharIdx := 0

		if _, ok := maps[rightChar]; ok {
			rightCharIdx = maps[rightChar]
		}

		if rightCharIdx > left {
			left = rightCharIdx
		}

		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}

		maps[rightChar] = right + 1
		right++

	}

	return maxLength
}

func main() {
	s := "abcabcbh"
	fmt.Println(lengthOfLongestSubstring(s))
}
