package main

import (
  "fmt"
  "math"
)

func main() {
  primes(10000)
}

func primes(x int) {
  if (x < 2){
    fmt.Println("Primes start at +2, do it right")
    return
  }
  primes := make([]bool, x)
  var sqrt = int( math.Ceil( math.Sqrt( float64(x)) ) )
  for i := 2; i<sqrt ; i++ {
    if (!primes[i]){
      for j := i*i; j<x; j += i{
        primes[j] = true
      } 
    }
  }
  for i := 2; i<x; i++{
    if (!primes[i]){
      fmt.Println(i);
    }
  } 
}