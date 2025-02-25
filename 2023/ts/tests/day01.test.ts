import { expect, test } from "bun:test";
import { part1, part2 } from "../src/day01";

test("part1", () => {
  const input = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`;
  const result = part1(input.trim());
  const expected = 142;
  expect(result).toBe(expected);
});

test("part2", () => {
  const input = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`;
  const result = part2(input.trim());
  const expected = 281;
  expect(result).toBe(expected);
});
