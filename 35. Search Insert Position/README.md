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
func searchInsert(nums []int, target int) int {
    if len(nums) == 1 {
        if nums[0] >= target {
            return 0
        }
        return 1
    }

    mid := len(nums) / 2

    if nums[mid] == target {
        return mid
    }

    if nums[mid] > target {
        i := 0
        for i <= mid {
            if target == nums[i] {
                return i
            }

            if nums[i] > target {
                return i;
            }
            
            i++
        }
    } else {
        i := len(nums) - 1         
        for i >= mid {
            if target == nums[i] {
                return i
            }
   
            if nums[i] < target {
                return i + 1;
            }

            i--
        }
    }

    return 0
}
```