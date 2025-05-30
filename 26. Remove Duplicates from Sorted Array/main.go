func removeDuplicates(nums []int) int {
    seenItems := map[int]bool{}
    i := 0

    for i < len(nums) {
		if !seenItems[nums[i]] {
            seenItems[nums[i]] = true
            i++
        } else {
            nums = append(nums[:i], nums[i+1:]...)
        }
    }
    return len(nums)
}
