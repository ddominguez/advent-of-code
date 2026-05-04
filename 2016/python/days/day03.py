import sys
from typing import Iterable


def part1(data: str) -> int:
    result = 0

    for line in data.splitlines():
        sides = [int(s) for s in line.split()]
        if is_triangle(sides):
            result += 1

    return result


def part2(data: str) -> int:
    result = 0
    group: list[list[int]] = []
    group_max_len = 3

    for line in data.splitlines():
        sides = [int(s) for s in line.split()]
        group.append(sides)
        if len(group) == group_max_len:
            for s in zip(*group):
                if is_triangle(s):
                    result += 1
            group.clear()

    return result


def is_triangle(sides: Iterable[int]) -> bool:
    a, b, c = sides
    return a + b > c and c + b > a and a + c > b


if __name__ == "__main__":
    with open("../input/03.txt") as f:
        data = f.read().strip()

    if sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
