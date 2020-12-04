file = File.open('./lovisa/lib/3.txt')
forest = file.read.split("\n").map { |line| line.split('') }

x = 0
shift = 3
tree = '#'

trees = forest.reduce(0) do |trees, line|
  position = line[x]
  trees += 1 if position == tree

  x += shift
  x %= forest[0].length

  trees
end

puts trees
