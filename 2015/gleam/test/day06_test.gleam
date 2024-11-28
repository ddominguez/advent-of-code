import day06
import gleeunit/should

// Note: did not use all examples.
// examples will slow down tests.

pub fn part1_test() {
  let input =
    "
toggle 0,0 through 999,0
"
  day06.part1(input) |> should.equal(1000)
}

pub fn part2_test() {
  let input =
    "
turn on 0,0 through 0,0
"
  day06.part2(input) |> should.equal(1)
}
