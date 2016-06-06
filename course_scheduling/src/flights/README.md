package flights

import (
  "errors"
)

func fewestHops(connectedAirports [][]string) ([]string, error) {

}
1) do this unweighted
2) breadth first search for weighted graphs :-)

var delay = map[string]int

map[[]string][]string

{
  { lax, sfo, 4 }
}


lax, foo
lax, ord
sfo, iah
foo, iah
ord, iah
iah, lga
iah, mia
ord, mia

convert to alist

{
  lax: { { foo, { duration: 4, delay: 2 } },
         { ord, 1 },
       },
   sfo: {}
}
f(LAX, MIA) -> [LAX, ORD, MIA]

f(source, dest) -> [PATH]
f(source, dest, ignore) -> [OTHER PATH]

f(source, dest, graph)
  queueToExplore []string
  relationships map[string]string // map[to]from
  while len(queueToExplore) > 0
    node = queueToExplore.shift
    foreach neighbor & not in relationships
      is the neighbor the dest?
        buildPathFromRelationships()
      relationships.append[neighbor: node]
      queueToExplore.append(neighbor)
  return err if you can't get there from here

f(source, dest, ignore)
f(source, dest,

// can probably remove source since it's the only one with ""
// can probably remove dest as the last entry
buildPathFromRelationShips(dest, source, relationships)
  until the dest is the source
    if dest, then look up the previous dest
      add to my return array

