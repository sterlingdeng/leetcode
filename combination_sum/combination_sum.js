/*
https://leetcode.com/problems/combination-sum/
lc medium

Given a set of candidate numbers (candidates) (without duplicates) and a target number (target), find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.
Example 1:

Input: candidates = [2,3,6,7], target = 7,
A solution set is:
[
  [7],
  [2,2,3]
]
Example 2:

Input: candidates = [2,3,5], target = 8,
A solution set is:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]

*/

function combinationSum(candidates, target) {
  return helper(candidates, target, 0);
}

function helper(candidates, currVal, idx, currPerm = [], output = []) {
  if (currVal === 0) {
    output.push([...currPerm]);
    return output;
  }
  if (currVal < 0) {
    return output;
  }
  for (let i = idx; i < candidates.length; i++) {
    const candidate = candidates[i];
    currPerm.push(candidate);
    helper(candidates, currVal - candidate, i, currPerm, output);
    currPerm.pop();
  }
  return output;
}

const candidates = [2, 3, 5];
const target = 8;
console.log(combinationSum(candidates, target));
