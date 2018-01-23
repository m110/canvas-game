package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/m110/canvas-game/pkg/broker"
	"github.com/m110/canvas-game/pkg/game/board"
	"github.com/m110/canvas-game/pkg/game/player"
)

func AddPlayerResource(r *chi.Mux, s PlayerResource) {
	r.Route("/players", func(r chi.Router) {
		r.Get("/", s.GetAll)
		r.Post("/{playerID}", s.Join)
		r.Post("/{playerID}/move/{direction}", s.MovePlayer)
		r.Get("/updates", s.FetchPlayerUpdates)
	})
}

type PlayerResource struct {
	board  *board.Board
	broker *broker.Broker
}

func NewPlayerResource(board *board.Board, broker *broker.Broker) PlayerResource {
	return PlayerResource{board, broker}
}

func (s PlayerResource) GetAll(w http.ResponseWriter, r *http.Request) {
	playersResponse := NewPlayerListResponse(s.board.Players())
	if err := render.RenderList(w, r, playersResponse); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

func (s PlayerResource) Join(w http.ResponseWriter, r *http.Request) {
	playerID := chi.URLParam(r, "playerID")
	s.board.SpawnPlayer(playerID)
}

func (s PlayerResource) MovePlayer(w http.ResponseWriter, r *http.Request) {
	playerID := chi.URLParam(r, "playerID")
	direction := chi.URLParam(r, "direction")
	s.board.MovePlayer(playerID, direction)
}

func (s PlayerResource) FetchPlayerUpdates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		render.Render(w, r, ErrRender(errors.New("flusher not supported")))
		return
	}

	notifier, ok := w.(http.CloseNotifier)
	if !ok {
		render.Render(w, r, ErrRender(errors.New("notifier not supported")))
		return
	}

	playerChan := s.broker.Subscribe(player.PlayerEvent{})

	running := true
	notify := notifier.CloseNotify()
	go func() {
		<-notify
		s.broker.Disconnect(playerChan)
		running = false
		log.Println("Client disconnected")
	}()

	for running {
		select {
		case event, more := <-playerChan:
			if !more {
				continue
			}

			playerEvent := event.(player.PlayerEvent)
			if !ok {
				log.Println("Failed to cast as PlayerEvent:", playerEvent)
				continue
			}

			response := NewPlayerEvent(playerEvent.Name(), playerEvent.Player())
			data, err := json.Marshal(response)
			if err != nil {
				log.Println("Error marshaling response:", err)
				continue
			}

			fmt.Fprintf(w, "data: %s\n\n", data)
			flusher.Flush()
		}
	}

	log.Println("End of event loop")
}
