package main

import "testing"

func TestPartOne(t *testing.T){
  testInput := `A Y
B X
C Z
`
  expected := 15
  actual := PartOne(testInput)

  if expected != actual {
    t.Errorf("got %d, wanted %d", actual, expected)
  }
}
