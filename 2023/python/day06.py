import os

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/06.txt", encoding="utf-8") as f:
    data = f.read().strip()

isPart1 = False
lines = data.split("\n")
possible_wins_per_race: list[int] = []

if isPart1:
    times = [int(t) for t in lines[0].split(":")[1].split()]
    distances = [int(d) for d in lines[1].split(":")[1].split()]
else:
    times = [int("".join(lines[0].split(":")[1].split()))]
    distances = [int("".join(lines[1].split(":")[1].split()))]

for i, ms in enumerate(times):
    wtw = 0
    for th in range(1, ms + 1):
        rd = (ms - th) * th
        wtw += 1 if rd > distances[i] else 0
    possible_wins_per_race.append(wtw)

result = 1
for w in possible_wins_per_race:
    result *= w

print(f"part {1 if isPart1 else 2} result:", result)
