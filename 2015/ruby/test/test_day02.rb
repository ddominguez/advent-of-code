require 'test/unit'
require_relative '../lib/day02'

class TestDay02 < Test::Unit::TestCase
  def test_part1
    input = "
2x3x4
1x1x10"
    assert_equal 101, Day02.part1(input)
  end

  def test_part2
    input = "
2x3x4
1x1x10"
    assert_equal 48, Day02.part2(input)
  end
end

