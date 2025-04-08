import unittest

from src.day06 import part1, part2


class TestDay06(unittest.TestCase):
    def test_part1(self):
        data = """
Time:      7  15   30
Distance:  9  40  200
"""
        self.assertEqual(part1(data), 288)

    def test_part2(self):
        data = """
Time:      7  15   30
Distance:  9  40  200
"""
        self.assertEqual(part2(data), 71503)
