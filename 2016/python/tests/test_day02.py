import unittest

from days.day02 import part1, part2


class TestDay02(unittest.TestCase):
    def test_part1(self):
        input = "ULL\nRRDDD\nLURDL\nUUUUD"
        expected = "1985"
        self.assertEqual(part1(input), expected)

    def test_part2(self):
        input = "ULL\nRRDDD\nLURDL\nUUUUD"
        expected = "5DB3"
        self.assertEqual(part2(input), expected)
