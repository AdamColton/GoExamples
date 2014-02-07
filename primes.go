package main

import (
  "fmt"
  "math"
  "os"
  "strconv"
)

func main() {
  i, _ := strconv.Atoi(os.Args[1])
  primes(i)
}

func primes(x int) {
  if (x < 2){
    fmt.Println("Primes start at +2, do it right")
    return
  }
  composites := make([]bool, x)
  var sqrt = int( math.Ceil( math.Sqrt( float64(x)) ) )
  for i := 2; i<sqrt ; i++ {
    if (!composites[i]){
      for j := i*i; j<x; j += i{
        composites[j] = true
      } 
    }
  }
  c := 0
  for i := 2; i<x; i++{
    if (!composites[i]){
      c++
    }
  }
  fmt.Println(c);
}