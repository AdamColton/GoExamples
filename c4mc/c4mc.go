package c4mc

import (
  "../c4"
  "math/rand"
  "runtime"
)

/*
Carlo the Monkey
    __
w c(..)o    (
\__(-)    __)
    /\   (
  /(_)___)
 w /|
  | \
  m m
*/

func randomMove(game *c4.Game) int {
  moves := game.Moves()
  l := len(moves)
  return moves[ rand.Intn(l) ]
}

func monkeyCarlo(game *c4.Game){
  game.Move( randomMove(game) )
}

func fullMonkeyCarloGame(game *c4.Game) int{
  positions := 0
  for ; game.Turn() != 0 ;{
    monkeyCarlo(game)
    positions++
  }
  return positions
}

func moveValue(startingGame, endingPosition *c4.Game) int {
  if (endingPosition.Winner() == startingGame.Turn()) {
    return 1
  }
  return 0
}

type MonkeyCarlo struct{}
func (self *MonkeyCarlo) Move(game *c4.Game) {
  monkeyCarlo(game)
}

type MonteCarlo struct{
  Simulations int
}
func (self *MonteCarlo) Move(game *c4.Game) {
  moveRecord := [7][2]int{}
  for i := 0; i<self.Simulations; i++{
    move := randomMove(game)
    sim := game.Copy()
    moveRecord[move][0]++
    sim.Move(move)
    fullMonkeyCarloGame(&sim)
    moveRecord[move][1] += moveValue(game, &sim)
  }
  game.Move( bestMove(&moveRecord) )
}

func bestMove(moves *[7][2]int) int {
  bestMove := 0
  bestScore := 0.0
  for i := 0; i<7; i++ {
    if (moves[i][0] > 0){
      score := float64(moves[i][1]) / float64(moves[i][0])
      if (score > bestScore){
        bestScore = score
        bestMove = i
      }
    }
  }
  return bestMove
}

type MonteCarlo_P struct{
  positions int
  threads int
}
func NewMonteCarlo_P(simulations int) MonteCarlo_P {
  threads := runtime.NumCPU()
  runtime.GOMAXPROCS(threads)
  return MonteCarlo_P{simulations, threads}
}
type moveData struct{
  col int
  val int
  positions int
}
func (self *MonteCarlo_P) Move(game *c4.Game) {
  moveRecord := [7][2]int{}
  ch := make(chan moveData)
  stop := make(chan bool, self.threads)
  for i:=0; i<self.threads; i++ {
    go gameThread(ch, stop, game)
  }
  for i:=0 ; i<self.positions;{
    move := <-ch
    moveRecord[move.col][0]++
    moveRecord[move.col][1] += move.val
    i += move.positions
  }
  for i:=0; i<self.threads; i++ {
    stop <- true
  }
  for ;; {
    shouldBreak := false
    select{
    case move := <-ch:
      moveRecord[move.col][0]++
      moveRecord[move.col][1] += move.val
    default:
      shouldBreak = true
    }
    if (shouldBreak){
      break
    }
  }
  game.Move( bestMove(&moveRecord) )
}

func gameThread(out chan moveData, stop chan bool, game *c4.Game){
  for ;; {
    select{
    case <-stop:
      return
    default:
      move := randomMove(game)
      sim := game.Copy()
      sim.Move(move)
      positions := fullMonkeyCarloGame(&sim)
      out <- moveData{move, moveValue(game, &sim), positions}
    } 
  }
  
}