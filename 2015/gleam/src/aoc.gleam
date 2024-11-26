import gleam/io

pub fn main() {
  io.println("Hello from aoc!")
}

@external(erlang, "file", "read_file")
pub fn read_file(file_path: String) -> Result(String, error)
