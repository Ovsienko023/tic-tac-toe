package repository

import (
	"net"
)

type DB interface {
	Init()
	GetUser(login string) UsersDB
}

type UsersDB struct {
	Login      string
	Password   string
	connection net.Conn
}

type DataBase struct {
	users []UsersDB
}

func (db *DataBase) Init() {
	db.users = append(
		db.users,
		UsersDB{Login: "bob", Password: "9001"},
		UsersDB{Login: "kop", Password: "9002"},
	)
}

func (db *DataBase) GetUser(login string) UsersDB {
	for _, user := range db.users {
		if user.Login == login {
			return user
		}
	}
	return UsersDB{}
}
