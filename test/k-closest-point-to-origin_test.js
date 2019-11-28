const assert = require("assert");
const {
  kClosest
} = require("../k_closest_point_to_origin/k-closest-point-to-origin");

describe("k closest to origin", () => {
  const main = [
    [1, 1],
    [2, 2],
    [3, 3],
    [4, 4],
    [5, 5],
    [6, 6],
    [7, 7],
    [8, 8],
    [9, 9],
    [10, 10]
  ];
  const testCases = [
    {
      name: "lc1",
      in: {
        arr: [
          [1, 3],
          [-2, 2]
        ],
        k: 1
      },
      expect: [[-2, 2]]
    },
    {
      name: "lc2",
      in: {
        arr: [
          [3, 3],
          [5, -1],
          [-2, 4]
        ],
        k: 2
      },
      expect: [
        [3, 3],
        [-2, 4]
      ]
    },
    {
      name: "empty array",
      in: {
        arr: [],
        k: 1
      },
      expect: []
    },
    {
      name: "test k = 1",
      in: {
        arr: main,
        k: 1
      },
      expect: [[1, 1]]
    },
    {
      name: "test k = 2",
      in: {
        arr: main,
        k: 2
      },
      expect: [
        [1, 1],
        [2, 2]
      ]
    },
    {
      name: "test k = 3",
      in: {
        arr: main,
        k: 3
      },
      expect: [
        [1, 1],
        [2, 2],
        [3, 3]
      ]
    },
    {
      name: "test k = 4",
      in: {
        arr: main,
        k: 4
      },
      expect: [
        [1, 1],
        [2, 2],
        [3, 3],
        [4, 4]
      ]
    },
    {
      name: "test k = 5",
      in: {
        arr: main,
        k: 5
      },
      expect: [
        [1, 1],
        [2, 2],
        [3, 3],
        [4, 4],
        [5, 5]
      ]
    },
    {
      name: "test k = 6",
      in: {
        arr: main,
        k: 6
      },
      expect: [
        [1, 1],
        [2, 2],
        [3, 3],
        [4, 4],
        [5, 5],
        [6, 6]
      ]
    }
  ];
  testCases.forEach(testCase => {
    it(testCase.name, () => {
      const got = kClosest(testCase.in.arr, testCase.in.k);
      assert.deepEqual(
        got,
        testCase.expect,
        `did not sort, got ${got}, but expected ${testCase.expect}`
      );
    });
  });
});
