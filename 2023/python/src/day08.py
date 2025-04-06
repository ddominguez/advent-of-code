import os

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/08.txt", encoding="utf-8") as f:
    lines = f.read().strip().split("\n")

instructions = lines[0]
network_map: dict[str, list[str]] = {}

for i in range(2, len(lines)):
    node, lrnodes = lines[i].split(" = ")
    network_map[node] = lrnodes.lstrip("(").rstrip(")").split(", ")

ins_map = {"L": 0, "R": 1}
curr_node = "AAA"
last_node = "ZZZ"
ins_i = 0
steps = 0

while True:
    steps += 1
    curr_node = network_map[curr_node][ins_map[instructions[ins_i]]]
    ins_i += 1
    if curr_node == last_node:
        break
    if ins_i == len(instructions):
        ins_i = 0

print("part 1 result:", steps)
