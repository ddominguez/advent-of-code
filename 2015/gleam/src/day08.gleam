import aoc
import argv
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
  |> list.fold(0, fn(acc, line) {
    let additional_quotes = -2
    let data_len = string.to_graphemes(line) |> get_data_len(additional_quotes)
    acc + string.length(line) - data_len
  })
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

pub fn part2(input: String) -> Int {
  input
  |> string.trim
  |> string.split("\n")
  |> list.fold(0, fn(acc, line) {
    let additional_quotes = 2
    let char_count =
      string.to_graphemes(line) |> get_char_count(additional_quotes)
    acc + char_count - string.length(line)
  })
}

fn get_char_count(line: List(String), count: Int) -> Int {
  case line {
    [] -> count
    ["\"", ..rest] -> get_char_count(rest, count + 2)
    ["\\", ..rest] -> get_char_count(rest, count + 2)
    [_, ..rest] -> get_char_count(rest, count + 1)
  }
}
