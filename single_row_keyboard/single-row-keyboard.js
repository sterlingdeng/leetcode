/*
https://leetcode.com/discuss/interview-question/356477

Input: keyboard = "abcdefghijklmnopqrstuvwxy", text = "cba"
Output: 4
Explanation:
Initially your finger is at index 0. First you have to type 'c'. The time taken to type 'c' will be abs(2 - 0) = 2 because character 'c' is at index 2.
The second character is 'b' and your finger is now at index 2. The time taken to type 'b' will be abs(1 - 2) = 1 because character 'b' is at index 1.
The third character is 'a' and your finger is now at index 1. The time taken to type 'a' will be abs(0 - 1) = 1 because character 'a' is at index 0.
The total time will therefore be 2 + 1 + 1 = 4.

*/
// strategy: create a map that maps the char -> index in the keyboard
// set the current idx = 0
// iterate through the text, calculate the distance from currentIdx to newIdx, update the sum, then update the idx
function singleRowKeyboard(keyboard, text) {
    var keyIdxMap = {};
    for (var i = 0; i < keyboard.length; i++) {
        keyIdxMap[keyboard[i]] = i;
    }
    var currIdx = 0;
    var distance = 0;
    for (var i = 0; i < text.length; i++) {
        distance += getDistance(currIdx, keyIdxMap[text[i]]);
        currIdx = keyIdxMap[text[i]];
    }
    return distance;
}
function getDistance(a, b) {
    return Math.abs(a - b);
}
console.log(singleRowKeyboard("abcdefghijklmnopqrstuvwxyz", "cba"));
