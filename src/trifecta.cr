class Trifeca
  def initialize(ac : Int32, b : Int32)
    @ac = ac
    @b = b
    @table = [] of Tuple(Int32, Int32)
  end

  def is_whole(x : Int32, y : Int32)
    return x % y == 0;
  end

  def make_table
    (1..@ac).to_a.each do |item|
      possible = @ac / item

      if is_whole(@ac, item)
        @table << { item, possible.to_i }
      end
    end
  end

  def is_factor(x : Int32, y : Int32)
    (x + y == @b) && (x * y == @ac)
  end

  def find
    make_table

    matches = 0

    @table.each do |row|
      if is_factor(row[0], row[1])
        matches += 1
        puts  "#{row[0]} and #{row[1]} are the right factors"
      end
    end

    if matches == 0
      puts "No factors found"
    end

  end
end

begin
  ac = ARGV[0].to_i
  b = ARGV [1].to_i

  factor = Trifeca.new(ac, b)
  factor.find
rescue ex
  puts "No command line args given"
end

