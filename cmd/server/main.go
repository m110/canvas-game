package main

import (
	"net/http"
	"time"

	"github.com/m110/canvas-game/pkg/broker"
	"github.com/m110/canvas-game/pkg/game/board"
	pkg_http "github.com/m110/canvas-game/pkg/http"
)

func main() {
	broker := broker.NewBroker()
	board := board.NewBoard(broker)

	board.SpawnPlayer("AAAAAA")
	board.SpawnPlayer("BBBBBB")
	board.SpawnPlayer("CCCCCC")
	board.SpawnPlayer("DDDDDD")
	board.SpawnPlayer("EEEEEE")

	go func() {
		i := 1
		for {
			//board.SpawnPlayer(player.SpawnPlayer(fmt.Sprintf("player-%d", i)))
			i++
			time.Sleep(time.Second * 1)
		}
	}()

	r := pkg_http.NewRouter()

	playerResource := pkg_http.NewPlayerResource(board, broker)
	pkg_http.AddPlayerResource(r, playerResource)

	http.ListenAndServe(":3000", r)
}
