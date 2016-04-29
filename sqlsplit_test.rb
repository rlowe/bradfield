#!/bin/ruby

require_relative 'sqlsplit'

examples = {
  "a" => ["a"],
  "a;b" => ["a;", "b"],
  "a;" => ["a;"],
  "';';" => ["';';"],
  "';';a;" => ["';';", "a;"],
  "';\\';';a;" => ["';\\';';", "a;"],
  "INSERT INTO comms VALUES ('f;o'); INSERT INTO comms VALUES ('bar')" =>
    ["INSERT INTO comms VALUES ('f;o');", " INSERT INTO comms VALUES ('bar')"],
  "INSERT INTO comms VALUES ('fo');" =>
    ["INSERT INTO comms VALUES ('fo');"],
  "INSERT INTO comms VALUES ('\\'fo;o\\''); INSERT INTO comms VALUES ('bar')" =>
    ["INSERT INTO comms VALUES ('\\'fo;o\\'');", " INSERT INTO comms VALUES ('bar')"],
  "INSERT INTO comms VALUES ('foo'); INSERT INTO comms VALUES ('bar')" =>
    ["INSERT INTO comms VALUES ('foo');", " INSERT INTO comms VALUES ('bar')"]
}

for given_stmt, expected_stmt in examples
  actual = split_sql_stmt given_stmt
  if actual != expected_stmt
    puts "FAIL"
    puts "EXPECTED: ", expected_stmt
    puts "GOT: ", actual
  else
    puts "passed"
  end
end
