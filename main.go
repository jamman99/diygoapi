package main

import (
	"fmt"
	"os"

	"github.com/gilcrest/go-api-basic/command"
)

func main() {
	if err := command.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error from commands.Run(): %s\n", err)
		os.Exit(1)
	}
}
