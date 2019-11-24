// https://leetcode.com/problems/longest-substring-without-repeating-characters/
// lc: medium

/*

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
 "abcabcbb"
 */

function lengthOfLongestSubstring(str) {
  const longest = [0, 0];
  let charIdxMap = new Map();
  let startIdx = 0;

  for (let i = 0; i < str.length; i++) {
    const char = str[i];
    if (charIdxMap.has(char)) {
      startIdx = Math.max(startIdx, charIdxMap.get(char) + 1);
    }
    if (i + 1 - startIdx > longest[1] - longest[0]) {
      longest[0] = startIdx;
      longest[1] = i + 1;
    }
    charIdxMap.set(char, i);
  }
  return str.slice(longest[0], longest[1]);
}

