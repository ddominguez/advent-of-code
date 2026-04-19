import sys

from enum import IntEnum


class Direction(IntEnum):
    North = 0
    East = 1
    South = 2
    West = 3

    def turn_left(self):
        return Direction((self.value - 1) % 4)

    def turn_right(self):
        return Direction((self.value + 1) % 4)

    def move_position(self):
        return {
            Direction.North: (0, 1),
            Direction.East: (1, 0),
            Direction.South: (0, -1),
            Direction.West: (-1, 0),
        }[self]

    def turn_direction(self, turn: str):
        if turn == "R":
            return self.turn_left()
        return self.turn_right()


def part1(data: str) -> int:
    pos = (0, 0)
    dir = Direction.North

    for seq in map(lambda s: s.strip(), data.split(",")):
        dir = dir.turn_direction(seq[0])
        blocks = int(seq[1:])
        x, y = dir.move_position()
        pos = (pos[0] + x * blocks, pos[1] + y * blocks)

    return abs(pos[0]) + abs(pos[1])


def part2(data: str) -> int:
    pos = (0, 0)
    dir = Direction.North
    visited = set()
    visited.add(pos)

    for seq in map(lambda s: s.strip(), data.split(",")):
        dir = dir.turn_direction(seq[0])
        blocks = int(seq[1:])
        x, y = dir.move_position()

        for _ in range(blocks):
            pos = (pos[0] + x, pos[1] + y)
            if pos in visited:
                return abs(pos[0]) + abs(pos[1])
            visited.add(pos)

    return abs(pos[0]) + abs(pos[1])


if __name__ == "__main__":
    with open("../input/01.txt") as f:
        data = f.read().strip()

    if sys.argv[1] == "part2":
        print(part1(data))
    else:
        print(part2(data))
