package day01

import "testing"

// func TestPartOne(t *testing.T){
//   testInput := `
// 1000
// 2000
// 3000
//
// 4000
//
// 5000
// 6000
//
// 7000
// 8000
// 9000
//
// 10000`
//   expected := 24000
//   actual := PartOne(testInput)
//
//   if expected != actual {
//     t.Errorf("got %d, wanted %d", expected, actual)
//   }
// }


func TestAdd(t *testing.T){
  testInput := []int{1000, 2000, 3000}
  expected := 6000
  actual := add(testInput)

  if expected != actual {
    t.Errorf("got %d, wanted %d", expected, actual)
  }
}
