use utils::get_input_lines;

fn main() {
    let day = 3;
    let input = get_input_lines(day);
    for arg in std::env::args().skip(1) {
        match arg.as_str() {
            "-part1" => println!("{}", part1(&input)),
            "-part2" => println!("{}", part2(&input)),
            _ => eprintln!("Unknown argument: {}", arg),
        }
    }
}

fn get_sides(data: &str) -> Vec<i32> {
    data.split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect()
}

fn is_triangle(s1: i32, s2: i32, s3: i32) -> bool {
    s1 + s2 > s3 && s2 + s3 > s1 && s1 + s3 > s2
}

fn part1(input: &str) -> i32 {
    let mut res = 0;
    for line in input.lines() {
        let sides = get_sides(line);
        if is_triangle(sides[0], sides[1], sides[2]) {
            res += 1;
        }
    }
    res
}

fn part2(input: &str) -> i32 {
    let mut res = 0;
    let group_size: usize = 3;
    let mut group: Vec<Vec<i32>> = Vec::with_capacity(group_size);
    for line in input.lines() {
        let sides = get_sides(line);
        group.push(sides);
        if group.len() == group_size {
            let zipped: Vec<(i32, i32, i32)> = group[0]
                .iter()
                .zip(group[1].iter())
                .zip(group[2].iter())
                .map(|((a, b), c)| (*a, *b, *c))
                .collect();
            for (s1, s2, s3) in zipped {
                if is_triangle(s1, s2, s3) {
                    res += 1;
                }
            }
            group.clear();
        }
    }
    res
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_get_triangle() {
        let input = "5 10 25";
        let result = get_sides(input);
        assert_eq!(result, vec![5, 10, 25])
    }

    #[test]
    fn test_part1() {
        let input = "5 10 25\n541  588  421";
        let result = part1(input);
        assert_eq!(result, 1);
    }

    #[test]
    fn test_part2() {
        let input = "101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603";
        let result = part2(input);
        assert_eq!(result, 6);
    }
}
