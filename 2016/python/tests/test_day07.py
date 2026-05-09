import unittest
from typing import TypedDict

from days.day07 import part1, part2, supports_tls


class TestDay07(unittest.TestCase):
    @unittest.skip("")
    def test_supports_tls(self):
        class TestCase(TypedDict):
            input: str
            expected: bool

        testcases = [
            TestCase(input="abba[mnop]qrst", expected=True),
            TestCase(input="abcd[bddb]xyyx", expected=False),
        ]

        for tc in testcases:
            self.assertEqual(supports_tls(tc["input"]), tc["expected"])

    def test_part1(self):
        input = """abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn
"""
        expected = 2
        self.assertEqual(part1(input), expected)

    def test_part2(self):
        input = """aba[bab]xyz
xyx[xyx]xyx
aaa[kek]eke
zazbz[bzb]cdb
"""
        expected = 3
        self.assertEqual(part2(input), expected)
