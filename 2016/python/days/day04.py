import sys
from collections import Counter

SECTOR_ID_SLICE = slice(-10, -7)
CHECKSUM_SLICE = slice(-6, -1)
ROOM_NAME_END = -11


def part1(data: str) -> int:
    result = 0
    for line in data.splitlines():
        sector_id = int(line[SECTOR_ID_SLICE])
        checksum = line[CHECKSUM_SLICE]
        letters_count: Counter[str] = Counter()

        for items in line[:ROOM_NAME_END]:
            if items == "-":
                continue
            letters_count.update(items)

        if is_real(letters_count, checksum):
            result += sector_id

    return result


def part2(data: str, search_term: str) -> str | None:
    for line in data.splitlines():
        sector_id = line[SECTOR_ID_SLICE]
        shift = int(sector_id) % 26
        decrypted = []
        for word in line[:ROOM_NAME_END].split("-"):
            for char in word:
                shift_char = chr((ord(char) - ord("a") + shift) % 26 + ord("a"))
                decrypted.append(shift_char)
            decrypted.append(" ")
        decrypted_str = "".join(decrypted)
        if search_term in decrypted_str:
            return sector_id
    return None


def is_real(counter: Counter[str], checksum: str) -> bool:
    sorted_counts = sorted(counter.items(), key=lambda item: (-item[1], item[0]))
    return "".join(c for c, _ in sorted_counts[:5]) == checksum


if __name__ == "__main__":
    with open("../input/04.txt") as f:
        data = f.read().strip()

    if len(sys.argv) > 1 and sys.argv[1] == "part2":
        print(part2(data, "northpole"))
    else:
        print(part1(data))
