package main

import(
  "./c4"
  "./c4mc"
  "math/rand"
  "time"
)


func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  game := c4.Game{}
  game.Init()
  ai := c4mc.NewMonteCarlo(1000000)
  ai.Move(&game)
}