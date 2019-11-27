/*
Google 

https://leetcode.com/discuss/interview-question/124724/

Given a sorted array of n elements, possibly with duplicates, find the number of occurrences of the target element.

Example 1:

Input: arr = [4, 4, 8, 8, 8, 15, 16, 23, 23, 42], target = 8
Output: 3
Example 2:

Input: arr = [3, 5, 5, 5, 5, 7, 8, 8], target = 6
Output: 0
Example 3:

Input: arr = [3, 5, 5, 5, 5, 7, 8, 8], target = 5
Output: 4
Expected O(logn) time solution.

*/

function getCount(arr, k) {
  let left = 0;
  let right = arr.length - 1;
  let midIdx;
  let found = false;
  while (left <= right) {
    const mid = Math.floor(left + (right - left) / 2);
    if (arr[mid] === k) {
      midIdx = mid;
      found = true;
      break;
    }
    if (arr[mid] < k) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }
  if (!found) {
    return 0;
  }

  // search for the first occurence of k on the left
  left = 0;
  right = midIdx;
  let firstIdx = Infinity;

  while (left <= right) {
    const mid = Math.floor(left + (right - left) / 2);
    if (arr[mid] === k) {
      firstIdx = Math.min(firstIdx, mid);
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }

  left = midIdx;
  right = arr.length - 1;
  let lastIdx = -Infinity;

  while (left <= right) {
    const mid = Math.floor(left + (right - left) / 2);
    if (arr[mid] === k) {
      lastIdx = Math.max(lastIdx, mid);
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }
  return lastIdx - firstIdx + 1;
}

function binarySearch(array, k, leftIdx) {
  let left = 0;
  let right = array.length - 1;
  let idx = -1;
  while (left <= right) {
    const mid = Math.floor(left + (right - left) / 2);
    if (array[mid] < k) {
      left = mid + 1;
    } else if (array[mid] > k) {
      right = mid - 1;
    } else {
      idx = mid;
      if (leftIdx) {
        right = mid - 1;
      } else {
        left = mid + 1;
      }
    }
  }
  return idx;
}

function driver(array, k) {
  const left = binarySearch(array, k, true);
  if (left < 0) {
    return 0;
  }
  const right = binarySearch(array, k, false);
  return right - left + 1;
}

console.log(driver([4, 4, 8, 8, 8, 15, 16, 23, 23, 42], 8));
console.log(driver([3, 5, 5, 5, 5, 7, 8, 8], 6));
console.log(driver([3, 5, 5, 5, 5, 7, 8, 8], 5));
