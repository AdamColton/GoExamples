package main

import(
  "./c4"
  "fmt"
  
)

type mine []int

func main() {
  game := c4.Game{}
  game.Init()
  move := 0
  for ; game.Turn() != 0 ; {
    fmt.Println(game)
    fmt.Print("Move (0-6): ")
    fmt.Scanf("%d", &move)
    game.Move(move)
  }
  fmt.Println(game)
}