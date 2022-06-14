package main

import (
  "encoding/json"
  "Chess/chess"
  "net"
  "fmt"
)


// takes JSON like Object: {"pos1":"E2", "pos2": E3}
type Data struct{
  Pos1 string
  Pos2 string
}

type Player struct {
  connection net.Conn
  color string
}


func main(){
  newPlayers := []Player{}
  ln, _ := net.Listen("tcp", ":5050")
  for {
    conn, _ := ln.Accept()
    if len(newPlayers) == 0 {
      newPlayers = append(newPlayers, Player{conn, "white"})
      fmt.Println("new Player connected!")
    } else {
      newPlayers = append(newPlayers, Player{conn, "black"})
      go handleGame(newPlayers[0], newPlayers[1])
      newPlayers = []Player{}
    }
  }
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
    var response string
    for i, player := range players{
      if player.color == game.Color {
        var turn Data = readPlayer(player)
        var msg string
        game, msg = chess.SetPiece(turn.Pos1, turn.Pos2, game)
        writePlayer(player, msg)
        if msg == "SET"{
          m, _ := json.Marshal(turn)
          response = string(m)
          if i == 0 {
            writePlayer(players[1], string(response))
          } else {
            writePlayer(players[0], string(response))
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

func readPlayer(player Player) Data {
  var data Data
  rcv := make([]byte, 28)
  _, err := player.connection.Read(rcv)
  if err != nil {
      player.connection.Close()
  }
  d := string(rcv)
  json.Unmarshal([]byte(d), &data)
  return data
}

func writePlayer(player Player, data string){
  player.connection.Write([]byte(data))
}
