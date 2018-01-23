package http

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/m110/canvas-game/pkg/game/player"
	"github.com/m110/canvas-game/pkg/game/position"
)

type PositionResponse struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func NewPositionResponse(position position.Position) PositionResponse {
	return PositionResponse{
		X: position.X(),
		Y: position.Y(),
	}
}

type PlayerResponse struct {
	ID       string           `json:"id"`
	Position PositionResponse `json:"position"`
}

func NewPlayerResponse(player player.Player) PlayerResponse {
	return PlayerResponse{
		ID:       player.ID(),
		Position: NewPositionResponse(player.Position()),
	}
}

func (PlayerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type PlayerListResponse []PlayerResponse

func NewPlayerListResponse(players []player.Player) []render.Renderer {
	list := []render.Renderer{}

	for _, p := range players {
		list = append(list, NewPlayerResponse(p))
	}

	return list
}

type PlayerEvent struct {
	Name   string         `json:"name"`
	Player PlayerResponse `json:"player"`
}

func NewPlayerEvent(name string, player player.Player) PlayerEvent {
	return PlayerEvent{
		Name:   name,
		Player: NewPlayerResponse(player),
	}
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}
