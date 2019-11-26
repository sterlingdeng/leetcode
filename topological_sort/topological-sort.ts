// write the algorithm to perform a topological sort

type adjList = {
  [key: string]: string[];
};

const adjList: adjList = {
  a: ["c"],
  b: ["c", "d"],
  c: ["e"],
  e: ["f"],
  d: ["f"],
  f: ["g"],
  g: []
};

function toplogicalSort(list: adjList): string[] {
  const toplogicallySorted: string[] = [];
  const visited = new Set<string>();
  const stack: string[] = [];

  function dfs(node: string) {
    if (visited.has(node)) {
      return;
    }
    visited.add(node);

    list[node].forEach(child => {
      dfs(child);
    });
    stack.push(node);
  }

  Object.keys(list).forEach(node => {
    dfs(node);
  });

  while (stack.length > 0) {
    toplogicallySorted.push(stack.pop());
  }

  return toplogicallySorted;
}

console.log(toplogicalSort(adjList));
