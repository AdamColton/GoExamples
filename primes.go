package main

import (
  "fmt"
  "math"
)

func main() {
  primes(100)
}

func primes(x int){
  if (x < 2){
    fmt.Println("Primes start at +2, do it right")
    return
  }
  x -= 2
  p := make([]bool, x)
  findPrimes(p)
  for i := 0; i<x; i++ {
    if (!p[i]){
      fmt.Println(i+2)
    }
  }
}

func findPrimes(primes []bool) {
  if (len(primes) == 1){
    return
  }
  var x = len(primes) + 1
  var sqrt = int( math.Ceil( math.Sqrt( float64(x)) ) )
  findPrimes(primes[0:sqrt-1])
  for i := 0; i<x-1 ; i++ {
    if (!primes[i]){
      for j := (i+2)*(i+2); j-1<x; j += (i+2){
        primes[j-2] = true
      } 
    }
  }
}