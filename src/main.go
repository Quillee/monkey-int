package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/quillee/monkey/src/repl"
)


func main() {
    user, err := user.Current()

    if err != nil {
        panic(err)
    }

    fmt.Printf("hello %s! This is the monkey programming language", user.Username)

    repl.Start(os.Stdin, os.Stdout)
}

