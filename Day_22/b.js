const dirs = [
  { x: 0, y: -1 },
  { x: 1, y: 0 },
  { x: 0, y: 1 },
  { x: -1, y: 0 },
];

const path = ".";
const forest = "#";

/** PARSE INPUT */
const input = Bun.file("Day_22/input.txt");
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
function getNodes(map) {
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

      for (let d = 0; d < dirs.length; ++d) {
        const dir = dirs[d];
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

      for (let i = 0; i < dirs.length; ++i) {
        const dir = dirs[i];

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
  const startIndex = pointToIndex(map, start);
  const endIndex = pointToIndex(map, end);
  const seen = {};

  function walk(index) {
    if (index === endIndex) {
      return 0;
    }

    let result = 0;

    seen[index] = true;
    for (let i = 0; i < weightedNodes[index].length; ++i) {
      const next = weightedNodes[index][i];
      if (seen[next.index]) {
        continue;
      }
      result = Math.max(result, walk(next.index) + next.weight);
    }
    seen[index] = false;

    return result;
  }

  return walk(startIndex);
}

const nodes = getNodes(map);
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
