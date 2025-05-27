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
function fizzBuzz(n: number): string[] {
    const fizzBuzz = [];
    let i = 1;
    
    while(i <= n) {
        let word = '';

        if(i % 3 === 0 ) word += 'Fizz';

        if(i % 5 === 0 ) word += 'Buzz';

        if(word === '') word = `${i}`;

        fizzBuzz.push(word);
        i++;
    }

    return fizzBuzz;
};
```