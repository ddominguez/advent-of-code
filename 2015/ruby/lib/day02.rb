module Day02
  def self.part1(input)
    res = 0
    input.strip.split("\n").each do |line|
      l, w, h = line.split("x").map {|c| c.to_i}
      area = 2 * l * w + 2 * w * h + 2 * h * l
      smallest_side = [l * w, w * h, h * l].min
      res += area + smallest_side
    end
    res
  end

  def self.part2(input)
    res = 0
    input.strip.split("\n").each do |line|
      lwh = line.split("x").map {|c| c.to_i}
      volume = lwh.inject(1) {|acc, x| x*acc}
      smallest_perimeter = lwh.sort.take(2).inject(0) {|acc, x| x+x+acc}
      res += volume + smallest_perimeter
    end
    res
  end
end

if __FILE__ == $0
  input = File.read('../input/02.txt')
  if ARGV[0] == 'part2'
    puts Day02.part2(input)
  else
    puts Day02.part1(input)
  end
end

