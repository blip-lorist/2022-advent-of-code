package day01
import (
  gm "github.com/spatialcurrent/go-math/pkg/math"
)

func PartOne(inputList string) int {
  return 0
}

func add(ints []int) int {
  sum, _ := gm.Sum(ints)

  // Cast the returned interface into an int
  intSum := sum.(int)
  return intSum
}
