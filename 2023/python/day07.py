import os
from dataclasses import dataclass
from collections import Counter
from functools import cmp_to_key

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/07.txt", encoding="utf-8") as f:
    lines = f.read().strip().split("\n")

cards = ("2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A")


def hand_sorter(a, b):
    for i in range(len(a)):
        if cards.index(a[i]) > cards.index(b[i]):
            return 1
        elif cards.index(a[i]) < cards.index(b[i]):
            return -1


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

for line in lines:
    h, b = line.split()
    hand_bids[h] = int(b)
    c = Counter(h)
    cl = len(c)
    if cl == 1:
        hand_type_collection["5k"].append(h)
    elif cl == 2:
        mc = c.most_common(1)[0][1]
        hand_type_collection["4k" if mc == 4 else "fh"].append(h)
    elif cl == 3:
        mc = c.most_common(1)[0][1]
        hand_type_collection["3k" if mc == 3 else "2p"].append(h)
    elif cl == 4:
        hand_type_collection["1p"].append(h)
    else:
        hand_type_collection["hc"].append(h)

total_winnings = 0
hand_rank = 0
for k, coll in hand_type_collection.items():
    coll.sort(key=cmp_to_key(hand_sorter))
    for h in coll:
        hand_rank += 1
        total_winnings += hand_bids[h] * hand_rank

print("part 1:", total_winnings)
