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
```typescript []
function twoSum(nums: number[], target: number): number[] {
    for(let i = 0; i < nums.length; i++){
        for(let x = 0; x < nums.length; x++){
            if(i === x) continue;

            if((nums[i] + nums[x]) === target) return [i, x];
        }
    }
};
```