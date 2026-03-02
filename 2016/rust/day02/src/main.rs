use utils::get_input_lines;

fn main() {
    let day = 2;
    let input = get_input_lines(day);
    for arg in std::env::args().skip(1) {
        match arg.as_str() {
            "-part1" => println!("{}", part1(&input)),
            "-part2" => println!("{}", part2(&input)),
            _ => eprintln!("Unknown argument: {}", arg),
        }
    }
}

fn is_in_bounds(pos: (i32, i32), rows: i32, cols: i32) -> bool {
    pos.0 >= 0 && pos.0 < rows && pos.1 >= 0 && pos.1 < cols
}

fn get_new_pos(pos: (i32, i32), move_char: char) -> Result<(i32, i32), String> {
    let mut new_pos = pos;
    match move_char {
        'U' => new_pos.0 -= 1,
        'R' => new_pos.1 += 1,
        'D' => new_pos.0 += 1,
        'L' => new_pos.1 -= 1,
        _ => return Err(format!("Invalid move character: {}", move_char)),
    }
    Ok(new_pos)
}

fn part1(input: &str) -> String {
    let mut result = String::new();
    let keypad: [[char; 3]; 3] = [['1', '2', '3'], ['4', '5', '6'], ['7', '8', '9']];
    const ROWS: i32 = 3;
    const COLS: i32 = 3;
    let mut btn_pos: (i32, i32) = (1, 1);

    for line in input.lines() {
        for c in line.chars() {
            match get_new_pos(btn_pos, c) {
                Ok(new_pos) => {
                    if is_in_bounds(new_pos, ROWS, COLS) {
                        btn_pos = new_pos;
                    }
                }
                Err(e) => eprintln!("{}", e),
            }
        }
        result.push(keypad[btn_pos.0 as usize][btn_pos.1 as usize])
    }

    result
}

fn part2(input: &str) -> String {
    let mut result = String::new();
    let keypad: [[char; 5]; 5] = [
        [' ', ' ', '1', ' ', ' '],
        [' ', '2', '3', '4', ' '],
        ['5', '6', '7', '8', '9'],
        [' ', 'A', 'B', 'C', ' '],
        [' ', ' ', 'D', ' ', ' '],
    ];
    const ROWS: i32 = 5;
    const COLS: i32 = 5;
    const EDGE: char = ' ';
    let mut btn_pos: (i32, i32) = (2, 0);

    for line in input.lines() {
        for c in line.chars() {
            match get_new_pos(btn_pos, c) {
                Ok(new_pos) => {
                    if is_in_bounds(new_pos, ROWS, COLS)
                        && keypad[new_pos.0 as usize][new_pos.1 as usize] != EDGE
                    {
                        btn_pos = new_pos;
                    }
                }
                Err(e) => eprintln!("{}", e),
            }
        }
        result.push(keypad[btn_pos.0 as usize][btn_pos.1 as usize])
    }

    result
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_get_new_pos_ok() -> Result<(), String> {
        let pos = (0, 0);
        let cases = [('U', (-1, 0)), ('R', (0, 1)), ('D', (1, 0)), ('L', (0, -1))];
        for (move_char, expected) in cases {
            let result = get_new_pos(pos, move_char)?;
            assert_eq!(result, expected);
        }
        Ok(())
    }

    #[test]
    fn test_get_new_pos_err() {
        let bad_char = 'X';
        let result = get_new_pos((0, 0), bad_char);
        assert!(result.is_err())
    }

    #[test]
    fn test_part1() {
        let input = "ULL\nRRDDD\nLURDL\nUUUUD";
        let result = part1(input);
        let expected = "1985";
        assert_eq!(result, expected);
    }

    #[test]
    fn test_part2() {
        let input = "ULL\nRRDDD\nLURDL\nUUUUD";
        let result = part2(input);
        let expected = "5DB3";
        assert_eq!(result, expected);
    }
}
