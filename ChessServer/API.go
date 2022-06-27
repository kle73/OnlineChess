package main

import (
  "encoding/json"
  "Chess/chess"
  "net"
  "fmt"
)

const (
  checkmate string = "CHECKMATE"
  set string = "SET"
  setCheck string = "SET CHECK"
  black string= "black"
  white string = "white"
)


// takes JSON like Object: {"pos1":"E2", "pos2": E3}
type Input struct{
  Pos1 string
  Pos2 string
}
// takes JSON like Object: {"pos1": "E2", "pos2": "E3", "Option1": "CHECKMATE", "Option2": true / false}
type Response struct{
  Pos1 string
  Pos2 string
  Option1 string
  Option2 string

}

type Player struct {
  connection net.Conn
  color string
}


func readPlayer(player Player) Input {
  var input Input
  rcv := make([]byte, 28)
  _, err := player.connection.Read(rcv)
  if err != nil {
      player.connection.Close()
  }
  data := string(rcv)
  json.Unmarshal([]byte(data), &input)
  return input
}

func writePlayer(player Player, message string){
  player.connection.Write([]byte(message))
}

func handleGame(player1 Player, player2 Player){
  fmt.Println("New Game startet")
  players := []Player{player1, player2}
  //give each player its color:
  for _, player := range players{
    writePlayer(player, player.color)
  }
  //Game starts:
  game := chess.NewGame()
  var playing bool = true
  //game loop:
  for playing {
    var message string
    for i, player := range players{
      if player.color == game.Color {
        //get move from player, whose turn it is
        var move Input = readPlayer(player)
        //make the move in the chss engine and send the response back to the player
        var response string
        var option string
        game, response, option = chess.SetPiece(move.Pos1, move.Pos2, game)
        if option == "0" {
          writePlayer(player, response)
        } else {
          writePlayer(player, response + " 1")
        }
        //send the move to the opponent of the player (normal case)
        if response == set{
          if option == "0"{
            jsonResp, _ := json.Marshal(move)
            message = string(jsonResp)
          }else {
            resp := Response{move.Pos1, move.Pos2, response, option}
            jsonResp, _ := json.Marshal(resp)
            message = string(jsonResp)
          }
          if i == 0 {
            writePlayer(players[1], string(message))
          } else {
            writePlayer(players[0], string(message))
          }
          //(special case)
        } else if response == setCheck || response == checkmate {

          resp := Response{move.Pos1, move.Pos2, response, option}
          jsonResp, _ := json.Marshal(resp)
          message = string(jsonResp)
          if i == 0 {
            writePlayer(players[1], string(message))
          } else {
            writePlayer(players[0], string(message))
          }
        }
        break
      }
    }
  }
  //close connections after game
  for _, player := range players {
    player.connection.Close()
  }
}

func main(){
  newPlayers := []Player{}
  ln, _ := net.Listen("tcp", ":5050")
  for {
    conn, _ := ln.Accept()
    if len(newPlayers) == 0 {
      newPlayers = append(newPlayers, Player{conn, white})
      fmt.Println("new Player connected!")
    } else {
      newPlayers = append(newPlayers, Player{conn, black})
      go handleGame(newPlayers[0], newPlayers[1])
      newPlayers = []Player{}
    }
  }
}
