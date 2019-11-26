// https://www.geeksforgeeks.org/find-the-longest-substring-with-k-unique-characters-in-a-given-string/

/*

Given a string you need to print longest possible substring that has exactly M unique characters.
If there are more than one substring of longest possible length, then print any one of them.
Examples:

"aabbcc", k = 1
Max substring can be any one from {"aa" , "bb" , "cc"}.

"aabbcc", k = 2
Max substring can be any one from {"aabb" , "bbcc"}.

"aabbcc", k = 3
There are substrings with exactly 3 unique characters
{"aabbcc" , "abbcc" , "aabbc" , "abbc" }
Max is "aabbcc" with length 6.

"aaabbb", k = 3
There are only two unique characters, thus show error message. 
Source: Google Interview Question.

*/

function solution(str: string, k: number): string {
  const charSet = new Map<string, number>();

  let leftPtr = 0;
  let rightPtr = 0;
  let longestSubstring = "";

  while (rightPtr < str.length) {
    const currChar = str[rightPtr];
    if (charSet.size <= k) {
      if (charSet.has(currChar)) {
        charSet.set(currChar, charSet.get(currChar) + 1);
      } else {
        charSet.set(currChar, 1);
      }
      rightPtr++;
      if (rightPtr - leftPtr > longestSubstring.length && charSet.size === k) {
        longestSubstring = str.slice(leftPtr, rightPtr);
      }
    } else {
      while (charSet.size > k) {
        const leftPtrChar = str[leftPtr];
        const count = charSet.get(leftPtrChar);
        if (count - 1 === 0) {
          charSet.delete(leftPtrChar);
        } else {
          charSet.set(leftPtrChar, count - 1);
        }
        leftPtr++;
      }
    }
  }
  return longestSubstring;
}

console.log(solution("aabbcc", 1));
console.log(solution("aabbcc", 2));
console.log(solution("aabbcc", 3));
console.log(solution("aaabbb", 3));
console.log(solution("aabacbebebe", 3));
