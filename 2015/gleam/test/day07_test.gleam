import day07
import gleeunit/should

pub fn part1_test() {
  let input =
    "
x AND y -> d
x OR y -> e
NOT x -> h
NOT y -> i
123 -> x
y RSHIFT 2 -> g
x LSHIFT 2 -> f
456 -> y
"

  day07.part1(input, "g") |> should.equal(114)
}
