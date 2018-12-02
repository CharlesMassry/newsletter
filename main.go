package main

import (
	"os"
)

func main() {
	languages := os.Args[1:]
	newsletter(languages)
}
