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