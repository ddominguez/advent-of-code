import aoc
import argv
import gleam/bit_array
import gleam/crypto
import gleam/int
import gleam/io
import gleam/string

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/04.txt")
  let contents = string.trim(contents)
  case argv.load().arguments {
    ["part2"] -> part2(contents)
    _ -> part1(contents)
  }
  |> io.debug
}

pub fn part1(input: String) -> Int {
  lowest_num_with_leading_zeroes(input, "00000", 0)
}

pub fn part2(input: String) -> Int {
  lowest_num_with_leading_zeroes(input, "000000", 0)
}

fn lowest_num_with_leading_zeroes(
  input: String,
  leading: String,
  acc: Int,
) -> Int {
  let encoded_hash =
    bit_array.from_string(input <> int.to_string(acc))
    |> crypto.hash(crypto.Md5, _)
    |> bit_array.base16_encode
  case string.starts_with(encoded_hash, leading) {
    True -> acc
    False -> lowest_num_with_leading_zeroes(input, leading, acc + 1)
  }
}
