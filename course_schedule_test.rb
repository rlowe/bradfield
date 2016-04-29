#!/usr/bin/ruby

require_relative 'course_schedule'

examples = [
  {
    :given => [[1,2]],
    :expected => [2, 1]
  },
  {
    :given => [[1,2],[2,3]],
    :expected => [3, 2, 1]
  },
  {
    :given => [[1,2],[2,1]],
    :expected => "Cycle Detected: {1=>[2], 2=>[1]}"
  }
]

for example in examples
  begin
    actual = course_order example[:given]
    if actual != example[:expected]
      puts "[FAIL] EXPECTED: "
      p example[:expected]
      puts ", GOT: "
      p actual
      exit 1
    end
  rescue Exception => e
    if e.message != example[:expected]
      puts "[FAIL] EXPECTED: "
      p example[:expected]
      puts ", GOT: "
      puts e.message
    end
  end
end

puts "PASSED ALL THE THINGS!"
