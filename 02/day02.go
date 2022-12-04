package main
import (
  "fmt"
  "strings"
  "io/ioutil"
)

func PartOne(inputList string) int {
  totalScore := 0
  rounds := strings.Split(inputList, "\n")
  for i := 0; i < len(rounds); i++ {
    round := rounds[i]
    if round == "" {
      continue
    }

    score := scoreRoundMoves(round)
    totalScore += score
  }

  return totalScore
}

func PartTwo(inputList string) int {
  totalScore := 0
  return totalScore
}

func scoreRoundMoves(roundString string) int {
  var opponentMineMapping = map[string]string {
    "A": "X",
    "B": "Y",
    "C": "Z",
  }

  var moveToWinningMove = map[string]string {
    "A": "Y",
    "B": "Z",
    "C": "X",
  }

  var playerMoveToPoints = map[string]int {
    "X": 1,
    "Y": 2,
    "Z": 3,
  }

  roundSlice := strings.Split(roundString, " ")
  opponentMove := roundSlice[0]
  myMove := roundSlice[1]

  myMovePoints := playerMoveToPoints[myMove]

  isDraw := opponentMineMapping[opponentMove] == myMove

  winningResponse := moveToWinningMove[opponentMove]
  isWin := myMove == winningResponse


  roundOutcome := 0
  if isDraw { // draw
    roundOutcome = 3
  } else if isWin {
    roundOutcome = 6
  } else {
    roundOutcome = 0
  }

  return roundOutcome + myMovePoints
}

func main() {
  puzzleInputByte, _ := ioutil.ReadFile("./input.txt")
  puzzleInput := string(puzzleInputByte)


  partOneSolution := PartOne(puzzleInput)
  fmt.Printf("PartOne: %d", partOneSolution)
  fmt.Println("")
  //fmt.Printf("PartTwo: %d", partTwoSolution)
  //fmt.Println("")
}
