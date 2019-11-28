// http://leetcode.liangjiateng.cn/leetcode/shortest-distance-from-all-buildings/description
// LC hard

const OBSTACLE = 2;
const BUILDING = 1;

function shortest(arr) {
  let shortestPath = -1;
  const visited = Array(arr.length)
    .fill(null)
    .map(() => Array(arr[0].length).fill(false));

  for (let r = 0; r < arr.length; r++) {
    for (let c = 0; c < arr[0].length; c++) {
      const houseVisited = new Map();
      if (arr[r][c] !== BUILDING && arr[r][c] !== OBSTACLE) {
        helper(r, c, arr, houseVisited, visited);
      }
      let distance = 0;
      houseVisited.forEach(v => {
        distance += v;
      });
      if (distance > 0 && shortestPath !== -1) {
        shortestPath = Math.min(shortestPath, distance);
      } else if (distance > 0) {
        shortestPath = distance;
      }
    }
  }
  return shortestPath;
}

function helper(r, c, arr, houseVisited, visited, distance = 0, sum = 0) {
  if (isOutside(r, c, arr)) return sum;
  if (arr[r][c] === OBSTACLE) return sum;
  if (visited[r][c]) return sum;
  if (arr[r][c] === BUILDING) {
    // serialize rc coordinate
    const serialized = JSON.stringify([r, c]);
    if (houseVisited.has(serialized)) {
      const currDistance = houseVisited.get(serialized);
      houseVisited.set(serialized, Math.min(currDistance, distance));
    } else {
      houseVisited.set(serialized, distance);
    }
  }
  visited[r][c] = true;

  sum = helper(r + 1, c, arr, houseVisited, visited, distance + 1, sum);
  sum = helper(r - 1, c, arr, houseVisited, visited, distance + 1, sum);
  sum = helper(r, c + 1, arr, houseVisited, visited, distance + 1, sum);
  sum = helper(r, c - 1, arr, houseVisited, visited, distance + 1, sum);

  visited[r][c] = false;
  return sum;
}

function isOutside(r, c, arr) {
  return r < 0 || c < 0 || r > arr.length - 1 || c > arr[0].length - 1;
}

const input = [
  [1, 0, 2, 0, 1],
  [0, 0, 0, 0, 0],
  [0, 0, 1, 0, 0]
];

console.log(shortest(input));

const input2 = [
  [1, 1],
  [1, 1]
];

console.log(shortest(input2));
