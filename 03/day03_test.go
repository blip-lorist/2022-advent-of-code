package main

import "testing"

func TestPartOne(t *testing.T){
  testInput := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

  expected := 157
  actual := PartOne(testInput)

  if expected != actual {
    t.Errorf("got %d, wanted %d", actual, expected)
  }
}

func TestGetPriority(t *testing.T){
  testInputs := [6]string{"p", "L", "P", "v", "t", "s"}

  expectedOutputs := [6]int{16, 38, 42, 22, 20, 19}

  for i := 0; i < len(testInputs); i++ {
    expected := expectedOutputs[i]
    actual := getPriority(testInputs[i])

    if expected != actual {
      t.Errorf("got %d, wanted %d", actual, expected)
    }
  }
}
