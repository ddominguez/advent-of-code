import day03
import gleeunit/should

pub fn part1_test() {
  day03.part1(">") |> should.equal(2)
  day03.part1("^>v<") |> should.equal(4)
  day03.part1("^v^v^v^v^v") |> should.equal(2)
}
