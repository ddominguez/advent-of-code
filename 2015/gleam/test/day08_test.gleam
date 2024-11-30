import day08
import gleeunit/should

pub fn part1_test() {
  let input =
    "
\"\"
\"abc\"
\"aaa\\\"aaa\"
\"\\x27\"
"

  day08.part1(input) |> should.equal(12)
}

pub fn part2_test() {
  let input =
    "
\"\"
\"abc\"
\"aaa\\\"aaa\"
\"\\x27\"
"

  day08.part2(input) |> should.equal(19)
}
