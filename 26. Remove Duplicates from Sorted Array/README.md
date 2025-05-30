# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->

# Approach
<!-- Describe your approach to solving the problem. -->

# Complexity
- Time complexity:
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity:
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```golang []
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

```