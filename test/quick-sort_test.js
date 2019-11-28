const assert = require("assert");
const { quickSort } = require("../quick_sort/quick-sort");

describe("quicksort", () => {
  const testCases = [
    {
      name: "easy",
      in: [5, 4, 3, 2, 1],
      expect: [1, 2, 3, 4, 5]
    },
    {
      name: "contains negatives",
      in: [-8, 13, 2, 4, -9, -1],
      expect: [-9, -8, -1, 2, 4, 13]
    },
    {
      name: "empty array",
      in: [],
      expect: []
    }
  ];
  testCases.forEach(testCase => {
    it(testCase.name, () => {
      const got = quickSort(testCase.in);
      assert.deepEqual(got, testCase.expect, "did not sort");
    });
  });
});
