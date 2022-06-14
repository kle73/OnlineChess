package chess


/*
checkmate Private
check Private
*/


type Game struct {
  positions  map[string]Position
  pieces     [32]Piece
  turn       int
  Color      string
}


func getCurrentColor(g Game) string {
  if g.turn % 2 == 0 {
    return "white"
  } else {
    return "black"
  }
}

func isCheck() string {
  return " "
}

func isCheckmate() string {
  return " "
}

func SetPiece(p1 string, p2 string, g Game) (game Game, msg string){
  var pos1 Position = g.positions[p1]
  var pos2 Position = g.positions[p2]
  for _, piece := range g.pieces {
    if piece.getPosition() == pos1 {

      if piece.getColor() != g.Color {
        game, msg = g, "INVALID, TRY AGAIN"
        return
      }

      for _, position := range piece.getPossibleSteps(g) {
        if pos2 == position {
          piece.setPosition(pos2)
          g.turn += 1
          g.Color = getCurrentColor(g)
          game, msg = g, "SET"
          return
        }
      }
    }
  }
  game, msg = g, "INVALID, TRY AGAIN"
  return
}
