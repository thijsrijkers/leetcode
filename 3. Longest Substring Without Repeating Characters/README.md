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
function lengthOfLongestSubstring(s: string): number {
    let longestSubstring: string[] = [];
    let substring: string[] = [];

    for (let char of s) {
        if (substring.includes(char)) {
            substring = substring.slice(substring.indexOf(char) + 1);
        }
        substring.push(char);

        if (substring.length > longestSubstring.length) {
            longestSubstring = [...substring];
        }
    }

    return longestSubstring.length;
}


```