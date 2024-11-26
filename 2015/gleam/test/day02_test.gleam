import day02
import gleeunit/should

pub fn part1_test() {
  let input =
    "2x3x4
1x1x10"
  day02.part1(input) |> should.equal(101)
}

pub fn part2_test() {
  let input =
    "2x3x4
1x1x10"
  day02.part2(input) |> should.equal(48)
}
