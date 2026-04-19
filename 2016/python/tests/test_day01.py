import unittest

from collections import namedtuple

from days.day01 import part1, part2


class TestDay01(unittest.TestCase):
    def test_part1(self):
        TestCase = namedtuple("TestCase", ["input", "expected"])
        cases = [
            TestCase("R2, L3", 5),
            TestCase("R2, R2, R2", 2),
            TestCase("R5, L5, R5, R3", 12),
        ]

        for tc in cases:
            self.assertEqual(part1(tc.input), tc.expected)

    def test_part2(self):
        input = "R8, R4, R4, R8"
        expected = 4
        self.assertEqual(part2(input), expected)
