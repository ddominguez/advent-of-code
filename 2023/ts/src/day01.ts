async function main() {
  const input = Bun.file("../input/01.txt");
  const data = await input.text();
  if (Bun.argv[2] === "part2") {
    console.log(part2(data.trim()));
  } else {
    console.log(part1(data.trim()));
  }
}

function part1(data: string) {
  let result = 0;
  for (const line of data.split("\n")) {
    const found: string[] = [];
    for (const char of line.split("")) {
      if (Number.parseInt(char)) {
        found.push(char);
      }
    }
    result += Number.parseInt(found[0] + found[found.length - 1]);
  }
  return result;
}

function part2(data: string) {
  let result = 0;
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
  for (const line of data.split("\n")) {
    const found: string[] = [];
    let start = 0;
    line.split("").forEach((char, i) => {
      if (Number.parseInt(char)) {
        found.push(char);
      }

      const wordSubstr = line.substring(start, i + 1);
      if (wordSubstr.length < 3) return;
      for (const [k, v] of numbers.entries()) {
        if (wordSubstr.includes(k) && numbers.has(k)) {
          found.push(v);
          start = i;
        }
      }
    });
    result += Number.parseInt(found[0] + found[found.length - 1]);
  }
  return result;
}

if (import.meta.path === Bun.main) {
  main();
}

export { part1, part2 };
