package main

import (
	"keyValueStore/src/storage"
	"keyValueStore/src/ui"
)

func main() {
	s := storage.InMemoryKVStorage{}.NewStorage()
	cli := ui.CommandLineInterface{}.NewCommandLineInterface(s)
	cli.Run()
}
