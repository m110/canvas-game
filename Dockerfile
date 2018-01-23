FROM golang:1.9-stretch

RUN go get github.com/cespare/reflex

RUN mkdir -p $GOPATH/src/github.com/m110/canvas-game
WORKDIR $GOPATH/src/github.com/m110/canvas-game

CMD reflex -s -r '\.go$' go run cmd/server/main.go
