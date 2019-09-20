package board

import (
	"fmt"
	"strings"
	"encoding/json"
	"github.com/gorilla/websocket"
)

// No constructors: https://stackoverflow.com/a/48446890/3339274

type Box struct {
	Enabled bool
	TileNum int
	TileId int
	Owner int
}

type Player struct {
	SessionId string
	Connection websocket.Conn // https://godoc.org/github.com/gorilla/websocket#Conn
	Name string
	Score int
	HasLost bool
}

type SimplifiedPlayer struct {
	Name string
	Score int
	HasLost bool
}

type Board struct {
	Boxes [][]Box
	NextTileId int
	Players []Player
	// GameBeganAt
	GameIsWon bool
}

func removeFromPlayerSlice(slice []Player, i int) []Player {
  copy(slice[i:], slice[i+1:])
  return slice[:len(slice)-1]
}

func (brd Board) addPlayer(sessionId string, connection websocket.Conn, name string) {
	brd.Players = append(brd.Players, Player{SessionId: sessionId, Connection: connection, Name: name})
}

func (brd Board) removePlayer(sessionId string) {
	for i := 0; i < len(brd.Players); i++ {
		if (strings.Compare(brd.Players[i].SessionId, sessionId) == 0) {
			brd.Players = removeFromPlayerSlice(brd.Players, i)
			return
		}
	}
}

func (brd Board) isBoardFull() bool { return len(brd.Players) >= 4 }

func (brd Board) getWhenGameBegan() { /* TODO(Neil): IMPLEMENT THIS*/ }

func (brd Board) checkIfGameWon() int {
	var playersStillInGame, playerId int
	playersStillInGame = 0
	playerId = -1

	for i := 0; i < len(brd.Players); i++ {
		if (!brd.Players[i].HasLost) {
			playersStillInGame++
			playerId = i
		}
	}

	if (playersStillInGame == 1) {
		return playerId
	} else {
		return -1
	}
}

func (brd Board) getPlayers() []Player {
	return brd.Players
}

func (brd Board) getAsJSON() string {

}

func newBoard() * Board {
	var board Board

	board.NextTileId = 1
	board.GameIsWon = false

	return &board
}
