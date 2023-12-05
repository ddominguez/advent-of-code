import os

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/04.txt", encoding="utf-8") as f:
    data = f.read().strip()

pt1_result = 0
for line in data.split("\n"):
    card, res = line.split(": ")
    left_nums, right_nums = res.split(" | ")
    l_list = [int(x) for x in left_nums.split(" ") if x]
    r_list = [int(x) for x in right_nums.split(" ") if x]
    winners = set(l_list) & set(r_list)
    points = 0
    for i, v in enumerate(winners):
        points = points + 1 if i == 0 else points * 2
    pt1_result += points

print(f"part 1: {pt1_result}")
