import aoc
import argv
import gleam/int
import gleam/io
import gleam/list
import gleam/string

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/08.txt")
  case argv.load().arguments {
    ["part2"] -> part2(contents)
    _ -> part1(contents)
  }
  |> io.debug
}

pub fn part1(input: String) -> Int {
  input
  |> string.trim
  |> string.split("\n")
  |> list.map(total_chars)
  |> int.sum
}

fn total_chars(line: String) -> Int {
  let character_len = string.length(line)
  let assert [_first, ..rest] = string.to_graphemes(line)
  let data_len = rest |> list.take(list.length(rest) - 1) |> get_data_len(0)

  character_len - data_len
}

fn get_data_len(chars: List(String), data_len: Int) -> Int {
  case chars {
    [] -> data_len
    ["\\", "x", _h1, _h2, ..tail] -> get_data_len(tail, data_len + 1)
    ["\\", "\"", ..tail] -> get_data_len(tail, data_len + 1)
    ["\\", "\\", ..tail] -> get_data_len(tail, data_len + 1)
    [_head, ..tail] -> get_data_len(tail, data_len + 1)
  }
}

pub fn part2(_input: String) -> Int {
  -100
}
