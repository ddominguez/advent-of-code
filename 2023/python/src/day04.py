import sys


def part1(data: str):
    result = 0

    for line in data.strip().split("\n"):
        _, card_numbers = line.split(": ")
        left_nums, right_nums = card_numbers.split(" | ")
        l_list = [int(x) for x in left_nums.split(" ") if x]
        r_list = [int(x) for x in right_nums.split(" ") if x]
        winners = set(l_list) & set(r_list)
        points = 0

        for i, _ in enumerate(winners):
            points = points + 1 if i == 0 else points * 2

        result += points

    return result


def part2(data: str):
    card_number = 0
    cards_processed: dict[int, int] = {}

    for line in data.strip().split("\n"):
        card_number += 1
        _, card_numbers = line.split(": ")
        left_nums, right_nums = card_numbers.split(" | ")
        l_list = [int(x) for x in left_nums.split(" ") if x]
        r_list = [int(x) for x in right_nums.split(" ") if x]
        winners = set(l_list) & set(r_list)

        cards_processed[card_number] = cards_processed.get(card_number, 0) + 1
        for i, _ in enumerate(winners):
            key = card_number + i + 1
            cards_processed[key] = (
                cards_processed.get(key, 0) + cards_processed[card_number]
            )

    return sum(cards_processed.values())


if __name__ == "__main__":
    with open("../input/04.txt") as f:
        data = f.read().strip()

    if sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
