import sys
from functools import reduce


def part1(data: str):
    possible_wins_per_race: list[int] = []
    lines = data.strip().split("\n")
    times = [int(t) for t in lines[0].split(":")[1].split()]
    distances = [int(d) for d in lines[1].split(":")[1].split()]

    for i, ms in enumerate(times):
        ways_to_win = 0
        for th in range(1, ms + 1):
            rd = (ms - th) * th
            ways_to_win += 1 if rd > distances[i] else 0
        possible_wins_per_race.append(ways_to_win)

    return reduce(lambda x, y: x * y, possible_wins_per_race)


def part2(data: str):
    lines = data.strip().split("\n")
    time = int("".join(lines[0].split(":")[1].split()))
    distance = int("".join(lines[1].split(":")[1].split()))

    ways_to_win = 0
    for th in range(1, time + 1):
        rd = (time - th) * th
        ways_to_win += 1 if rd > distance else 0

    return ways_to_win


if __name__ == "__main__":
    with open("../input/06.txt") as f:
        data = f.read().strip()

    if sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
