import aoc
import argv
import gleam/io

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/01.txt")
  case argv.load().arguments {
    ["part2"] -> part2(contents)
    _ -> part1(contents)
  }
  |> io.debug
}

pub fn part1(contents: String) -> Int {
  let assert Ok(floor) = get_floor(contents, 0)
  floor
}

fn get_floor(input: String, floor: Int) -> Result(Int, Nil) {
  case input {
    "(" <> rest -> get_floor(rest, floor + 1)
    ")" <> rest -> get_floor(rest, floor - 1)
    "" -> Ok(floor)
    _ -> Error(Nil)
  }
}

pub fn part2(contents: String) -> Int {
  let assert Ok(position) = get_position(contents, 0, 0)
  position
}

fn get_position(input: String, floor: Int, position: Int) -> Result(Int, Nil) {
  case input, floor {
    _input, -1 -> Ok(position)
    "(" <> rest, _floor -> get_position(rest, floor + 1, position + 1)
    ")" <> rest, _floor -> get_position(rest, floor - 1, position + 1)
    _, _ -> Error(Nil)
  }
}
