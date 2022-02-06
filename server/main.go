package main

import (
	"TicTacToeServer/pkg/server"
)

func main() {
	host := "0.0.0.0"
	port := 8080

	server.Run(host, port)
}
