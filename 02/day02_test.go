package main

import "testing"

func TestPartOne(t *testing.T){
  testInput := `A Y
B X
C Z`

  expected := 15
  actual := PartOne(testInput)

  if expected != actual {
    t.Errorf("got %d, wanted %d", actual, expected)
  }
}

func TestPartTwo(t *testing.T){
  testInput := `A Y
B X
C Z`

  expected := 12
  actual := PartOne(testInput)

  if expected != actual {
    t.Errorf("got %d, wanted %d", actual, expected)
  }
}

func TestScoreRound(t *testing.T){
  testInputs := [3]string{"A Y", "B X", "C Z"}

  expectedOutputs := [3]int{8, 1, 6}

  for i := 0; i < len(testInputs); i++ {
    expected := expectedOutputs[i]
    actual := scoreRound(testInputs[i])

    if expected != actual {
      t.Errorf("got %d, wanted %d", actual, expected)
    }
  }
}
