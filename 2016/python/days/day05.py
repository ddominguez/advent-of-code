import hashlib
import sys


def part1(data: str) -> str:
    passwd_chars = []
    passwd_len = 8
    search_count = 0
    index = 0
    hash_prefix = "00000"
    while search_count < passwd_len:
        hash = hashlib.md5(f"{data}{index}".encode("utf-8")).hexdigest()
        if hash.startswith(hash_prefix):
            passwd_chars.append(hash[5])
            search_count += 1
        index += 1

    return "".join(passwd_chars)


def part2(data: str) -> str:
    passwd_len = 8
    passwd_chars = [""] * passwd_len
    search_count = 0
    index = 0
    hash_prefix = "00000"
    while search_count < passwd_len:
        hash = hashlib.md5(f"{data}{index}".encode("utf-8")).hexdigest()
        if hash.startswith(hash_prefix):
            try:
                pos = int(hash[5])
            except ValueError:
                pass
            else:
                if pos < passwd_len and passwd_chars[pos] == "":
                    passwd_chars[pos] = hash[6]
                    search_count += 1
        index += 1

    return "".join(passwd_chars)


if __name__ == "__main__":
    with open("../input/05.txt") as f:
        data = f.read().strip()

    if len(sys.argv) > 1 and sys.argv[1] == "part2":
        print(part2(data))
    else:
        print(part1(data))
