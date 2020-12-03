file = File.open('./lovisa/lib/2.txt')
input = []

file.each_line do |line|
  limits = line.split(' ').first.split('-')
  input << {
    character: line.split(' ')[1][0],
    password: line.split(':').last.strip,
    min: limits.first.to_i,
    max: limits.last.to_i
  }
end

valid = input.reduce(0) do |valid, entry|
  exists_on_first = entry[:password][entry[:min] - 1] == entry[:character]
  exists_on_second = entry[:password][entry[:max] - 1] == entry[:character]

  if (exists_on_first ^ exists_on_second)
    valid = valid + 1
  end

  valid
end

puts valid
