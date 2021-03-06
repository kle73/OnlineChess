package chess


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

func isCheckmate(g Game, color string) bool {
  for _, piece := range g.pieces {
    if piece.getColor() == color{
      for _, position := range piece.getPossibleSteps(g) {
        var cachePosition Position = piece.getPosition()
        piece.setPosition(position)
        //check throwout
        for _, otherPiece := range g.pieces {
          if otherPiece.getColor() != color && otherPiece.getPosition() == position {
            var nilPos Position = Position{0, 0}
            otherPiece.setPosition(nilPos)
          } else if otherPiece.getColor() == color && otherPiece.getPosition() == position {
            piece.setPosition(cachePosition)
            continue
          }
        }

        if isCheck(g) == color {
          piece.setPosition(cachePosition)
          continue
        } else if isCheck(g) != color {
          piece.setPosition(cachePosition)
          return false
        }
      }
    }
  }
  return true
}


func SetPiece(p1 string, p2 string, g Game) (game Game, msg string, opt string){
  opt = "0"
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
                }else if isCheck(g) != "none"{
                  g.turn += 1
                  g.Color = getCurrentColor(g)
                  if isCheckmate(g, g.Color){
                    game, msg, opt = g, "CHECKMATE", "1"
                    return
                  }
                  game, msg, opt = g, "SET CHECK", "1"
                  return
                }else{
                  g.turn += 1
                  g.Color = getCurrentColor(g)
                  game, msg, opt = g, "SET", "1"
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
          }else if isCheck(g) != "none"{
            g.turn += 1
            g.Color = getCurrentColor(g)
            if isCheckmate(g, g.Color) {
              game, msg = g, "CHECKMATE"
              return
            }
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
