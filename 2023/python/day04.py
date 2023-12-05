import os

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/04.txt", encoding="utf-8") as f:
    data = f.read().strip()

pt1_result = 0
card_winners: dict[int, int] = {}
card_processed: dict[int, int] = {}


def process_card(card: int):
    if card_winners[card] < 1:
        return
    for c in range(card + 1, card + 1 + card_winners[card]):
        process_card(c)
        card_processed[c] += 1


for line in data.split("\n"):
    card, res = line.split(": ")
    card_num = int(card.split()[1])
    left_nums, right_nums = res.split(" | ")
    l_list = [int(x) for x in left_nums.split(" ") if x]
    r_list = [int(x) for x in right_nums.split(" ") if x]
    winners = set(l_list) & set(r_list)
    points = 0
    for i, v in enumerate(winners):
        points = points + 1 if i == 0 else points * 2
    pt1_result += points

    card_winners[card_num] = len(winners)
    card_processed[card_num] = card_processed.get(card_num, 0) + 1

print(f"part 1: {pt1_result}")


# part 2
for card_num, win_count in card_winners.items():
    if win_count < 1:
        continue
    process_card(card_num)

print(f"part 2: {sum(card_processed.values())}")
