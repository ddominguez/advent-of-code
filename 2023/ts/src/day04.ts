async function main() {
  const input = Bun.file("../input/04.txt");
  const data = await input.text();
  if (Bun.argv[2] === "part2") {
    console.log(part2(data.trim()));
  } else {
    console.log(part1(data.trim()));
  }
}

export function part1(data: string) {
  let result = 0;
  for (const line of data.split("\n")) {
    const [_, cardNumbers] = line.split(": ");
    const [left, right] = cardNumbers.split(" | ");
    const winners = new Set(left.split(" ").filter(Boolean).map(Number));
    const matches = right
      .split(" ")
      .filter(Boolean)
      .map(Number)
      .filter((v) => winners.has(v));

    result += matches.reduce((acc, _, i) => {
      if (i === 0) return acc + 1;
      return acc * 2;
    }, 0);
  }
  return result;
}

export function part2(data: string) {
  let result = 0;
  let cardNumber = 0;
  const cardsProcessed: Record<number, number> = {};
  for (const line of data.split("\n")) {
    cardNumber += 1;
    const [_, cardNumbers] = line.split(": ");
    const [left, right] = cardNumbers.split(" | ");
    const winners = new Set(left.split(" ").filter(Boolean).map(Number));
    const matches = right
      .split(" ")
      .filter(Boolean)
      .map(Number)
      .filter((v) => winners.has(v));

    cardsProcessed[cardNumber] = (cardsProcessed[cardNumber] ?? 0) + 1;
    for (let i = 0; i < matches.length; i++) {
      const key = cardNumber + i + 1;
      cardsProcessed[key] =
        (cardsProcessed[key] ?? 0) + cardsProcessed[cardNumber];
    }
  }
  result = Object.values(cardsProcessed).reduce((acc, cur) => acc + cur, 0);
  return result;
}

if (import.meta.path === Bun.main) {
  main();
}
