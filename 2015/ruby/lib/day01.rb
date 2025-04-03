module Day01
  DIRECTIONS = {'('=> 1, ')'=> -1}

  def self.part1(input)
    floor = 0
    input.chars.each do |char|
      floor += DIRECTIONS[char]
    end
    floor
  end

  def self.part2(input)
    floor = 0
    position = 0
    input.chars.each do |char|
      floor += DIRECTIONS[char]
      position += 1
      break if floor == -1
    end
    position
  end
end

if __FILE__ == $0
  input = File.read('../input/01.txt')
  if ARGV[0] == 'part2'
    puts Day01.part2(input)
  else
    puts Day01.part1(input)
  end
end
