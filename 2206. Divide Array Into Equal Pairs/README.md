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
import "fmt"

func divideArray(nums []int) bool {
    startOfLoop := 0;

    for startOfLoop < len(nums) {
        secondLoop := startOfLoop + 1;
        found := false;

        for secondLoop < len(nums) {
            if nums[startOfLoop] == nums[secondLoop] {
                nums = append(nums[:secondLoop], nums[secondLoop+1:]...)
                nums = append(nums[:startOfLoop], nums[startOfLoop+1:]...)
                found = true;
                break;
            }

            secondLoop++;
        }

        if !found {
            return false;
        }
    }

    return true;
}
```