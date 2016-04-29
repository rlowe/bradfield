#!/usr/bin/ruby

def elist_to_alist(course_dep_elist)
  course_alist = {}
  for course_id, dep_course_id in course_dep_elist
    course_alist[course_id] ||= []
    course_alist[dep_course_id] ||= []
    course_alist[course_id] << dep_course_id
  end

  course_alist
end

def course_order(course_dep_elist)
  raise Exception.new("No Courses Defined") if course_dep_elist == []

  course_alist = elist_to_alist course_dep_elist

  ordered_courses = []
  previous_alist = {}

  while !course_alist.empty?
    previous_alist = course_alist.dup
    for course_id, dependent_courses in course_alist
      if dependent_courses == []
        ordered_courses << course_id
        course_alist.delete(course_id)
      end
    end

    for course_id, dependent_courses in course_alist
      for course in ordered_courses
        if dependent_courses.include?(course)
          course_alist[course_id].delete(course)
        end
      end
    end

    if course_alist == previous_alist
      raise Exception.new("Cycle Detected: #{course_alist.inspect}")
    end
  end

  return ordered_courses
end
