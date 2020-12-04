file = File.open('./lovisa/lib/3.txt')
forest = file.read.split("\n").map { |line| line.split('') }

x, y = 0, 0
tree = '#'

shifts = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]

product = shifts.reduce([]) do |count, shift|
  trees = 0
  while y < forest.length
    position = forest[y][x]

    trees += 1 if position == tree

    y += shift[1]
    x += shift[0]
    x %= forest[0].length
  end

  x, y = 0, 0
  count.concat([trees])
end.reduce(:*)

puts product
