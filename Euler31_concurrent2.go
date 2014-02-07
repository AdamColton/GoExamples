package main

import(
  "fmt"
  "runtime"
)

type Accumulator struct{
  sum int
  ch chan int
}

func (self *Accumulator) loop(listener chan bool) {
  for ;; {
    self.sum += <- self.ch
    if (self.sum == 0){
      listener <- true
      return
    }
  }
}

func main(){
  runtime.GOMAXPROCS(runtime.NumCPU())
  coins := [...]int{200,100,50,20,10,5,2,1}
  accumulator := Accumulator{0, make(chan int, 2000)}
  runCounter := Accumulator{0, make(chan int, 2000)}
  listener := make(chan bool)
  go accumulator.loop(listener)
  go runCounter.loop(listener)
  go waysToMakeChange(500, coins[:], &accumulator, &runCounter)
  <-listener
  fmt.Println(accumulator.sum)
}

func waysToMakeChange(amount int, coins []int, accumulator, runCounter *Accumulator){
  if (amount == 0){
    accumulator.ch <- 1
    runCounter.ch <- -1
    return
  }
  if (len(coins) == 1){
    if (amount % coins[0] == 0){
      accumulator.ch <- 1
    }
    runCounter.ch <- -1
    return
  }

  max := 1 + amount/coins[0]
  runCounter.ch <- max
  for i := 0; i<max; i++{
    go waysToMakeChange(amount - i*coins[0], coins[1:], accumulator, runCounter)
  }
  runCounter.ch <- -1
}