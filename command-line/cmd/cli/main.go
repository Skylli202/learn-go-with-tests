package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/Skylli202/learn-go-with-tests/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFile()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	poker.NewCLI(store, os.Stdin).PlayPoker()
}
