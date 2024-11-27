import aoc
import argv
import gleam/bool
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

fn get_new_position(position: Position, direction: String) -> Position {
  case direction {
    "^" -> update_position(position, "y", 1)
    ">" -> update_position(position, "x", 1)
    "v" -> update_position(position, "y", -1)
    "<" -> update_position(position, "x", -1)
    _ -> panic as "Unknown value"
  }
}

fn house_visits(input: List(String), visited: Visited, position: Position) {
  case input {
    [] -> set.size(visited)
    [head, ..rest] -> {
      let new_position = get_new_position(position, head)
      house_visits(rest, update_visited(visited, new_position), new_position)
    }
  }
}

type Visitor {
  Visitor(visits: Visited, position: Position)
}

pub fn part2(input: String) -> Int {
  let initial_visited = set.from_list([[0, 0]])
  let initial_position = dict.from_list([#("x", 0), #("y", 0)])
  let santa_state = Visitor(visits: initial_visited, position: initial_position)
  let robo_state = Visitor(visits: initial_visited, position: initial_position)
  let state = dict.from_list([#("santa", santa_state), #("robo", robo_state)])
  let is_santa = True
  string.to_graphemes(input) |> robo_santa_visits(state, is_santa)
}

fn robo_or_santa(is_santa: Bool) -> String {
  case is_santa {
    True -> "santa"
    False -> "robo"
  }
}

fn robo_santa_visits(
  input: List(String),
  state: Dict(String, Visitor),
  is_santa: Bool,
) {
  case input {
    [] -> {
      let assert Ok(santa) = dict.get(state, "santa")
      let assert Ok(robo) = dict.get(state, "robo")
      santa.visits |> set.union(robo.visits) |> set.size
    }
    [head, ..rest] -> {
      let visitor = robo_or_santa(is_santa)
      let assert Ok(visitor_state) = dict.get(state, visitor)
      let new_position = get_new_position(visitor_state.position, head)
      let updated_visits = update_visited(visitor_state.visits, new_position)
      let visitor_state =
        Visitor(visits: updated_visits, position: new_position)
      let new_state = dict.insert(state, visitor, visitor_state)
      robo_santa_visits(rest, new_state, bool.negate(is_santa))
    }
  }
}
