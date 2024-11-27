import day05
import gleam/string
import gleeunit/should

pub fn part1_test() {
  let input =
    "
ugknbfddgicrmopn
aaa
jchzalrnumimnmhp
haegwjzuvuyypxyu
dvszwmarrgswjxmb
"
    |> string.trim

  day05.part1(input) |> should.equal(2)
}
