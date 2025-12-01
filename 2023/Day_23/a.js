const dirs = {
  "^": [{ x: 0, y: -1 }],
  ">": [{ x: 1, y: 0 }],
  v: [{ x: 0, y: 1 }],
  "<": [{ x: -1, y: 0 }],
  ".": [
    { x: 0, y: -1 },
    { x: 1, y: 0 },
    { x: 0, y: 1 },
    { x: -1, y: 0 },
  ],
};

const path = ".";
const forest = "#";

/** PARSE INPUT */
const input = Bun.file("Day_23/input.txt");
const text = await input.text();

/** @type {string[][]} */
const map = text.split("\n").map((col) => col.split(""));

const startRow = map[0].findIndex((x) => x === path);
const start = { x: startRow, y: 0 };

const endRow = map[map.length - 1].findIndex((x) => x === path);
const end = { x: endRow, y: map.length - 1 };
/** END PARSE INPUT */

/**
 * @function
 * @template T
 * @param {T[][]} map
 * @param {T} path
 * @param {T} forest
 * @returns {{[index: string]: []}}
 */
function getNodes(map, path, forest) {
  const result = {
    [pointToIndex(map, start)]: [],
    [pointToIndex(map, end)]: [],
  };

  for (let r = 0; r < map.length; ++r) {
    for (let c = 0; c < map[0].length; ++c) {
      if (map[r][c] === forest) {
        continue;
      }

      let nextDirs = 0;

      for (let d = 0; d < dirs[map[r][c]].length; ++d) {
        const dir = dirs[map[r][c]][d];
        const next = { x: c + dir.x, y: r + dir.y };

        const { x, y } = next;

        if (x < 0 || y < 0 || x >= map[0].length || y >= map.length) {
          continue;
        } else if (map[y][x] === forest) {
          continue;
        } else {
          nextDirs += 1;
        }
      }

      // prev, next, ...others
      if (nextDirs > 2) {
        result[pointToIndex(map, { x: c, y: r })] = [];
      }
    }
  }

  return result;
}

/**
 * @function
 * @template T
 * @param {T[][]} map
 * @param {number} node
 * @param {number} currWeight
 * @param {{[index: string]: []}} nodes
 * @returns {{[index: string]: {index: number; weight: number;}[];}}
 */
function getWeights(map, nodes) {
  const result = nodes;

  for (const node in nodes) {
    const stack = [{ index: node, weight: 0 }];
    const seen = { [node]: true };

    while (stack.length > 0) {
      const curr = stack.pop();
      const currPoint = indexToPoint(map, curr.index);

      if (nodes[curr.index] && curr.weight > 0) {
        result[node].push(curr);
        continue;
      }

      for (let i = 0; i < dirs[map[currPoint.y][currPoint.x]].length; ++i) {
        const dir = dirs[map[currPoint.y][currPoint.x]][i];

        const nextPoint = { x: currPoint.x + dir.x, y: currPoint.y + dir.y };
        const nextIndex = pointToIndex(map, nextPoint);

        const { x, y } = nextPoint;

        if (x < 0 || y < 0 || x >= map[0].length || y >= map.length) {
          continue;
        } else if (map[y][x] === forest) {
          continue;
        } else if (seen[nextIndex]) {
          continue;
        } else {
          seen[nextIndex] = true;
          stack.push({ index: nextIndex, weight: curr.weight + 1 });
        }
      }
    }
  }

  return result;
}

/**
 * @function
 * @template T
 * @param {T[][]} map
 * @param {{[index: string]: {index: number; weight: number}[]}} weightedNodes
 * @returns {number}}
 */
function getLongestHike(map, weightedNodes) {
  let result = 0;

  const startIndex = pointToIndex(map, start);
  const stack = [{ index: startIndex, accWeight: 0 }];

  while (stack.length) {
    const curr = stack.pop();

    if (curr.accWeight > result) {
      result = curr.accWeight;
    }

    for (let i = 0; i < weightedNodes[curr.index].length; ++i) {
      const next = weightedNodes[curr.index][i];

      stack.push({
        index: next.index,
        accWeight: curr.accWeight + next.weight,
      });
    }
  }

  return result;
}

const nodes = getNodes(map, path, forest);
const weightedNodes = getWeights(map, nodes);
const length = getLongestHike(map, weightedNodes);

console.log(length);

/** @type {(map: string, point: {x: number, y: number}) => number}  */
function pointToIndex(map, { x, y }) {
  return y * map[0].length + x;
}

/**
 * @function
 * @template T
 * @param {T[][]} map
 * @param {number} index
 * @returns {{x: number, y: number}}
 */
function indexToPoint(map, index) {
  return { x: index % map[0].length, y: Math.floor(index / map[0].length) };
}
