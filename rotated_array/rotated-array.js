// somebody rotated a sorted array. find the smallest element
// example: [6, 7, 9, 15, 19, 2, 3]

function smallestElement(arr) {
  let smallest = Infinity;
  let left = 0;
  let right = arr.length - 1;
  while (left <= right) {
    const mid = left + Math.floor((right - left) / 2);
    if (arr[mid] > arr[right]) {
      left = mid + 1;
    } else {
      smallest = Math.min(smallest, arr[mid]);
      right = mid - 1;
    }
  }
  return smallest;
}

const input = [6, 7, 9, 15, 19, 3, 4, 5];
console.log(smallestElement(input));
