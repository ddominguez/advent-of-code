<?php
declare(strict_types=1);

require 'aoc.php';

enum Direction: int {
    case North = 0;
    case East = 1;
    case South = 2;
    case West = 3;

    public function turn_left(): self {
        return match($this) {
            Direction::North => Direction::West,
            Direction::East => Direction::North,
            Direction::South => Direction::East,
            Direction::West => Direction::South,
        };
    }

    public function turn_right(): self {
        return match($this) {
            Direction::North => Direction::East,
            Direction::East => Direction::South,
            Direction::South => Direction::West,
            Direction::West => Direction::North,
        };
    }

    public function move_position(): array {
        return match($this) {
            Direction::North => [0, 1],
            Direction::East => [1, 0],
            Direction::South =>  [0, -1],
            Direction::West => [-1, 0],
        };
    }

    public function turn_direction(string $turn): self {
        if ($turn === "R") {
            return self::turn_left();
        }
        return self::turn_right();
    }
}

function part1(string $contents): int {
    $position = [0, 0];
    $direction = Direction::North;

    foreach (array_map('trim', explode(',', $contents)) as $seq) {
        $direction = $direction->turn_direction(substr($seq, 0, 1));
        $blocks = (int)substr($seq, 1);
        [$x, $y] = $direction->move_position();
        $position = [$position[0] + $x * $blocks, $position[1] + $y * $blocks];
    }

    return abs($position[0]) + abs($position[1]);
}

function part2(string $contents): int {
    $position = [0, 0];
    $visited = [];
    $visited["$position[0],$position[1]"] = true;
    $direction = Direction::North;

    foreach (array_map('trim', explode(',', $contents)) as $seq) {
        $direction = $direction->turn_direction(substr($seq, 0, 1));
        $blocks = (int)substr($seq, 1);
        [$x, $y] = $direction->move_position();

        for ($i=0; $i<$blocks; $i++) {
            $position = [$position[0] + $x, $position[1] + $y];
            if (array_key_exists("$position[0],$position[1]", $visited)) {
                return abs($position[0]) + abs($position[1]);
            }
            $visited["$position[0],$position[1]"] = true;
        }
    }

    return abs($position[0]) + abs($position[1]);
}

function test_part1(): void {
    $input = 'R2, L3';
    expect(part1($input), 5, __FUNCTION__);
}

function test_part2(): void {
    $input = 'R8, R4, R4, R8';
    expect(part2($input), 4, __FUNCTION__);
}

main([
    'cmd'=> $argv[1] ?? 'part1',
    'day'=> 1,
    'part1'=> 'part1',
    'part2'=> 'part2',
    'tests'=> function() {
        test_part1();
        test_part2();
    },
]);
