package register

import (
	// "fmt"
	
)

func reg_user(name string) {
	user_id := users.AddUser(name)
	TopUpBalance(user_id, 0)
}

func TopUpBalance(userID int, amount int) {
	if users.UserExists(userID) {
		balance.SetBalance(userID, amount)
	}
}