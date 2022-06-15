package chess

import (
  "math"
)

type Piece interface {
  getPosition() Position
  setPosition(position Position)
  getColor() string
  getType() string
  getPossibleSteps(g Game) map[string]Position
}
type Bishop struct{
  position Position
  color string
}

func (b *Bishop) getPossibleSteps(g Game) map[string]Position {
 var steps = map[string]Position{}
  for key, position := range g.positions {
    if math.Abs(float64(position.Char - b.position.Char)) == math.Abs(float64(position.Num - b.position.Num)){
      steps[key] = position
    }
  }
	for _, piece := range g.pieces{
    if piece.getPosition() != b.getPosition(){
      for _, step := range steps {
        if step == piece.getPosition(){

				  var charDifference int = piece.getPosition().Char - b.getPosition().Char
				  var numDifference int = piece.getPosition().Num - b.getPosition().Num

					if charDifference <= -1 && numDifference >= 1 {
            for _, position := range g.positions{
				      for stepKey, step := range steps {
								if step == position {
								  if position.Char - b.getPosition().Char < charDifference && position.Num - b.getPosition().Num > numDifference {
								    delete(steps, stepKey)
									}
								}
							}
						}
					}	else if charDifference >= 1 && numDifference >= 1 {
            for _, position := range g.positions{
				      for stepKey, step := range steps {
								if step == position {
								  if position.Char - b.getPosition().Char > charDifference && position.Num - b.getPosition().Num > numDifference {
								    delete(steps, stepKey)
									}
								}
							}
						}
					}	else if charDifference <= -1 && numDifference <= -1 {
            for _, position := range g.positions{
				      for stepKey, step := range steps {
								if step == position {
								  if position.Char - b.getPosition().Char < charDifference && position.Num - b.getPosition().Num < numDifference {
								    delete(steps, stepKey)
									}
								}
							}
						}
          }	else if charDifference >= 1 && numDifference <= -1 {
            for _, position := range g.positions{
				      for stepKey, step := range steps {
								if step == position {
								  if position.Char - b.getPosition().Char > charDifference && position.Num - b.getPosition().Num < numDifference {
								    delete(steps, stepKey)
									}
								}
							}
						}
					}
				}
			}
		}
	}

  return steps
}
func (b *Bishop) getPosition() Position {
  return b.position
}
func (bp *Bishop) setPosition(position Position) {
  bp.position = position
}
func (b *Bishop) getColor() string {
  return b.color
}
func (b *Bishop) getType() string {
  return "bishop"
}



type Knight struct{
  position Position
  color string
}

func (k *Knight) getPossibleSteps(g Game) map[string]Position {
  var steps = map[string]Position{}
  for key, position := range g.positions{
    if position.Char == k.position.Char + 2{
      if position.Num == k.position.Num + 1{
        steps[key] = position
      } else if position.Num == k.position.Num - 1{
        steps[key] = position
      }
    } else if position.Char == k.position.Char + 1{
      if position.Num == k.position.Num + 2{
        steps[key] = position
      } else if position.Num == k.position.Num - 2{
        steps[key] = position
      }
    } else if position.Char == k.position.Char -1 {
      if position.Num == k.position.Num + 2{
        steps[key] = position
      } else if position.Num == k.position.Num - 2{
        steps[key] = position
      }
    } else if position.Char == k.position.Char - 2{
      if position.Num == k.position.Num + 1{
        steps[key] = position
      } else if position.Num == k.position.Num - 1{
        steps[key] = position
      }
    }
  }
  return steps
}
func (k *Knight) getPosition() Position {
  return k.position
}
func (kp *Knight) setPosition(position Position) {
  kp.position = position
}
func (k *Knight) getColor() string {
  return k.color
}
func (k *Knight) getType() string {
  return "knight"
}


type Rook struct{
  position Position
  color string
}

