import day01
import gleeunit/should

pub fn part1_test() {
  day01.part1("(())") |> should.equal(0)
  day01.part1("(()(()(") |> should.equal(3)
  day01.part1(")())())") |> should.equal(-3)
}

pub fn part2_test() {
  day01.part2(")") |> should.equal(1)
  day01.part2("()())") |> should.equal(5)
}

