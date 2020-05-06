package two_sum

func TwoSum(nums []int, target int) []int {
	imap := make(map[int]int)
	for i, v := range nums {
		need := target - v
		if need_i, ok := imap[need]; ok {
			return []int{need_i, i}
		}
		imap[v] = i
	}

	return nil
}
