package main
import (
  //"bufio"
  "fmt"
  "strings"
  "io/ioutil"
  "strconv"
  gm "github.com/spatialcurrent/go-math/pkg/math"
)

func PartOne(inputList string) int {
  maxCalsFound := 0

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
    if elfPackCals > maxCalsFound {
      maxCalsFound = elfPackCals
    }
  }

  return maxCalsFound
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

  solution := PartOne(puzzleInput)
  fmt.Println(solution)
}
