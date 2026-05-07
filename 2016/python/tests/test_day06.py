import unittest

from days.day06 import part1, part2


class TestDay06(unittest.TestCase):
    def test_part1(self):
        input = """eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar"""
        expected = "easter"
        self.assertEqual(part1(input), expected)

    def test_part2(self):
        input = """eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar"""
        expected = "advent"
        self.assertEqual(part2(input), expected)
