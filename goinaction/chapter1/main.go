package main

import (
	"go-code/goinaction/chapter1/search"
	"log"
	"os"

)

func main() {
	search.Run("president")
}

func init() {
	log.SetOutput(os.Stdout)
}
