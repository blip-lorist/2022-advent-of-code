package main
import (
  "fmt"
  "io/ioutil"
  "strings"
  "github.com/ghetzel/go-stockutil/sliceutil"
)

var lowerCasePri = map[string] int {
  "a": 1,
  "b": 2,
  "c": 3,
  "d": 4,
  "e": 5,
  "f": 6,
  "g": 7,
  "h": 8,
  "i": 9,
  "j": 10,
  "k": 11,
  "l": 12,
  "m": 13,
  "n": 14,
  "o": 15,
  "p": 16,
  "q": 17,
  "r": 18,
  "s": 19,
  "t": 20,
  "u": 21,
  "v": 22,
  "w": 23,
  "x": 24,
  "y": 25,
  "z": 26,
}

var upperCasePri = map[string] int {
  "A": 27,
  "B": 28,
  "C": 29,
  "D": 30,
  "E": 31,
  "F": 32,
  "G": 33,
  "H": 34,
  "I": 35,
  "J": 36,
  "K": 37,
  "L": 38,
  "M": 39,
  "N": 40,
  "O": 41,
  "P": 42,
  "Q": 43,
  "R": 44,
  "S": 45,
  "T": 46,
  "U": 47,
  "V": 48,
  "W": 49,
  "X": 50,
  "Y": 51,
  "Z": 52,
}


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
    if (i % 3) == 0 {
      firstSet := [2]string{ rucksacks[i], rucksacks[i+1] }
      secondSet := [2]string{ rucksacks[i+1], rucksacks[i+2] }
      firstDupe := dupeFinder(firstSet)
      secondDupe := dupeFinder(secondSet)

      if (firstDupe != nil) {
        thirdSet := [2]string{ strings.Join(firstDupe[:], ","),  strings.Join(secondDupe[:], ",")}
        thirdDupe := dupeFinder(thirdSet)
        if thirdDupe != nil {
          prioritySum += getPriority(thirdDupe[0])
        }
      }
    }
  }
  return prioritySum
}

func getPriority(letter string) int {
  isUpper := letter == strings.ToUpper(letter)

  if isUpper {
    return upperCasePri[letter]
  } else {
    return lowerCasePri[letter]
  }
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
  //partTwoSolution := PartTwo(puzzleInput)

  fmt.Printf("PartOne: %d", partOneSolution)
  fmt.Println("")
  //fmt.Printf("PartTwo: %d", partTwoSolution)
  //fmt.Println("")
}

