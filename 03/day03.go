package main
import (
  "fmt"
  "io/ioutil"
  "strings"
  "github.com/ghetzel/go-stockutil/sliceutil"
)

const priorityOrder = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PartOne(inputList string) int {
  prioritySum := 0
  rucksacks := strings.Split(inputList, "\n")
  for i := 0; i < len(rucksacks); i++ {
    rucksack := rucksacks[i]
    midIdx := len(rucksack) / 2
    halves := [2]string{ rucksack[0: midIdx], rucksack[midIdx:] }

    dupe := dupeFinder(halves)
    if dupe != nil {
      prioritySum += getPriority(dupe[0])
    }
  }
  return prioritySum
}

func PartTwo(inputList string) int {
  prioritySum := 0
  rucksacks := strings.Split(inputList, "\n")
  for i := 0; i < len(rucksacks); i++ {
    if (i%3 == 0) && (i+2 < len(rucksacks))  {
      firstSet := [2]string{ rucksacks[i], rucksacks[i+1] }
      secondSet := [2]string{ rucksacks[i+1], rucksacks[i+2] }
      firstDupe := dupeFinder(firstSet)
      secondDupe := dupeFinder(secondSet)

      if (firstDupe != nil) {
        thirdSet := [2]string{ strings.Join(firstDupe[:], ""),  strings.Join(secondDupe[:], "")}
        thirdDupe := dupeFinder(thirdSet)
        if thirdDupe != nil {
          fmt.Println(thirdDupe)
          prioritySum += getPriority(thirdDupe[0])
        }
      }
    }
  }
  return prioritySum
}

func getPriority(letter string) int {
  idx := strings.Index(priorityOrder, letter)
  return idx + 1
}

func dupeFinder(compartments_pair [2]string) []string {

  firstSlice := strings.Split(compartments_pair[0], "")
  secondSlice := strings.Split(compartments_pair[1], "")

  lettersIntersection := sliceutil.IntersectStrings(firstSlice, secondSlice)

  if len(lettersIntersection) > 0 {
    return lettersIntersection
  } else {
    return nil
  }
}

func main() {
  puzzleInputByte, err := ioutil.ReadFile("./input.txt")
  if err != nil {
    panic(err)
  }

  puzzleInput := string(puzzleInputByte)


  partOneSolution := PartOne(puzzleInput)
  partTwoSolution := PartTwo(puzzleInput)

  fmt.Printf("PartOne: %d", partOneSolution)
  fmt.Println("")
  fmt.Printf("PartTwo: %d", partTwoSolution)
  fmt.Println("")
}

