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