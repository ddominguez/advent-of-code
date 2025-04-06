import os

root_path = os.popen("git rev-parse --show-toplevel").read().strip()
with open(f"{root_path}/2023/input/05.txt", encoding="utf-8") as f:
    data = f.read().strip()


seeds: list[int] = []
for i, item in enumerate(data.split("\n\n")):
    if i == 0:
        seeds = list(map(int, item.split(": ")[1].split(" ")))
        continue
    mapName, mapItem = item.split(" map:\n")
    ranges = [list(map(int, r.split(" "))) for r in mapItem.split("\n")]
    for i, sv in enumerate(seeds):
        for r in ranges:
            dsr, srs, rl = r
            if srs <= sv <= srs + rl - 1:
                seeds[i] = sv - srs + dsr
                break

print(f"part 1: {min(seeds)}")
