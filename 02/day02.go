package main
import (
  "fmt"
  "strings"
  "io/ioutil"
)

var moveToWinningMove = map[string]string {
  "A": "PAPER",
  "B": "SCISSORS",
  "C": "ROCK",
}

var moveToLosingMove = map[string]string {
  "A": "SCISSORS",
  "B": "ROCK",
  "C": "PAPER",
}

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
  rounds := strings.Split(inputList, "\n")
  for i := 0; i < len(rounds); i++ {
    round := rounds[i]
    if round == "" {
      continue
    }

    score := scoreRoundOutcomes(round)
    totalScore += score
  }

  return totalScore
}

func scoreRoundMoves(roundString string) int {
  var opponentMineMapping = map[string]string {
    "A": "X",
    "B": "Y",
    "C": "Z",
  }

  var playerMoveToPoints = map[string]int {
    "X": 1,
    "Y": 2,
    "Z": 3,
  }

  var playerMoveMapping = map[string]string {
    "ROCK": "X",
    "PAPER": "Y",
    "SCISSORS": "Z",
  }

  roundSlice := strings.Split(roundString, " ")
  opponentMove := roundSlice[0]
  myMove := roundSlice[1]

  myMovePoints := playerMoveToPoints[myMove]

  isDraw := opponentMineMapping[opponentMove] == myMove

  winningResponse := moveToWinningMove[opponentMove]
  winningResponseEncoded := playerMoveMapping[winningResponse]
  isWin := myMove == winningResponseEncoded

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


func scoreRoundOutcomes(roundString string) int {
  var opponentEncodingToMove = map[string]string {
    "A": "ROCK",
    "B": "PAPER",
    "C": "SCISSORS",
  }

  var playerMoveToPoints = map[string]int {
    "ROCK": 1,
    "PAPER": 2,
    "SCISSORS": 3,
  }

  var outcomeCodeToOutcome = map[string]string {
    "X": "LOSE",
    "Y": "DRAW",
    "Z": "WIN",
  }

  roundSlice := strings.Split(roundString, " ")
  opponentMove := roundSlice[0]
  desiredOutcome := roundSlice[1]

  wantDraw :=  outcomeCodeToOutcome[desiredOutcome] == "DRAW"
  wantWin :=  outcomeCodeToOutcome[desiredOutcome] == "WIN"

  roundOutcome := 0
  myMovePoints := 0
  if wantDraw { // draw
    roundOutcome = 3
    drawMove := opponentEncodingToMove[opponentMove]
    myMovePoints = playerMoveToPoints[drawMove]
  } else if wantWin {
    roundOutcome = 6
    winningMove := moveToWinningMove[opponentMove]
    myMovePoints = playerMoveToPoints[winningMove]
  } else {
    roundOutcome = 0
    losingMove := moveToLosingMove[opponentMove]
    myMovePoints = playerMoveToPoints[losingMove]
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
