import sys
from collections import Counter
from functools import cmp_to_key


def get_total_winnings(
    hand_bids: dict[str, int],
    hand_type_collection: dict[str, list[str]],
    cards: tuple[str, ...],
):
    def hand_sorter(a: str, b: str) -> int:
        for i in range(len(a)):
            if cards.index(a[i]) > cards.index(b[i]):
                return 1
            elif cards.index(a[i]) < cards.index(b[i]):
                return -1
        return 0

    total_winnings = 0
    hand_rank = 0
    for collection in hand_type_collection.values():
        collection.sort(key=cmp_to_key(hand_sorter))
        for hand in collection:
            hand_rank += 1
            total_winnings += hand_bids[hand] * hand_rank

    return total_winnings


def hand_type_key(card_type_count: int, most_common_count: int) -> str:
    match card_type_count:
        case 1:
            return "5k"
        case 2 if most_common_count == 4:
            return "4k"
        case 2:
            return "fh"
        case 3 if most_common_count == 3:
            return "3k"
        case 3:
            return "2p"
        case 4:
            return "1p"
        case _:
            return "hc"


def part1(data: str):
    cards = ("2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A")
    hand_bids: dict[str, int] = {}
    hand_type_collection: dict[str, list[str]] = {
        "hc": [],
        "1p": [],
        "2p": [],
        "3k": [],
        "fh": [],
        "4k": [],
        "5k": [],
    }

    for line in data.strip().split("\n"):
        h, b = line.split()
        hand_bids[h] = int(b)
        c = Counter(h)
        cl = len(c)
        mc = c.most_common()
        mcc = mc[0][1]
        hand_type_collection[hand_type_key(cl, mcc)].append(h)

    return get_total_winnings(hand_bids, hand_type_collection, cards)


def part2(data: str):
    cards = ("J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A")
    hand_bids: dict[str, int] = {}
    hand_type_collection: dict[str, list[str]] = {
        "hc": [],
        "1p": [],
        "2p": [],
        "3k": [],
        "fh": [],
        "4k": [],
        "5k": [],
    }

    for line in data.strip().split("\n"):
        h, b = line.split()
        hand_bids[h] = int(b)
        c = Counter(h)
        cl = len(c)
        mc = c.most_common()
        mcc = mc[0][1]
        jokers = list(filter(lambda tpl: tpl[0] == "J", mc))
        if cl > 1 and len(jokers) > 0:
            cl -= 1
            if mc[0][0] == "J":
                mcc = mc[1][1] + jokers[0][1]
            else:
                mcc = mcc + jokers[0][1]
        hand_type_collection[hand_type_key(cl, mcc)].append(h)

    return get_total_winnings(hand_bids, hand_type_collection, cards)


if __name__ == "__main__":
    with open("../input/07.txt") as f:
        data = f.read().strip()

    if sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
