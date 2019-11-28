function groupAnagrams(words) {
  // Write your code here.
  const map = new Map();
  words.forEach(word => {
    const sorted = Array.from(word)
      .sort()
      .join("");
    if (!map.has(sorted)) {
      map.set(sorted, [word]);
    } else {
      map.get(sorted).push(word);
    }
  });

  const output = [];
  map.forEach((v, k) => {
    output.push(v);
  });
  return output;
}

console.log(groupAnagrams(["yo", "act", "flop", "tac", "cat", "oy", "olfp"]));
