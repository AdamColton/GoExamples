package main

import(
  "./c4"
  "./c4mc"
  "fmt"
  "math/rand"
  "time"
)


func main() {
  rand.Seed( time.Now().UTC().UnixNano())
  game := c4.Game{}
  game.Init()
  move := 0
  ai := c4mc.NewMonteCarlo_P(1000000)
  for ; game.Turn() != 0 ; {
    if (game.Turn() == 1){
      ai.Move(&game)
    } else { 
      fmt.Println(game)
      fmt.Print("Move (0-6): ")
      fmt.Scanf("%d", &move)
      game.Move(move)
    }
  }
  fmt.Println(game)
}