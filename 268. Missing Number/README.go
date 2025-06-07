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
func missingNumber(nums []int) int {
    length := len(nums)

    total := length * (length + 1) / 2
    actual := 0

    for _, num := range nums {
        actual += num
    }

    return total - actual
}
```