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

func checker(amount int, coins []int)(int, bool){
  if (amount == 0){
    return 1, true
  }
  if (len(coins) == 1){
    if (amount % coins[0] == 0){
      return 1, true
    }
    return 0, true
  }
  return 0, false
}

func waysToMakeChange_concurrent(amount int, coins []int, out chan int, depth int){
  ways, shouldReturn := checker(amount, coins)
  if (shouldReturn){
    out <- ways
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
    sum = waysToMakeChange(amount, coins)
  }
  out <- sum
}

func waysToMakeChange(amount int, coins []int)(int){
  ways, shouldReturn := checker(amount, coins)
  if (shouldReturn){
    return ways
  }

  sum := 0
  max := 1 + amount/coins[0]
  for i := 0; i<max; i++{
    sum += waysToMakeChange(amount - i*coins[0], coins[1:])
  }

  return sum
}