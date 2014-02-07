package main

import(
  "fmt"
  "runtime"
)

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())
  coins := [...]int{200,100,50,20,10,5,2,1}
  ch := make(chan int)
  go waysToMakeChange(500, coins[:], ch)
  fmt.Println(<- ch)
}

func waysToMakeChange(amount int, coins []int, out chan int){
  if (amount == 0){
    out <- 1
    return
  }
  if (len(coins) == 1){
    if (amount % coins[0] == 0){
      out <- 1
      return
    } else {
      out <- 0
    }
    return
  }

  max := 1 + amount/coins[0]
  ch := make(chan int)
  sum := 0
  for i := 0; i<max; i++{
    go waysToMakeChange(amount - i*coins[0], coins[1:], ch)
  }

  for i := 0; i<max; i++{
    sum += <- ch
  }

  out <- sum
}