import { expect, test } from "bun:test";
import { part1, part2 } from "../src/day03";

test("part1", () => {
  const input = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`;
  const result = part1(input.trim());
  const expected = 4361;
  expect(result).toBe(expected);
});

test("part2", () => {
  const input = `
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`;
  const result = part2(input.trim());
  const expected = 467835;
  expect(result).toBe(expected);
});
