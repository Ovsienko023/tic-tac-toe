package auth

import "TicTacToeServer/pkg/repository"

func authentication(db repository.DB, login string, password string) bool {
	user := db.GetUser(login)
	if user.Login == "" {
		return false
	}
	if user.Password == password {
		return true
	}
	return false
}
