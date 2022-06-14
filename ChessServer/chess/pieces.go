package chess

import (
  "math"
)

type Piece interface {
  getPosition() Position
  setPosition(position Position)
  getColor() string
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


type Pawn struct{
  position Position
  color string
}

func (p *Pawn) getPossibleSteps(g Game) map[string]Position {
  var steps = map[string]Position{"A3": Position{65,3}}
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
