package main

import (
	"fmt"
	"regexp"

)

func testValidUserName(username string) bool {

	validUsername := regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	if len(username) > 3 && len(username) < 39 && validUsername.MatchString(username) {
        fmt.Println("Username is valid", username)
        return true
    } else {
        fmt.Println("Username is invalid, please try again.")
        return false
    }
}

func main() {
	var username string
	for {
		fmt.Print("Enter username: ")
		fmt.Scanln(&username)
		testValidUserName(username)
	}
}