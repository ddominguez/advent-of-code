import sys


def part1(data: str) -> str:
    codes: list[str] = []
    keypad = [
        ["1", "2", "3"],
        ["4", "5", "6"],
        ["7", "8", "9"],
    ]
    ROWS = len(keypad)
    COLS = ROWS
    btn_pos = (1, 1)

    for line in data.splitlines():
        for char in line:
            new_pos = get_new_pos(btn_pos, char)
            if is_in_bounds(new_pos, ROWS, COLS):
                btn_pos = new_pos
        codes.append(keypad[btn_pos[0]][btn_pos[1]])

    return "".join(codes)


def part2(data: str) -> str:
    codes: list[str] = []
    keypad = [
        [" ", " ", "1", " ", " "],
        [" ", "2", "3", "4", " "],
        ["5", "6", "7", "8", "9"],
        [" ", "A", "B", "C", " "],
        [" ", " ", "D", " ", " "],
    ]
    ROWS = len(keypad)
    COLS = ROWS
    EDGE = " "
    btn_pos = (2, 0)

    for line in data.splitlines():
        for char in line:
            new_pos = get_new_pos(btn_pos, char)
            if (
                is_in_bounds(new_pos, ROWS, COLS)
                and keypad[new_pos[0]][new_pos[1]] != EDGE
            ):
                btn_pos = new_pos
        codes.append(keypad[btn_pos[0]][btn_pos[1]])

    return "".join(codes)


def get_new_pos(pos: tuple[int, int], move_char: str) -> tuple[int, int]:
    match move_char:
        case "U":
            return (pos[0] - 1, pos[1])
        case "R":
            return (pos[0], pos[1] + 1)
        case "D":
            return (pos[0] + 1, pos[1])
        case "L":
            return (pos[0], pos[1] - 1)
        case _:
            raise Exception(f"Invalid move character: {move_char}")


def is_in_bounds(pos: tuple[int, int], rows: int, cols: int) -> bool:
    return pos[0] >= 0 and pos[0] < rows and pos[1] >= 0 and pos[1] < cols


if __name__ == "__main__":
    with open("../input/02.txt") as f:
        data = f.read().strip()

    if sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
