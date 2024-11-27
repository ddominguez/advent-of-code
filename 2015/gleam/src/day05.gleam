import aoc
import argv
import gleam/io
import gleam/list
import gleam/string

const vowels = ["a", "e", "i", "o", "u"]

const invalid = ["ab", "cd", "pq", "xy"]

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/05.txt")
  let contents = string.trim(contents)
  case argv.load().arguments {
    ["part2"] -> part2(contents)
    _ -> part1(contents)
  }
  |> io.debug
}

pub fn part1(input: String) -> Int {
  string.split(input, "\n")
  |> list.fold(0, fn(count, line) {
    let line_is_nice = is_nice(string.to_graphemes(line), 0, 0, 0)
    increment(count, when: line_is_nice)
  })
}

fn increment(value: Int, when predicate: Bool) {
  case predicate {
    True -> value + 1
    False -> value
  }
}

fn is_nice(
  chars: List(String),
  vowels_count: Int,
  doubles_count: Int,
  invalid_count: Int,
) -> Bool {
  case chars {
    [] -> vowels_count >= 3 && doubles_count > 0 && invalid_count == 0
    [char] -> {
      is_nice(
        [],
        increment(vowels_count, when: list.contains(vowels, char)),
        doubles_count,
        invalid_count,
      )
    }
    [char, ..rest] -> {
      let assert Ok(second) = list.first(rest)
      is_nice(
        rest,
        increment(vowels_count, when: list.contains(vowels, char)),
        increment(doubles_count, when: char == second),
        increment(invalid_count, when: list.contains(invalid, char <> second)),
      )
    }
  }
}

pub fn part2(input: String) -> Int {
  string.split(input, "\n")
  |> list.fold(0, fn(count, line) {
    let line_is_nice = is_nice_too(string.to_graphemes(line), 0, 0)
    increment(count, when: line_is_nice)
  })
}

fn is_nice_too(chars: List(String), pairs: Int, repeats: Int) -> Bool {
  case chars {
    [] -> pairs > 0 && repeats > 0
    [_] | [_, _] -> is_nice_too([], pairs, repeats)
    [a, b, c, ..rest] -> {
      let pairs_exist = string.join([c, ..rest], "") |> string.contains(a <> b)
      is_nice_too(
        [b, c, ..rest],
        increment(pairs, when: pairs_exist),
        increment(repeats, when: a == c),
      )
    }
  }
}
