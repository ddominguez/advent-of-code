import unittest

from days.day05 import part1, part2


class TestDay05(unittest.TestCase):
    def test_part1(self):
        input = "abc"
        expected = "18f47a30"
        self.assertEqual(part1(input), expected)

    def test_part2(self):
        input = "abc"
        expected = "05ace8e3"
        self.assertEqual(part2(input), expected)
