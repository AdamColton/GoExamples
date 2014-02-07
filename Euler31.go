package main

import(
  "fmt"
)

func main(){
  coins := [...]int{200,100,50,20,10,5,2,1}
  fmt.Println(waysToMakeChange(500, coins[:]))
}

func waysToMakeChange(amount int, coins []int)(int){
  if (amount == 0){
    return 1
  }
  if (len(coins) == 1){
    if (amount % coins[0] == 0){
      return 1
    }
    return 0
  }

  sum := 0
  max := 1 + amount/coins[0]
  for i := 0; i<max; i++{
    sum += waysToMakeChange(amount - i*coins[0], coins[1:])
  } 

  return sum
}