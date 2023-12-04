import { readFile } from "node:fs";
import { rootPath } from "@utils";

type NumberPosition = {
  val: string;
  pos: number[];
};

readFile(`${rootPath}/2023/input/03.txt`, "utf-8", (err, data) => {
  if (err) {
    console.log(err);
    return;
  }

  const isPart1 = true;
  let result = 0;

  const numPositions: NumberPosition[][] = [];
  const symbolPos: number[][] = [];

  data
    .trim()
    .split("\n")
    .forEach((line, row) => {
      numPositions[row] = [];
      symbolPos[row] = [];
      let trackedNum = "";
      let trackNumPos: number[] = [];
      for (let i = 0; i < line.length; i++) {
        let trackIt = false;
        if (parseInt(line[i]) >= 0) {
          trackedNum += line[i];
          trackNumPos.push(i);
          if (i == line.length - 1 && trackedNum) {
            trackIt = true;
          }
        } else {
          if (/^[0-9.]+$/.test(line[i]) === false) {
            symbolPos[row].push(i);
          }
          if (trackedNum) {
            trackIt = true;
          }
        }

        if (trackIt) {
          numPositions[row].push({
            val: trackedNum,
            pos: trackNumPos,
          });
          trackedNum = "";
          trackNumPos = [];
        }
      }
    });

  numPositions.map((np, row) => {
    np.forEach((n) => {
      let isAdjacent = false;
      const updatedPos = [n.pos[0] - 1, ...n.pos, n.pos[n.pos.length - 1] + 1];

      if (symbolPos[row - 1]?.some((p) => updatedPos.includes(p))) {
        isAdjacent = true;
      } else if (symbolPos[row + 1]?.some((p) => updatedPos.includes(p))) {
        isAdjacent = true;
      } else if (symbolPos[row].some((p) => updatedPos.includes(p))) {
        isAdjacent = true;
      }
      if (isAdjacent) {
        result += Number(n.val);
      }
    });
  });
  console.log(`part ${isPart1 ? 1 : 2} result:`, result);
});
