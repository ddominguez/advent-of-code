require 'test/unit'
require_relative '../lib/day01'

class TestDay01 < Test::Unit::TestCase
  def test_part1
    assert_equal 0, Day01.part1('(())')
    assert_equal 3, Day01.part1('(()(()(')
    assert_equal -3, Day01.part1(')())())')
  end

  def test_part2
    assert_equal 1, Day01.part2(')')
    assert_equal 5, Day01.part2('()())')
  end
end
