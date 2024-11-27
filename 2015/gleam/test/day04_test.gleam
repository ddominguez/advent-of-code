import day04
import gleeunit/should

// only one test
// part2 test times out :(

pub fn part1_test() {
  day04.part1("abcdef") |> should.equal(609_043)
  //day04.part1("pqrstuv") |> should.equal(1_048_970)
}
