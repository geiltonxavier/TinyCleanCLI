package main

import (
	"log"
	"os"

	"github.com/geiltonxavier/TinyCleanCLI/internal/cli"
)

func main() {
	if err := cli.Execute(os.Args); err != nil {
		log.Fatal(err)
	}
}
