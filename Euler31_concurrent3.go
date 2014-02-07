package main

import(
  "fmt"
  "runtime"
)

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())
  coins := [...]int{200,100,50,20,10,5,2,1}
  ch := make(chan int)
  go waysToMakeChange_concurrent(1100, coins[:], ch, 2)
  fmt.Println(<- ch)
}

func waysToMakeChange_concurrent(amount int, coins []int, out chan int, depth int){
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

  sum := 0
  if (depth > 0){
    depth--
    max := 1 + amount/coins[0]
    ch := make(chan int)
    for i := 0; i<max; i++{
      go waysToMakeChange_concurrent(amount - i*coins[0], coins[1:], ch, depth)
    }

    for i := 0; i<max; i++{
      sum += <- ch
    }
  } else {
    max := 1 + amount/coins[0]
    for i := 0; i<max; i++{
      sum += waysToMakeChange(amount - i*coins[0], coins[1:])
    }
  }
  out <- sum
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