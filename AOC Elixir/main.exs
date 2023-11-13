x = 1
addOne = fn a -> a + x end

x = 2
addTwo = fn a -> a + x end

IO.puts(addOne.(1))
IO.puts(addTwo.(1))
