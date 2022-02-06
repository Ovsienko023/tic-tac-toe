package repository

import "net"

type UsersDB struct {
	Login      string
	Password   string
	connection net.Conn
}

type DataBase struct {
	user1 UsersDB
	user2 UsersDB
}

//db := repo.DataBase{
//	user1: repo.UsersDB{Login: "bob", Password: "9001"},
//	user2: repo.UsersDB{Login: "kop", Password: "9002"},
//}