func (r *Rook) getPossibleSteps(g Game) map[string]Position {
  var steps = map[string]Position{}
  for key, position := range g.positions{
    if position.Num == r.position.Num{
      steps[key] = position
    } else if position.Char == r.position.Char{
        steps[key] = position
      }
    }
  //dont jumpe
  for _, piece := range g.pieces {
    if piece.getPosition() != r.getPosition() {
      if piece.getPosition().Num > r.getPosition().Num && piece.getPosition().Char == r.getPosition().Char {
        for key, p := range g.positions{
          for _, pos := range steps {
            if pos == p {
              if p.Num > piece.getPosition().Num{
                delete(steps, key)
              }
            }
          }
        }
      } else if piece.getPosition().Num < r.getPosition().Num && piece.getPosition().Char == r.getPosition().Char {
          for key, p := range g.positions{
            for _, pos := range steps {
              if pos == p {
                if p.Num < piece.getPosition().Num{
                  delete(steps, key)
                }
              }
            }
          }
      } else if piece.getPosition().Char < r.getPosition().Char && piece.getPosition().Num == r.getPosition().Num {
          for key, p := range g.positions{
            for _, pos := range steps {
              if pos == p {
                if p.Char < piece.getPosition().Char{
                  delete(steps, key)
                }
              }
            }
          }
      } else if piece.getPosition().Char > r.getPosition().Char && piece.getPosition().Num == r.getPosition().Num {
          for key, p := range g.positions{
            for _, pos := range steps {
              if pos == p {
                if p.Char > piece.getPosition().Char{
                  delete(steps, key)
                }
              }
            }
          }
        }
    }
  }
  return steps
}

func (r *Rook) getPosition() Position {
  return r.position
}
func (r *Rook) setPosition(position Position) {
  r.position = position
}
func (r *Rook) getColor() string {
  return r.color
}
func (r *Rook) getType() string {
  return "rook"
}


type Pawn struct{
  position Position
  color string
}

func (pawn *Pawn) getPossibleSteps(g Game) map[string]Position {
  var steps = map[string]Position{}
  position := pawn.getPosition()
  var numberSign int = 1;
  var indicator int = 7;
  if pawn.getColor()  == "white"{
    numberSign = -1;
    indicator = 2;
  }

  if position.Num == indicator {
    for key, p := range g.positions{
      if p.Char == position.Char && p.Num + (numberSign * 1) == position.Num {
        steps[key] = p
      } else if p.Char == position.Char && p.Num + (numberSign * 2) == position.Num {
        steps[key] = p
        //check for jumped
        for _, piece := range g.pieces {
          if piece.getPosition().Char == p.Char && piece.getPosition().Num == p.Num + (numberSign * 1) {
            delete(steps, key)
          }
        }
      }
    }
  } else {
    for key, p := range g.positions{
      if p.Char == position.Char && p.Num + (numberSign * 1) == position.Num {
        steps[key] = p
      }
    }
  }

  for _, piece := range g.pieces {
    //stellt sicher, dass der bauer nur dann zur seite ziehen kann, wennd ort ein Gegner steht
    for key, position := range steps {
      if piece.getPosition() == position {
        delete(steps, key)
      }
    }
    //stellt sicher, dass der bauer nur zur seite ziehen kann, wenn dort ein gegner steht
    for key, p := range g.positions {
      if p == piece.getPosition() {
        if piece.getPosition().Char - 1 == position.Char && piece.getPosition().Num + (numberSign* 1) == position.Num {
          if piece.getColor() != pawn.getColor() {
            steps[key] = piece.getPosition()
          }
        } else if piece.getPosition().Char + 1 == position.Char && piece.getPosition().Num + (numberSign * 1) == position.Num{
          if piece.getColor() != pawn.getColor() {
            steps[key] = piece.getPosition()
          }
        }
      }
    }
  }
  return steps
}


func (p *Pawn) getPosition() Position {
  return p.position
}
func (pp *Pawn) setPosition(position Position) {
  pp.position = position
}
func (p *Pawn) getColor() string {
  return p.color
}
func (p *Pawn) getType() string {
  return "pawn"
}


type King struct{
  position Position
  color string
}

func (k *King) getPossibleSteps(g Game) map[string]Position {
  var steps = map[string]Position{}
  for key, position := range g.positions {
    if position.Num == k.position.Num && position.Char == k.position.Char + 1{
        steps[key] = position;

    } else if position.Num == k.position.Num && position.Char == k.position.Char - 1{
        steps[key] = position;

    } else if position.Num == k.position.Num + 1{
        if position.Char == k.position.Char + 1 {
          steps[key] = position;

        } else if position.Char == k.position.Char - 1  {
          steps[key] = position;

        } else if position.Char == k.position.Char{
          steps[key] = position;
        }

    } else if position.Num == k.position.Num - 1 {
        if position.Char == k.position.Char + 1 {
            steps[key] = position;

        } else if position.Char == k.position.Char - 1 {
          steps[key] = position;

        } else if position.Char == k.position.Char {
          steps[key] = position;
        }
      }
  }
  return steps
}
func (k *King) getPosition() Position {
  return k.position
}
func (kp *King) setPosition(position Position) {
  kp.position = position
}
func (k *King) getColor() string {
  return k.color
}
func (k *King) getType() string {
  return "king"
}


