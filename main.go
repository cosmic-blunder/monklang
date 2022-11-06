package main

import (
	"fmt"

	"os"
	"os/user"

	// "github.com/oscuro1111/golang-libs/semaphore"
	// "github.com/google/go-cmp/cmp"
	"github.com/cosmic-blunder/monklang/repl"
)

func main() {

	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey Programming Language!\n", user.Username)

	fmt.Printf("Feel FRee to type in command\n")

	repl.Start(os.Stdin, os.Stdout)

}
