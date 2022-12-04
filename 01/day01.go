package main
import (
  "fmt"
  "strings"
  "sort"
  "io/ioutil"
  "strconv"
  gm "github.com/spatialcurrent/go-math/pkg/math"
)

func PartOne(inputList string) int {
  maxCalsFound := 0
  elfPackCals := 0

  elfPackStrs := strings.Split(inputList, "\n\n")
  for i := 0; i < len(elfPackStrs); i++ {
    elfPack := elfPackStrs[i]
    elfPackItems := strings.Split(elfPack, "\n")

    itemsAsInts := []int{}
    for j := 0; j < len(elfPackItems); j++ {
      strVar := elfPackItems[j]
      intVar, _ := strconv.Atoi(strVar)
      itemsAsInts = append(itemsAsInts, intVar)
    }

    elfPackCals = add(itemsAsInts)
    if elfPackCals > maxCalsFound {
      maxCalsFound = elfPackCals
    }
  }

  return maxCalsFound
}

func PartTwo(inputList string) int {
  allTotals := []int{}

  elfPackStrs := strings.Split(inputList, "\n\n")
  for i := 0; i < len(elfPackStrs); i++ {
    elfPack := elfPackStrs[i]
    elfPackItems := strings.Split(elfPack, "\n")

    itemsAsInts := []int{}
    for j := 0; j < len(elfPackItems); j++ {
      strVar := elfPackItems[j]
      intVar, _ := strconv.Atoi(strVar)
      itemsAsInts = append(itemsAsInts, intVar)
    }

    elfPackCals := add(itemsAsInts)
    allTotals = append(allTotals, elfPackCals)
  }

  sort.Ints(allTotals)
  topThree := allTotals[len(allTotals)-3:]
  topThreeTotal := add(topThree)

  return topThreeTotal
}

func add(ints []int) int {
  sum, _ := gm.Sum(ints)

  // Cast the returned interface into an int
  intSum := sum.(int)
  return intSum
}

func main(){
  puzzleInputByte, _ := ioutil.ReadFile("./input.txt")
  puzzleInput := string(puzzleInputByte)

  partOneSolution := PartOne(puzzleInput)
  partTwoSolution := PartTwo(puzzleInput)
  fmt.Printf("PartOne: %d", partOneSolution)
  fmt.Println("")
  fmt.Printf("PartTwo: %d", partTwoSolution)
  fmt.Println("")
}
