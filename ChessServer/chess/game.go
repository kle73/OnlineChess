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

func isCheck(g Game) string {
  kings := []Piece{}
  for _, piece := range g.pieces {
    if piece.getType() == "king" {
      kings = append(kings, piece)
    }
  }
  for _, king := range kings {
    for _, piece := range g.pieces {
      if piece.getColor() != king.getColor(){
        for _, position := range piece.getPossibleSteps(g){
          if position == king.getPosition() {
            return king.getColor()
          }
        }
      }
    }
  }
  return "none"
}

func isCheckmate(g Game) string {
  return "none"
}


func SetPiece(p1 string, p2 string, g Game) (game Game, msg string){
  if p1 == p2 {
    game, msg = g, "INVALID, TRY AGAIN"
    return
  }
  var pos1 Position = g.positions[p1]
  var pos2 Position = g.positions[p2]
  for _, piece := range g.pieces {
    if piece.getPosition() == pos1 {

      if piece.getColor() != g.Color {
        game, msg = g, "INVALID, TRY AGAIN"
        return
      }

      //position in possiblesteps?
      for _, position := range piece.getPossibleSteps(g) {
        if pos2 == position {
          //check if there already is a peice on Pos2
          for _, otherPiece := range g.pieces {
            if otherPiece.getPosition() == position {
              //there already is a piece with same color
              if otherPiece.getColor() == g.Color {
                game, msg = g, "INVALID, TRY AGAIN"
                return
              } else {
                var nilPos Position = Position{0, 0}
                otherPiece.setPosition(nilPos)
                piece.setPosition(pos2)
                //check if it would be check
                if g.Color == isCheck(g) {
                  otherPiece.setPosition(pos2)
                  piece.setPosition(pos1)
                  game, msg = g, "INVALID, THAT WOULD BE CHECK"
                  return
                } else if isCheckmate(g) != "none"{
                  game, msg = g, "CHECKMATE"
                  return
                }else if isCheck(g) != "none"{
                  g.turn += 1
                  g.Color = getCurrentColor(g)
                  game, msg = g, "SET CHECK"
                  return
                }else{
                  g.turn += 1
                  g.Color = getCurrentColor(g)
                  game, msg = g, "SET"
                  return
                }
              }
            }
          }
          piece.setPosition(pos2)
          //check if it would be check
          if g.Color == isCheck(g) {
            piece.setPosition(pos1)
            game, msg = g, "INVALID, THAT WOULD BE CHECK"
            return
          } else if isCheckmate(g) != "none"{
            game, msg = g, "CHECKMATE"
            return
          }else if isCheck(g) != "none"{
            g.turn += 1
            g.Color = getCurrentColor(g)
            game, msg = g, "SET CHECK"
            return
          }else {
            g.turn += 1
            g.Color = getCurrentColor(g)
            game, msg = g, "SET"
            return
          }
        }
      }
    }
  }
  game, msg = g, "INVALID, TRY AGAIN"
  return
}
