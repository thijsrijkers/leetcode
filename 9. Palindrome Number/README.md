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
function isPalindrome(x: number): boolean {
    const stringified = x.toString();
    const characters = stringified.split('');

    let result = "";

    for(let x = characters.length - 1; x >= 0; x--) {
        result += characters[x]
    }

    return result === stringified
};
```