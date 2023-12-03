import { readFile } from "node:fs";
import { dirname } from "node:path";

const rootDir = process.cwd();

readFile(`${rootDir}/2023/input/02.txt`, "utf-8", (err, data) => {
  if (err) {
    console.log(err);
    return;
  }

  const isPart1 = true;
  const partNum = isPart1 ? 1 : 2;

  let results: number[] = [];

  const maxColorCounts = new Map([
    ["red", 12],
    ["green", 13],
    ["blue", 14],
  ]);

  data
    .trim()
    .split("\n")
    .forEach((line) => {
      const [game, sets] = line.split(":");
      let isPossible = true;
      const minColorCount = new Map<string, number>([
        ["red", 0],
        ["green", 0],
        ["blue", 0],
      ]);
      sets
        .trim()
        .split(";")
        .forEach((set) => {
          if (!isPossible) {
            return;
          }
          set
            .trim()
            .split(", ")
            .forEach((cube) => {
              const [count, color] = cube.split(" ");
              const countInt = parseInt(count);
              if (isPart1) {
                const maxCount = maxColorCounts.get(color);
                if (maxCount && maxCount < countInt) {
                  isPossible = false;
                  return;
                }
              } else {
                const minCount = minColorCount.get(color);
                if (minCount !== undefined) {
                  minColorCount.set(color, Math.max(minCount, countInt));
                }
              }
            });
        });

      if (isPart1 && isPossible) {
        results.push(parseInt(game.replace(/\D/g, "")));
      }

      if (!isPart1) {
        const [c1, c2, c3] = Array.from(minColorCount.values());
        results.push(c1 * c2 * c3);
      }
    });

  console.log(
    `part ${partNum} result:`,
    results.reduce((acc, currVal) => acc + currVal, 0),
  );
});
