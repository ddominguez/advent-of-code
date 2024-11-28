import aoc
import argv
import gleam/bool
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
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
    _ -> dict.get(acc, [x, y]) |> result.unwrap(False) |> bool.negate
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

pub fn part2(input: String) -> Int {
  string.trim(input)
  |> string.split("\n")
  |> list.map(parse)
  |> list.fold(dict.new(), adjust_brightness)
  |> dict.fold(0, fn(acc, _key, val) { acc + val })
}

type BrightnessGrid =
  Dict(List(Int), Int)

fn adjust_brightness(
  grid: BrightnessGrid,
  instruction: Instruction,
) -> BrightnessGrid {
  let assert #(action, [x1, y1], [x2, y2]) = instruction
  let x_range = list.range(x1, x2)
  let y_range = list.range(y1, y2)

  use acc, x <- list.fold(x_range, grid)
  use acc, y <- list.fold(y_range, acc)
  let curr_brightness = dict.get(acc, [x, y]) |> result.unwrap(0)
  let brightness = case action {
    "on" -> curr_brightness + 1
    "off" -> int.max(curr_brightness - 1, 0)
    _ -> curr_brightness + 2
  }
  dict.insert(acc, [x, y], brightness)
}
