/*

We have a list of points on the plane.  Find the K closest points to the origin (0, 0).

(Here, the distance between two points on a plane is the Euclidean distance.)

You may return the answer in any order.  The answer is guaranteed to be unique (except for the order that it is in.)



Example 1:

Input: points = [[1,3],[-2,2]], K = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], K = 2
Output: [[3,3],[-2,4]]
(The answer [[-2,4],[3,3]] would also be accepted.)


Note:

1 <= K <= points.length <= 10000
-10000 < points[i][0] < 10000
-10000 < points[i][1] < 10000
 */

var kClosest = function(arr, k) {
  if (!arr.length) {
    return arr;
  }
  const kIdx = k - 1;
  kClosestHelper(arr, 0, arr.length - 1, kIdx);
  return arr.slice(0, k);
};

function kClosestHelper(arr, low, high, kIdx) {
  if (low < high) {
    const pIdx = partition(arr, low, high);
    if (pIdx === kIdx) {
      return;
    }
    if (pIdx < kIdx) {
      return kClosestHelper(arr, pIdx + 1, high, kIdx);
    } else {
      return kClosestHelper(arr, low, pIdx - 1, kIdx);
    }
  }
}

function partition(arr, low, high) {
  const partitionValue = getDistance(arr[high]);
  let partitionIdx = low;
  for (let i = low; i < high; i++) {
    if (getDistance(arr[i]) < partitionValue) {
      swap(arr, partitionIdx, i);
      partitionIdx++;
    }
  }
  swap(arr, partitionIdx, high);
  return partitionIdx;
}

function swap(arr, a, b) {
  [arr[a], arr[b]] = [arr[b], arr[a]];
}

function getDistance(points) {
  return Math.sqrt(points[0] ** 2 + points[1] ** 2);
}

module.exports = { kClosest };
