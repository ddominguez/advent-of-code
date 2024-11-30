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
  //let input = "\"byc\\x9dyxuafof\\\\\\xa6uf\\\\axfozomj\\\\olh\\x6a\""

  day08.part1(input) |> should.equal(12)
}
