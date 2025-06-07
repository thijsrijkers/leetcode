func missingNumber(nums []int) int {
    length := len(nums)

    total := length * (length + 1) / 2
    actual := 0

    for _, num := range nums {
        actual += num
    }

    return total - actual
}