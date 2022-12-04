package main
import (
  "fmt"
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
  return 0
}

func getPriority(letter string) int {
  isUpper := letter == strings.ToUpper(letter)

  if isUpper {
    return upperCasePri[letter]
  } else {
    return lowerCasePri[letter]
  }
}

func dupeFinder(compartments_pair [2]string) string {

  firstSlice := strings.Split(compartments_pair[0], "")
  secondSlice := strings.Split(compartments_pair[1], "")

  lettersIntersection := sliceutil.IntersectStrings(firstSlice, secondSlice)

  if len(lettersIntersection) > 0 {
    return lettersIntersection[0]
  } else {
    return ""
  }
}

