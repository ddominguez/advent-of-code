import aoc
import argv
import gleam/dict.{type Dict}
import gleam/io
import gleam/set.{type Set}
import gleam/string

type Position =
  Dict(String, Int)

type Visited =
  Set(List(Int))

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/03.txt")
  let contents = string.trim(contents)
  case argv.load().arguments {
    ["part2"] -> part2(contents)
    _ -> part1(contents)
  }
  |> io.debug
}

pub fn part1(input: String) -> Int {
  let initial_visited = set.from_list([[0, 0]])
  let initial_position = dict.from_list([#("x", 0), #("y", 0)])
  string.to_graphemes(input) |> house_visits(initial_visited, initial_position)
}

fn update_position(position: Position, key: String, bump_val: Int) -> Position {
  let assert Ok(pos_val) = dict.get(position, key)
  dict.insert(position, key, pos_val + bump_val)
}

fn update_visited(visited: Visited, position: Position) -> Visited {
  let pos_vals = dict.values(position)
  case set.contains(visited, pos_vals) {
    True -> visited
    False -> set.insert(visited, pos_vals)
  }
}

fn house_visits(input: List(String), visited: Visited, position: Position) {
  case input {
    [] -> set.size(visited)
    _ -> {
      let assert [head, ..rest] = input
      let new_position = case head {
        "^" -> update_position(position, "y", 1)
        ">" -> update_position(position, "x", 1)
        "v" -> update_position(position, "y", -1)
        "<" -> update_position(position, "x", -1)
        _ -> panic
      }
      house_visits(rest, update_visited(visited, new_position), new_position)
    }
  }
}

pub fn part2(_input: String) -> Int {
  io.debug("Todo")
  -100
}
