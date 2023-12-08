import os
from collections import Counter
from functools import cmp_to_key

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/07.txt", encoding="utf-8") as f:
    lines = f.read().strip().split("\n")

isPart2 = True
if isPart2:
    cards = ("J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A")
else:
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
    mc = c.most_common()
    fmc = mc[0][1]
    jokers = list(filter(lambda tpl: tpl[0] == "J", mc))
    if isPart2 and cl > 1 and len(jokers) > 0:
        cl -= 1
        if mc[0][0] == "J":
            fmc = mc[1][1] + jokers[0][1]
        else:
            fmc = fmc + jokers[0][1]

    if cl == 1:
        k = "5k"
    elif cl == 2:
        k = "4k" if fmc == 4 else "fh"
    elif cl == 3:
        k = "3k" if fmc == 3 else "2p"
    elif cl == 4:
        k = "1p"
    else:
        k = "hc"
    hand_type_collection[k].append(h)

total_winnings = 0
hand_rank = 0
for k, coll in hand_type_collection.items():
    coll.sort(key=cmp_to_key(hand_sorter))
    for h in coll:
        hand_rank += 1
        total_winnings += hand_bids[h] * hand_rank

print(f"part {2 if isPart2 else 1}:", total_winnings)
