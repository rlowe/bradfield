#!/usr/bin/ruby

def split_sql_stmt(sql_string)
  ignore_next_char = false
  statements = []
  current_query = ""
  is_in_quote = false

  for c in sql_string.chars
   current_query += c

    if ignore_next_char
      ignore_next_char = false
      next
    end

    if c == "\\"
      ignore_next_char = true
    elsif c == "'"
      is_in_quote = !is_in_quote
    elsif c == ";" && !is_in_quote
      statements << current_query
      current_query = ""
    end
  end

  if current_query != ""
    statements << current_query
  end

  statements
end
