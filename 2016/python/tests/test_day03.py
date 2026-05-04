import unittest

from days.day03 import part1, part2


class TestDay03(unittest.TestCase):
    def test_part1(self):
        input = "5 10 25"
        expected = 0
        self.assertEqual(part1(input), expected)

    def test_part2(self):
        input = """101 301 501
102 302 502
103 303 503
201 401 601
202 402 602
203 403 603"""
        expected = 6
        self.assertEqual(part2(input), expected)
