import aoc
import argv
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

const max_signal = 65_535

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/07.txt")
  case argv.load().arguments {
    ["part2"] -> part2(contents, "a")
    _ -> part1(contents, "a")
  }
  |> io.debug
}

pub fn part1(input: String, wire: String) -> Int {
  let lines =
    string.trim(input)
    |> string.split("\n")
    |> list.map(string.split(_, " "))

  find_wire(lines, wire, dict.new(), [])
}

pub fn part2(input: String, wire: String) -> Int {
  let lines =
    string.trim(input)
    |> string.split("\n")
    |> list.map(string.split(_, " "))

  let wire_val = find_wire(lines, wire, dict.new(), [])
  let new_state = dict.new() |> dict.insert("b", wire_val)
  find_wire(lines, wire, new_state, [])
}

fn find_wire(
  lines: List(List(String)),
  wire: String,
  map: Dict(String, Int),
  processed: List(List(String)),
) -> Int {
  case dict.get(map, wire) {
    Ok(val) -> val
    Error(_) -> {
      case lines {
        [] -> find_wire(processed, wire, map, [])
        [head, ..tail] -> {
          let updated_map = update_map(map, head)
          find_wire(tail, wire, updated_map, [head, ..processed])
        }
      }
    }
  }
}

fn update_map(map: Dict(String, Int), line: List(String)) {
  case line {
    [a, _, key] -> {
      case int.parse(a), dict.get(map, a), dict.get(map, key) {
        Ok(num), _, Error(_) -> dict.insert(map, key, num)
        Error(_), Ok(val), _ -> dict.insert(map, key, val)
        _, _, _ -> map
      }
    }
    [a, "AND", b, _, key] -> {
      case int.parse(a), dict.get(map, a), dict.get(map, b) {
        Ok(a_val), _, Ok(b_val) ->
          dict.insert(map, key, int.bitwise_and(a_val, b_val))
        Error(_), Ok(a_val), Ok(b_val) ->
          dict.insert(map, key, int.bitwise_and(a_val, b_val))
        _, _, _ -> map
      }
    }
    [a, "OR", b, _, key] -> {
      case dict.get(map, a), dict.get(map, b) {
        Ok(a_val), Ok(b_val) ->
          dict.insert(map, key, int.bitwise_or(a_val, b_val))
        _, _ -> map
      }
    }
    [a, "LSHIFT", b, _, key] -> {
      case dict.get(map, a), int.parse(b) {
        Ok(a_val), Ok(b_val) ->
          dict.insert(map, key, int.bitwise_shift_left(a_val, b_val))
        _, _ -> map
      }
    }
    [a, "RSHIFT", b, _, key] -> {
      case dict.get(map, a), int.parse(b) {
        Ok(a_val), Ok(b_val) ->
          dict.insert(map, key, int.bitwise_shift_right(a_val, b_val))
        _, _ -> map
      }
    }
    ["NOT", a, _, key] -> {
      case dict.get(map, a) {
        Ok(a_val) -> {
          let new_val = int.bitwise_not(a_val) |> int.bitwise_and(max_signal)
          dict.insert(map, key, new_val)
        }
        _ -> map
      }
    }
    _ -> map
  }
}
