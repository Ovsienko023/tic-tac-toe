package client

import (
	users "TicTacToeClient/pkg/users"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

//type User struct {
//	Login    string
//	Password string
//	Data     int
//}

func ConnReader(c net.Conn) {
	//data := make([]byte, 128)
	//_, err := c.Read(data)
	message, _ := bufio.NewReader(c).ReadString('}')
	//if _, err := io.Copy(message, src); err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(message)

	//fmt.Println(string(data))
}

func MoveOpponent(data int) int {
	//data := make([]byte, 128)
	user := users.User{Login: "bob", Password: "9001"}
	user.Data = data
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	conn, _ := net.Dial("tcp", "192.168.31.192:8080")
	conn.Write(b)
	go ConnReader(conn)
	return 1
}

func main1() {
	user := users.User{Login: "bob", Password: "9001"}
	b, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	serv := os.Args[1]               // берем адрес сервера из аргументов командной строки
	conn, _ := net.Dial("tcp", serv) // открываем TCP-соединение к серверу
	//conn.Write([]byte(""))
	fmt.Println("!!!!")
	conn.Write(b)
	go copyTo(os.Stdout, conn) // читаем из сокета в stdout
	//fmt.Println(data)
	//conn.Write(b)
	copyTo(conn, os.Stdin) // пишем в сокет из stdin
}

func copyTo(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
