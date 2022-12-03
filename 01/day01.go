package day01
import (
  "bufio"
  "strconv"
  "strings"
  gm "github.com/spatialcurrent/go-math/pkg/math"
)

func PartOne(inputList string) int {
  elfPack := []int{}
  maxCalsFound := 0
  scanner := bufio.NewScanner(strings.NewReader(inputList))
  for scanner.Scan() {
    strVar := scanner.Text()
    if strVar == "" {
      elfPackCals := add(elfPack)
      if elfPackCals > maxCalsFound {
        maxCalsFound = elfPackCals
      }
      elfPack = []int{}
    }

    intVar, _ := strconv.Atoi(strVar)
    elfPack = append(elfPack, intVar)
  }
  return maxCalsFound
}

func add(ints []int) int {
  sum, _ := gm.Sum(ints)

  // Cast the returned interface into an int
  intSum := sum.(int)
  return intSum
}
