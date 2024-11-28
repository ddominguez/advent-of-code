import day06
import gleeunit/should

pub fn part1_test() {
  let input =
    "
turn on 0,0 through 999,999
toggle 0,0 through 999,0
"
  day06.part1(input) |> should.equal(999_000)
}
