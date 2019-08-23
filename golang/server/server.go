package server

import (
	"fmt"
)

func main() {
	username := "002099"
	password := "1"
	userID, err := lis.ValidUser(username, password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(userID)
	}
}
