package game

import (
    "github.com/gorilla/websocket"
)

type player struct {
	connection *websocket.Conn
}
