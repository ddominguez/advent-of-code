use std::fs::File;
use std::io::{BufRead, BufReader};
use std::path::PathBuf;

pub fn get_input_lines(day: u8) -> std::io::Lines<BufReader<File>> {
    let mut path = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
    path.pop();
    path.pop();
    path.push(format!("input/{:02}.txt", day));
    let file = File::open(path).expect("Failed to open input file");
    BufReader::new(file).lines()
}
