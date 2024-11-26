import aoc
import argv
import gleam/int
import gleam/io
import gleam/list
import gleam/string

pub fn main() {
  let assert Ok(contents) = aoc.read_file("../input/02.txt")
  case argv.load().arguments {
    ["part2"] -> string.trim(contents) |> part2
    _ -> string.trim(contents) |> part1
  }
  |> io.debug
}

pub fn part1(contents: String) -> Int {
  string.split(contents, "\n") |> get_total_feet(0)
}

fn to_dimensions_list(line: String) -> List(Int) {
  string.split(line, "x")
  |> list.map(fn(x) {
    let assert Ok(num) = int.parse(x)
    num
  })
}

fn get_total_feet(lines: List(String), total_feet: Int) {
  case lines {
    [] -> total_feet
    [line, ..rest] -> {
      let assert [l, w, h] = to_dimensions_list(line)
      let area = 2 * l * w + 2 * w * h + 2 * h * l
      let assert Ok(smallest_side) = list.reduce([l * w, w * h, h * l], int.min)
      let surface_area = area + smallest_side
      get_total_feet(rest, total_feet + surface_area)
    }
  }
}

pub fn part2(contents: String) -> Int {
  string.split(contents, "\n") |> get_ribbon_feet(0)
}

fn get_ribbon_feet(lines: List(String), total_feet: Int) -> Int {
  case lines {
    [] -> total_feet
    [line, ..rest] -> {
      let lwh = to_dimensions_list(line)
      let volume = int.product(lwh)
      let assert [a, b, _c] = list.sort(lwh, int.compare)
      let smallest_perimeter =
        [a, b] |> list.fold(0, fn(acc, x) { x + x + acc })
      get_ribbon_feet(rest, total_feet + volume + smallest_perimeter)
    }
  }
}
