use std::fs;
use std::path::PathBuf;

pub fn get_input_lines(day: u8) -> String {
    let mut path = PathBuf::from(env!("CARGO_MANIFEST_DIR"));
    path.pop();
    path.pop();
    path.push(format!("input/{:02}.txt", day));
    return fs::read_to_string(path).expect("Failed to read input");
}
