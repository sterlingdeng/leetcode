/*
https://www.geeksforgeeks.org/longest-subarray-sum-elements-atmost-k/

Given an array of integers, our goal is to find the length of largest subarray having sum of its elements atmost ‘k’ where k>0.

Examples:

Input : arr[] = {1, 2, 1, 0, 1, 1, 0}, 
					 k = 4
Output : 5
Explanation:
 {1, 2, 1} => sum = 4, length = 3
 {1, 2, 1, 0}, {2, 1, 0, 1} => sum = 4, length = 4
 {1, 0, 1, 1, 0} =>5 sum = 3, length = 5


*/
// caterpillar
function longest(arr, k) {
  let longest = 0;
  let maxSum = 0;
  let front = 0;
  let back = 0;
  let currSum = 0;
  while (front < arr.length) {
    if (currSum + arr[front] <= k) {
      currSum += arr[front];
      maxSum = Math.max(maxSum, currSum);
      front++;
    } else {
      currSum -= arr[back];
      back++;
    }
    longest = Math.max(longest, front - back);
  }
  console.log("maxSum:", maxSum);
  return longest;
}

console.log(longest([1, 2, 1, 0, 1, 1, 0], 4));

console.log(longest([6, 8, 9], 20));
