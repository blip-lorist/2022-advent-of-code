package main
import (
  "strings"
)

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

func PartOne(inputList string) int {
  totalScore := 0
  rounds := strings.Split(inputList, "\n")
  for i := 0; i < len(rounds); i++ {
    round := rounds[i]
    score := scoreRound(round)
    totalScore += score
  }

  return totalScore
}

func scoreRound(roundString string) int {
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

