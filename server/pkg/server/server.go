package server

import (
	"TicTacToeServer/pkg/users"
	"encoding/json"
	"fmt"
	"net"
)

func Run(host string, port int) {

	address := fmt.Sprintf("%s:%d", host, port)

	listener, _ := net.Listen("tcp", address) // открываем слушающий сокет
	fmt.Println("Server run: ", address)
	user := users.User{}
	buf := make([]byte, 128)
	for {
		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		if err != nil {
			continue
		}

		readLen, err := conn.Read(buf) // читаем из сокета
		if err != nil {
			fmt.Println(err)
			break
		}
		err = json.Unmarshal(buf[:readLen], &user)
		if err != nil {
			panic(err)
		}
		fmt.Println(user.Login, user.Password, user.Data)
		//if user.Login == db.user1.Login && user.Password == db.user1.Password && db.user1.connection == nil {
		//	db.user1.connection = conn
		//	go handlers.HandleClient(conn, &db)
		//	continue
		//}
		//if user.Login == db.user2.Login && user.Password == db.user2.Password && db.user2.connection == nil {
		//	db.user2.connection = conn
		//	go handleClient(conn, &db)
		//	continue
		//}
	}

}
