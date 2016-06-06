package flights

import (
  "errors"
  "fmt"
  "myUtil"
)

func FewestHops(originAirport string, destinationAirport string, connectedAirports [][]string, exclusions []string) ([]string, error) {
  aList := myUtil.EListToAList(connectedAirports)
  var queueToExplore = []string{ originAirport }
  var shortestPath []string // This is what we shall return
  var relationships = make(map[string]string) // map[to]from

  if originAirport == destinationAirport {
    return shortestPath, errors.New("You're already there!")
  }

  // Is the origin or destination in the exclusions list?
  for _, exc := range exclusions {
    if originAirport == exc || destinationAirport == exc {
      return shortestPath, errors.New("Airport is excluded")
    }
  }

  // Non-existent Aiport(s) -- RYAN can we use an OR?
  if _, ok := aList[originAirport]; !ok {
    return shortestPath, errors.New("Airport does not exist")
  }
  if _, ok := aList[destinationAirport]; !ok {
    return shortestPath, errors.New("Airport does not exist")
  }

  for len(queueToExplore) > 0 {
    node, _ := myUtil.Shift(&queueToExplore)
    for _, neighbor := range aList[node] {
      // Ignore it if we've already done it
      if _, ok := relationships[neighbor]; !ok {
        if !myUtil.InArray(neighbor, exclusions) {
          if neighbor == destinationAirport {
            relationships[neighbor] = node
            shortestPath = BuildPathFromRelationships(originAirport, destinationAirport, relationships)
            return shortestPath, nil
          } else {
            relationships[neighbor] = node
            queueToExplore = append(queueToExplore, neighbor)
            fmt.Println(relationships)
          }
        }
      }
    }
  }
  return shortestPath, errors.New("You can't get there from here")
}

func BuildPathFromRelationships(originAirport string, destinationAirport string, relationships map[string]string) []string {
  fmt.Println(relationships)
  var shortestPath []string
  currentNode := destinationAirport
  for {
    for k, v := range relationships {
      if currentNode == originAirport {
        shortestPath = append([]string{currentNode}, shortestPath...)
        return shortestPath
      } else if k == currentNode {
        shortestPath = append([]string{currentNode}, shortestPath...)
        currentNode = v
      }
    }
  }
  return shortestPath
}