type Queen struct{
  position Position
  color string
}

func (q *Queen) getPossibleSteps(g Game) map[string]Position {
  var steps = map[string]Position{}
  for key, position := range g.positions{
    if math.Abs(float64(position.Char - q.position.Char)) == math.Abs(float64(position.Num - q.position.Num)){
      steps[key] = position
    } else if position.Num == q.position.Num{
      steps[key] = position
    } else if position.Char == q.position.Char{
        steps[key] = position
      }
    }
    //dont jumpe
    for _, piece := range g.pieces {
      if piece.getPosition() != q.getPosition() {
        if piece.getPosition().Num > q.getPosition().Num && piece.getPosition().Char == q.getPosition().Char {
          for key, p := range g.positions{
            for _, pos := range steps {
              if pos == p {
                if p.Num > piece.getPosition().Num && p.Char == q.getPosition().Char{
                  delete(steps, key)
                }
              }
            }
          }
        } else if piece.getPosition().Num < q.getPosition().Num && piece.getPosition().Char == q.getPosition().Char {
            for key, p := range g.positions{
              for _, pos := range steps {
                if pos == p {
                  if p.Num < piece.getPosition().Num && p.Char == q.getPosition().Char{
                    delete(steps, key)
                  }
                }
              }
            }
        } else if piece.getPosition().Char < q.getPosition().Char && piece.getPosition().Num == q.getPosition().Num {
            for key, p := range g.positions{
              for _, pos := range steps {
                if pos == p {
                  if p.Char < piece.getPosition().Char && p.Num == q.getPosition().Num{
                    delete(steps, key)
                  }
                }
              }
            }
        } else if piece.getPosition().Char > q.getPosition().Char && piece.getPosition().Num == q.getPosition().Num {
            for key, p := range g.positions{
              for _, pos := range steps {
                if pos == p {
                  if p.Char > piece.getPosition().Char && p.Num == q.getPosition().Num{
                    delete(steps, key)
                  }
                }
              }
            }
          }
      }
    }
    for _, piece := range g.pieces{
      if piece.getPosition() != q.getPosition(){
        for _, step := range steps {
          if step == piece.getPosition(){

  				  var charDifference int = piece.getPosition().Char - q.getPosition().Char
  				  var numDifference int = piece.getPosition().Num - q.getPosition().Num

  					if charDifference <= -1 && numDifference >= 1 {
              for _, position := range g.positions{
  				      for stepKey, step := range steps {
  								if step == position {
  								  if position.Char - q.getPosition().Char < charDifference && position.Num - q.getPosition().Num > numDifference {
  								    delete(steps, stepKey)
  									}
  								}
  							}
  						}
  					}	else if charDifference >= 1 && numDifference >= 1 {
              for _, position := range g.positions{
  				      for stepKey, step := range steps {
  								if step == position {
  								  if position.Char - q.getPosition().Char > charDifference && position.Num - q.getPosition().Num > numDifference {
  								    delete(steps, stepKey)
  									}
  								}
  							}
  						}
  					}	else if charDifference <= -1 && numDifference <= -1 {
              for _, position := range g.positions{
  				      for stepKey, step := range steps {
  								if step == position {
  								  if position.Char - q.getPosition().Char < charDifference && position.Num - q.getPosition().Num < numDifference {
  								    delete(steps, stepKey)
  									}
  								}
  							}
  						}
            }	else if charDifference >= 1 && numDifference <= -1 {
              for _, position := range g.positions{
  				      for stepKey, step := range steps {
  								if step == position {
  								  if position.Char - q.getPosition().Char > charDifference && position.Num - q.getPosition().Num < numDifference {
  								    delete(steps, stepKey)
  									}
  								}
  							}
  						}
  					}
  				}
  			}
  		}
  	}
  return steps
}
func (q *Queen) getPosition() Position {
  return q.position
}
func (qp *Queen) setPosition(position Position) {
  qp.position = position
}
func (q *Queen) getColor() string {
  return q.color
}
func (q *Queen) getType() string {
  return "queen"
}
