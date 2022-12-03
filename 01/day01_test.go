package day01

import "testing"

func TestHello(t *testing.T){
  testInput := `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`
  expected := 24000
  actual := DayOne(testInput)

  if expected != actual {
    t.Errorf("got %d, wanted %d", expected, actual)
  }
}
