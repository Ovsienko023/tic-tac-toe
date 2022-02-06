package main

import "TicTacToeServer/pkg/httpServer"

func main() {
	//host := "0.0.0.0"
	//port := 8080
	//
	//socketServer.Run(host, port)
	//token.Token()
	httpServer.Run("localhost:8888")
}
