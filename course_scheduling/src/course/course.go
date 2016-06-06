package course

import (
  "errors"
  "fmt"
  "util"
)

func TopOrder(elist [][]int) ([]int, error) {
  // +1 to the number of iterations that have to be done
  alist := util.EListToAList(elist)

  // Keys from alist
  // var sortQueue = keys(alist)
  var sortQueue = []int{}
  for k, _ := range alist {
    sortQueue = append(sortQueue, k)
  }

  var topSort = []int{}
  for len(sortQueue) > 0 {
    // We will shift later
    node := sortQueue[0]

    // Last item in the queue
    // Saves one function call wheeee
    //if len(sortQueue) == 0 && !inIntArray(node, topSort) {
    //  topSort = append(topSort, node)
    //  break
    //}

    if !inIntArray(node, topSort) {
      dfw, err := DepthFirstWalk(alist, node, topSort)
      if err != nil {
        return topSort, err
      }
      topSort = append(topSort, dfw[len(dfw)-1])

      // Remove value from sortQueue
      // Delete(sortQueue,dfw[len(dfw)-1]) -- Remember SETS?
      for i, n := range sortQueue {
        if n == dfw[len(dfw)-1] {
          sortQueue = append(sortQueue[0:i], sortQueue[i+1:]...)
        }
      }

      // Re-order the sortQueue
      // ShiftToFront(sortQueue, dfw[len(dfw)-2]) -- WAT?!
      if len(dfw) > 1 {
        for i, n := range sortQueue {
          if n == dfw[len(dfw)-2] {
            sortQueue = append(sortQueue[0:i], sortQueue[i+1:]...)
            sortQueue = append([]int{n}, sortQueue...)
          }
        }
      }
    }
  }

  return topSort, nil
}

func DepthFirstWalk(alist map[int][]int, node int, ignorable []int) ([]int, error) {
  var toExplore = []int{ node }
  var path = []int{ node }
  var currentNode int

  for {
    // Nothing More To Explore
    if len(toExplore) == 0 {
      break
    }
    currentNode, toExplore = toExplore[0], toExplore[1:]
    neighbors := alist[currentNode]
    if len(neighbors) == 0 {
      path = append(path, currentNode)
      return path, nil
    } else {
      for i, _ := range neighbors {
        if inIntArray(neighbors[i], path) {
          return path, errors.New(fmt.Sprintf("Cycle detected for course %d", neighbors[i]))
        }
        if !inIntArray(neighbors[i], ignorable) && !inIntArray(neighbors[i], path) {
          path = append(path, neighbors[i])
          toExplore = append(toExplore, neighbors[i])
          break
        }
      }
    }
  }
  return path, nil
}

func inIntArray(needle int, haystack []int) bool {
  for _, i := range haystack {
    if needle == i {
      return true
    }
  }

  return false
}
