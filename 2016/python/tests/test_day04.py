import unittest

from days.day04 import part1, part2


class TestDay03(unittest.TestCase):
    def test_part1(self):
        input = """aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]"""
        expected = 1514
        self.assertEqual(part1(input), expected)

    def test_part2(self):
        input = "qzmt-zixmtkozy-ivhz-343[ihxyz]"
        expected = "343"
        self.assertEqual(part2(input, "very encrypted name"), expected)
