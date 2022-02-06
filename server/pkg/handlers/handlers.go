package handlers

import (
	"TicTacToeServer/pkg/repository"
	"TicTacToeServer/pkg/users"
	"encoding/json"
	"fmt"
	"net"
)

func HandleClient(conn net.Conn, db *repository.DataBase) {
	defer conn.Close() // закрываем сокет при выходе из функции

	user := users.User{}
	buf := make([]byte, 128) // буфер для чтения клиентских данных
	for {
		readLen, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			fmt.Println(err)
			break
		}
		err = json.Unmarshal(buf[:readLen], &user)
		if err != nil {
			panic(err)
		}

		for _, data := range buf[:readLen] {
			fmt.Print(string(data))
		}
		//conn.Write(append([]byte("Goodbye, "), buf[:readLen]...)) // пишем в сокет
	}
}
