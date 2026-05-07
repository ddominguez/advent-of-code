import sys
from collections import Counter


def build_column_frequencies(lines: list[str]) -> dict[int, Counter[str]]:
    res: dict[int, Counter[str]] = {}
    cols = len(lines[0])
    for line in lines:
        for i in range(cols):
            if i not in res:
                res[i] = Counter()
            res[i].update(line[i])
    return res


def part1(data: str) -> str:
    result = []
    col_frequencies = build_column_frequencies(data.splitlines())
    for i in range(len(col_frequencies)):
        mc = col_frequencies[i].most_common()[0]
        result.append(mc[0])

    return "".join(result)


def part2(data: str) -> str:
    result = []
    col_frequencies = build_column_frequencies(data.splitlines())
    for i in range(len(col_frequencies)):
        lc = col_frequencies[i].most_common()[-1]
        result.append(lc[0])

    return "".join(result)


if __name__ == "__main__":
    with open("../input/06.txt") as f:
        data = f.read().strip()

    if len(sys.argv) > 1 and sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
