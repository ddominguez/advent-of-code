import { readFile } from "node:fs";
import { dirname } from "node:path";

const rootDir = process.cwd();

const numbers = new Map([
  ["one", "1"],
  ["two", "2"],
  ["three", "3"],
  ["four", "4"],
  ["five", "5"],
  ["six", "6"],
  ["seven", "7"],
  ["eight", "8"],
  ["nine", "9"],
]);

readFile(`${rootDir}/2023/input/01.txt`, "utf-8", (err, data) => {
  if (err) {
    console.log(err);
    return;
  }

  const isPart1 = true;
  let result = 0;

  data
    .trim()
    .split("\n")
    .forEach((line) => {
      const found: string[] = [];
      let start = 0;
      line.split("").forEach((char, i) => {
        if (parseInt(char)) {
          found.push(char);
          start = i;
        }

        if (!isPart1) {
          const wordSubstr = line.substring(start, i + 1);
          if (wordSubstr.length < 3) return;
          for (const [k, v] of numbers.entries()) {
            if (wordSubstr.includes(k) && numbers.has(k)) {
              found.push(v);
              start = i;
            }
          }
        }
      });
      result += parseInt(found[0] + found[found.length - 1]);
    });

  console.log(`part ${isPart1 ? 1 : 2} result:`, result);
});
