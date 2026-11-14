func hasDuplicate(nums []int) bool {
    st := make(map[int]bool)
	for _, num := range nums {
		if st[num] {
			return true
		}
		st[num] = true
	}
	return false
}