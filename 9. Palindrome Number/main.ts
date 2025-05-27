function isPalindrome(x: number): boolean {
    const stringified = x.toString();
    const characters = stringified.split('');

    let result = "";

    for(let x = characters.length - 1; x >= 0; x--) {
        result += characters[x]
    }

    return result === stringified
};