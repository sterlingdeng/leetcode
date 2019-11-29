const assert = require("assert");
const {
  evalRPN
} = require("../evaluate_reverse_polish_notations/evaluate-reverse-polish-notations");

describe("evaluate reverse polish notation", () => {
  const testCases = [
    {
      name: "leetcode test case 1",
      in: ["2", "1", "+", "3", "*"],
      expect: 9
    },
    {
      name: "leet code test case 2",
      in: ["4", "13", "5", "/", "+"],
      expect: 6
    },
    {
      name: "leet code test case 3",
      in: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"],
      expect: 22
    }
  ];
  testCases.forEach(testCase => {
    it(testCase.name, () => {
      const got = evalRPN(testCase.in);
      assert.equal(
        got,
        testCase.expect,
        `evalRPN() = ${got}, but expected ${testCase.expect}`
      );
    });
  });
});
