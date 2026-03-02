use std::collections::HashSet;

use utils::get_input_lines;

#[derive(Clone, Copy)]
enum Direction {
    North,
    East,
    South,
    West,
}

impl Direction {
    fn turn_right(self) -> Self {
        match self {
            Direction::North => Direction::East,
            Direction::East => Direction::South,
            Direction::South => Direction::West,
            Direction::West => Direction::North,
        }
    }

    fn turn_left(self) -> Self {
        match self {
            Direction::North => Direction::West,
            Direction::East => Direction::North,
            Direction::South => Direction::East,
            Direction::West => Direction::South,
        }
    }

    fn move_position(self) -> (i32, i32) {
        match self {
            Direction::North => (0, 1),
            Direction::East => (1, 0),
            Direction::South => (0, -1),
            Direction::West => (-1, 0),
        }
    }
}

fn main() {
    let day = 1;
    let input = get_input_lines(day);
    for arg in std::env::args().skip(1) {
        match arg.as_str() {
            "-part1" => println!("{}", part1(&input)),
            "-part2" => println!("{}", part2(&input)),
            _ => eprintln!("Unknown argument: {}", arg),
        }
    }
}

fn part1(input: &str) -> i32 {
    let mut pos = (0, 0);
    let mut dir = Direction::North;

    for seq in input.split(",").map(str::trim) {
        let (turn, blocks_str) = seq.split_at(1);
        let blocks: i32 = blocks_str.parse().expect("Invalid blocks amount");

        dir = match turn {
            "R" => dir.turn_right(),
            "L" => dir.turn_left(),
            _ => panic!("Invalid turn: {}", turn),
        };

        let (x, y) = dir.move_position();
        pos.0 += x * blocks;
        pos.1 += y * blocks;
    }

    pos.0.abs() + pos.1.abs()
}

fn part2(input: &str) -> i32 {
    let mut pos = (0, 0);
    let mut dir = Direction::North;
    let mut visited: HashSet<(i32, i32)> = HashSet::new();
    visited.insert(pos);

    for seq in input.split(",").map(str::trim) {
        let (turn, blocks_str) = seq.split_at(1);
        let blocks: i32 = blocks_str.parse().expect("Invalid blocks amount");

        dir = match turn {
            "R" => dir.turn_right(),
            "L" => dir.turn_left(),
            _ => panic!("Invalid turn: {}", turn),
        };

        let (x, y) = dir.move_position();
        for _ in 0..blocks {
            pos.0 += x;
            pos.1 += y;
            if !visited.insert(pos) {
                return pos.0.abs() + pos.1.abs();
            }
        }
    }

    pos.0.abs() + pos.1.abs()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let cases = [("R2, L3", 5), ("R2, R2, R2", 2), ("R5, L5, R5, R3", 12)];

        for (input, expected) in cases {
            let result = part1(input);
            assert_eq!(result, expected, "got {}, expected {}", result, expected);
        }
    }

    #[test]
    fn test_part2() {
        let input = "R8, R4, R4, R8";
        let expected = 4;
        let result = part2(input);
        assert_eq!(result, expected, "got {}, expected {}", result, expected);
    }
}
