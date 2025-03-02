async function main() {
  const input = Bun.file("../input/03.txt");
  const data = await input.text();
  if (Bun.argv[2] === "part2") {
    console.log(part2(data.trim()));
  } else {
    console.log(part1(data.trim()));
  }
}

type Position = {
  row: number;
  col: number;
};

const positions: Position[] = [
  { row: -1, col: 0 },
  { row: 0, col: 1 },
  { row: 1, col: 0 },
  { row: 0, col: -1 },
  { row: -1, col: -1 },
  { row: -1, col: 1 },
  { row: 1, col: 1 },
  { row: 1, col: -1 },
];

const isSpecialChar = (char: string): boolean => {
  return !/^[0-9.]+$/.test(char);
};

export function part1(data: string) {
  let result = 0;
  const lines = data.split("\n");
  lines.forEach((line, row) => {
    const trackedIdxNum = new Map<number, string>();
    line.split("").forEach((char, col) => {
      if (Number.parseInt(char) >= 0) {
        trackedIdxNum.set(col, char);
        if (col < line.length - 1) {
          return;
        }
      }

      if (trackedIdxNum.size === 0) {
        return;
      }

      let hasAdjacent = false;
      for (const trackedCol of trackedIdxNum.keys()) {
        if (hasAdjacent) {
          break;
        }
        for (const pos of positions) {
          if (
            lines[row + pos.row] === undefined ||
            lines[row + pos.row][trackedCol + pos.col] === undefined
          ) {
            continue;
          }

          if (isSpecialChar(lines[row + pos.row][trackedCol + pos.col])) {
            const partNum = Number([...trackedIdxNum.values()].join(""));
            result += partNum;
            hasAdjacent = true;
          }
        }
      }

      trackedIdxNum.clear();
    });
  });
  return result;
}

export function part2(data: string) {
  let result = 0;
  const specialNumbers = new Map<string, number[]>();
  const lines = data.split("\n");
  lines.forEach((line, row) => {
    const trackedIdxNum = new Map<number, string>();
    line.split("").forEach((char, col) => {
      if (Number.parseInt(char) >= 0) {
        trackedIdxNum.set(col, char);
        if (col < line.length - 1) {
          return;
        }
      }

      if (trackedIdxNum.size === 0) {
        return;
      }

      let hasAdjacent = false;
      for (const trackedCol of trackedIdxNum.keys()) {
        if (hasAdjacent) {
          break;
        }
        for (const pos of positions) {
          if (
            lines[row + pos.row] === undefined ||
            lines[row + pos.row][trackedCol + pos.col] === undefined
          ) {
            continue;
          }

          if (lines[row + pos.row][trackedCol + pos.col] === "*") {
            const partNum = Number([...trackedIdxNum.values()].join(""));
            const key = `${row + pos.row},${trackedCol + pos.col}`;
            const existingNums = specialNumbers.get(key);
            if (existingNums) {
              specialNumbers.set(key, [...existingNums, partNum]);
            } else {
              specialNumbers.set(key, [partNum]);
            }
            hasAdjacent = true;
          }
        }
      }

      trackedIdxNum.clear();
    });
  });

  result = specialNumbers
    .values()
    .filter((v) => v.length === 2)
    .map((v) => v.reduce((acc, cur) => acc * cur, 1))
    .reduce((acc, cur) => acc + cur, 0);
  return result;
}

if (import.meta.path === Bun.main) {
  main();
}
