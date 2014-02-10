package main

import "fmt"

func main() {
  fmt.Println( waysToMakeChange(1100, []int{200, 100, 50, 20, 10, 5, 2, 1}) )
}

func waysToMakeChange(targetAmount int, coins []int) int{
  var recur func (total_so_far, coin_index int) int

  recur = func (total_so_far, coin_index int) int {
    val := coins[coin_index]
    if total_so_far == targetAmount || val == 1 {
      return 1
    }
    will_fit := (targetAmount - total_so_far) / val
    sum := 0
    for n := 0; n < will_fit+1; n++ {
      sum += recur(total_so_far+val*n, coin_index+1)
    }
    return sum
  }

  return recur(0,0)
}