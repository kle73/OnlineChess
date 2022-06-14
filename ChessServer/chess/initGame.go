package chess

import "strconv"


type Position struct {
  Char int
  Num int
}

//Initialize the Positions
func initPositions() map[string]Position {
  var positions = map[string]Position{}
  for c := 65; c <= 72; c++ {
    for i := 1; i < 9; i++ {
      var key string = string(c) + strconv.Itoa(i)
      positions[key] = Position{c, i}
    }
  }
  return positions
}

//initialize the Pieces
func initPieces(positions map[string]Position) [32]Piece {
  var pieces = [32]Piece{}
  pieces[0] = &King{positions["E1"], "white"}
  pieces[1] = &King{positions["E8"], "black"}
  pieces[2] = &Queen{positions["D1"], "white"}
  pieces[3] = &Queen{positions["D8"], "black"}

  pieces[4] = &Bishop{positions["C1"], "white"}
  pieces[5] = &Bishop{positions["F1"], "white"}
  pieces[6] = &Bishop{positions["C8"], "black"}
  pieces[7] = &Bishop{positions["F8"], "black"}

  pieces[8] = &Knight{positions["B1"], "white"}
  pieces[9] = &Knight{positions["G1"], "white"}
  pieces[10] = &Knight{positions["B8"], "black"}
  pieces[11] = &Knight{positions["G8"], "black"}

  pieces[12] = &Rook{positions["A1"], "white"}
  pieces[13] = &Rook{positions["H1"], "white"}
  pieces[14] = &Rook{positions["A8"], "black"}
  pieces[15] = &Rook{positions["H8"], "black"}

  for c, i := 65, 16; c <= 72; c, i = c + 1, i + 1{
    for _, position := range positions{
      if position.Char == c && position.Num == 7 {
        pieces[i] = &Pawn{position, "black"}
      }
    }
  }
  for c, i := 65, 24; c <= 72; c, i = c+1, i+1{
    for _, position := range positions{
      if position.Char == c && position.Num == 2 {
        pieces[i] = &Pawn{position, "white"}
      }
    }
  }
  return pieces
}



//initialize a new Game
func NewGame() Game {
  var game Game = Game{}
  positions := initPositions()
  pieces := initPieces(positions)
  game.positions = positions
  game.pieces = pieces
  game.turn = 0
  game.Color = "white"
  return game
}
