package two_sum

import "fmt"

func TwoSum(nums []int, target int) []int {
	imap := make(map[int]int)
	for i, v := range nums {
		need := target - v
		if need_i, ok := imap[i]; ok {
			return []int{i, need_i}
		}
		imap[v] = i
	}

	return nil
}
