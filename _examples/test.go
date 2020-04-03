package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
)

func GetCurrentUser(username *string) error {

	currentUser, err := user.Current()

	if err != nil {
		return err
	}

	*username = currentUser.Username

	return nil
}

func main() {
	var username string

	err := GetCurrentUser(&username)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println("username: ",username)

	Print("KatWeb " + currentVersion + ", built for " + runtime.GOOS + "-" + runtime.GOARCH + ", using " + runtime.Compiler + " compiler.")
}
