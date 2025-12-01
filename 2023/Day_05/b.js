const file = Bun.file("Day_05/input.txt");
const input = await file.text();

const dupa = input.split("\n\n");
const seeds = dupa[0].split(": ")[1].split(" ").map(Number);
const maps = dupa.slice(1).map((x) =>
  x
    .split("\n")
    .slice(1)
    .map((x) => x.split(" ").map(Number))
);

let rangeSeeds = [];
for (let i = 0; i < seeds.length; i += 2) {
  rangeSeeds.push([seeds[i], seeds[i] + seeds[i + 1]]);
}

maps.forEach((map) => {
  rangeSeeds = mappingFn(rangeSeeds, map);
});

const s = rangeSeeds.map(([a]) => a);
console.log(Math.min.apply(null, s));

// [55, 68), [79, 93)

// [57, 70), [81, 95)

// soil-to-fertilizer map:
// [57, 70), [81, 95)

// fertilizer-to-water map:
// [53, 57), [61, 70), [81, 95)

// water-to-light map:
// [46, 50), [54, 63), [74, 88)

// light-to-temperature map:
// [45, 56), [78, 81), [82, 86), [90, 99)

// temperature-to-humidity map:
// [46, 57), [78, 81], [82, 86), [90, 99)

// humidity-to-location map:
// 60 56 37
// 56 93 4

// [46, 56), [56, 60), [60, 61), [82, 85), [86, 90), [94, 97), [97, 99)

function mappingFn(range, maps) {
  let R = [...range];
  const A = [];
  maps.forEach((map) => {
    const NR = [];
    const [dst, src, len] = map;
    while (R.length) {
      const [a, b] = R.pop();

      // max a, because the range cannot be smaller than a
      // min b, because the range cannot be larger than b
      const back = [a, Math.min(src, b)];
      const middle = [Math.max(a, src), Math.min(src + len, b)];
      const forward = [Math.max(a, src + len), b];

      if (back[1] > back[0]) {
        NR.push(back);
      }
      if (middle[1] > middle[0]) {
        A.push([middle[0] - src + dst, middle[1] - src + dst]);
      }
      if (forward[1] > forward[0]) {
        NR.push(forward);
      }
    }
    R = NR;
  });

  return [...A, ...R];
}
