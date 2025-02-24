import { readFile } from "node:fs";

readFile("../input/02.txt", "utf-8", (err, data) => {
  if (err) {
    console.log(err);
    return;
  }

  const isPart1 = true;
  let result = 0;

  const maxColorCounts = new Map([
    ["red", 12],
    ["green", 13],
    ["blue", 14],
  ]);

  data
    .trim()
    .split("\n")
    .forEach((line) => {
      const [game, sets] = line.split(": ");
      let isPossible = true;
      const minColorCount = new Map([
        ["red", 0],
        ["green", 0],
        ["blue", 0],
      ]);
      sets.split("; ").forEach((set) => {
        if (!isPossible) {
          return;
        }
        set.split(", ").forEach((cube) => {
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
        result += parseInt(game.replace(/\D/g, ""));
      } else if (!isPart1) {
        result += Array.from(minColorCount.values()).reduce(
          (acc, curr) => acc * curr,
          1,
        );
      }
    });

  console.log(`part ${isPart1 ? 1 : 2} result:`, result);
});
