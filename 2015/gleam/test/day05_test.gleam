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

pub fn part2_test() {
  let input =
    "
aaaa
aaa
"
    |> string.trim()

  day05.part2(input) |> should.equal(1)

  let input =
    "
qjhvhtzxzqqjkmpb
xxyxx
uurcxstgmygtbstg
ieodomkazucvgmuy
"
    |> string.trim()

  day05.part2(input) |> should.equal(2)
}
