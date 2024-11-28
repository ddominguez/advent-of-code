import aoc
import argv
import gleam/bool
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/string

type Grid =
  Dict(List(Int), Bool)

type Instruction =
  #(String, List(Int), List(Int))

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/06.txt")
  case argv.load().arguments {
    ["part2"] -> part2(contents)
    _ -> part1(contents)
  }
  |> io.debug
}

pub fn part1(input: String) -> Int {
  string.trim(input)
  |> string.split("\n")
  |> list.map(parse)
  |> list.fold(dict.new(), handle_instruction)
  |> dict.fold(0, fn(acc, _key, val) {
    case val {
      True -> acc + 1
      False -> acc
    }
  })
}

fn handle_instruction(grid: Grid, instruction: Instruction) {
  let assert #(action, [x1, y1], [x2, y2]) = instruction
  let x_range = list.range(x1, x2)
  let y_range = list.range(y1, y2)

  use acc, x <- list.fold(x_range, grid)
  use acc, y <- list.fold(y_range, acc)
  let on_off = case action {
    "on" -> True
    "off" -> False
    _ -> {
      case dict.get(acc, [x, y]) {
        Ok(val) -> bool.negate(val)
        Error(_) -> True
      }
    }
  }
  dict.insert(acc, [x, y], on_off)
}

fn to_ints_list(str: String) -> List(Int) {
  string.split(str, ",")
  |> list.map(fn(x) {
    let assert Ok(num) = int.parse(x)
    num
  })
}

fn parse(line: String) -> Instruction {
  case string.split(line, " ") {
    [_, "on", from, _, to] -> #("on", to_ints_list(from), to_ints_list(to))
    [_, "off", from, _, to] -> #("off", to_ints_list(from), to_ints_list(to))
    ["toggle", from, _, to] -> #("toggle", to_ints_list(from), to_ints_list(to))
    _ -> panic as "Unknown instruction"
  }
}

pub fn part2(_input: String) -> Int {
  -100
}
