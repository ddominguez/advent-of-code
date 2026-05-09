import sys
from typing import NamedTuple


class Sequences(NamedTuple):
    inner: list[str]
    outer: list[str]


def generate_sequences(line: str) -> Sequences:
    items = line.split("[")
    sequences = Sequences(inner=[], outer=[items[0]])

    for item in items[1:]:
        inner, outer = item.split("]")
        sequences.inner.append(inner)
        sequences.outer.append(outer)

    return sequences


def part1(data: str) -> int:
    res = 0
    for line in data.splitlines():
        res += supports_tls(line)
    return res


def supports_tls(line: str) -> bool:
    sequences = generate_sequences(line)

    if any(has_abba(s) for s in sequences.inner):
        return False

    return any(has_abba(s) for s in sequences.outer)


def has_abba(s: str) -> bool:
    for i in range(len(s) - 3):
        if s[i + 1] == s[i + 2] and s[i] == s[i + 3] and s[i] != s[i + 1]:
            return True
    return False


def part2(data: str) -> int:
    res = 0
    for line in data.splitlines():
        res += supports_ssl(line)
    return res


def supports_ssl(line: str) -> bool:
    sequences = generate_sequences(line)
    abas = find_3_character_sequences(sequences.outer)
    babs = find_3_character_sequences(sequences.inner)
    for aba in abas:
        for bab in babs:
            if aba[0] == bab[1] and aba[1] == bab[0]:
                return True
    return False


def find_3_character_sequences(seqs: list[str]) -> set[str]:
    res = set()
    for s in seqs:
        for i in range(len(s) - 2):
            if s[i] == s[i + 2] and s[i] != s[i + 1]:
                res.add("".join(s[i : i + 3]))
    return res


if __name__ == "__main__":
    with open("../input/07.txt") as f:
        data = f.read().strip()

    if len(sys.argv) > 1 and sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
