import { readFileSync } from "node:fs";

const data = readFileSync("../input/03.txt", "utf-8");
const grid: string[][] = [];
const partNumbers: number[] = [];
const isSpecialChar = (char: string) => {
  if (!char) return false;
  return !/^[0-9.]+$/.test(char);
};

// return speacial char pos if tracked num is adjacent
const getAdacent = (row: number, pos: number[]): number[] | null => {
  const checkPos = [pos[0] - 1, ...pos, pos[pos.length - 1] + 1];
  let col: number | undefined;
  // curr row
  col = checkPos.find((i) => isSpecialChar(grid[row][i]));
  if (col) {
    return [row, col];
  }
  // prev row
  col = checkPos.find((i) => data[row - 1] && isSpecialChar(grid[row - 1][i]));
  if (col) {
    return [row - 1, col];
  }
  // next row
  col = checkPos.find((i) => data[row + 1] && isSpecialChar(grid[row + 1][i]));
  if (col) {
    return [row + 1, col];
  }
  return null;
};

data
  .trim()
  .split("\n")
  .forEach((line) => grid.push(Array.from(line)));

const spMap: Record<string, number[]> = {};
grid.forEach((row, i) => {
  let trackedNum = "";
  let trackedNumPos: number[] = [];
  row.forEach((char, j) => {
    let checkAdj = false;
    if (parseFloat(char) >= 0) {
      trackedNum += char;
      trackedNumPos.push(j);
      checkAdj = j === row.length - 1 && trackedNum.length > 0;
    } else {
      checkAdj = trackedNum.length > 0;
    }

    if (checkAdj) {
      const adjChar = getAdacent(i, trackedNumPos);
      if (adjChar) {
        partNumbers.push(Number(trackedNum));
        const [sr, sc] = adjChar;
        if (grid[sr][sc] === "*") {
          const key = `${sr},${sc}`;
          if (!Object.hasOwn(spMap, key)) {
            spMap[key] = [Number(trackedNum)];
          } else {
            spMap[key].push(Number(trackedNum));
          }
        }
      }
      trackedNum = "";
      trackedNumPos = [];
      checkAdj = false;
    }
  });
});

console.log(
  "part 1:",
  partNumbers.reduce((acc, curr) => acc + curr, 0),
);

const gearRatios: number[] = [];
for (const [k, v] of Object.entries(spMap)) {
  if (v.length == 2) {
    gearRatios.push(v.reduce((acc, curr) => acc * curr, 1));
  }
}
console.log(
  "part 2:",
  gearRatios.reduce((acc, curr) => acc + curr, 0),
);
