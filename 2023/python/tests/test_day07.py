import unittest

from src.day07 import part1, part2


class TestDay07(unittest.TestCase):
    def test_part1(self):
        data = """
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""
        self.assertEqual(part1(data), 6440)

    def test_part2(self):
        data = """
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""
        self.assertEqual(part2(data), 5905)
