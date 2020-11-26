package ui

import (
	"bufio"
	"fmt"
	"keyValueStore/src/storage"
	"os"
	"strings"
)

type CommandLineInterface struct {
	storage storage.KVStorage
	reader  *bufio.Reader
}

func (i CommandLineInterface) NewCommandLineInterface(s storage.KVStorage) CommandLineInterface {
	fmt.Println("building cli for", s)

	return CommandLineInterface{storage: s, reader: bufio.NewReader(os.Stdin)}
}

func (i CommandLineInterface) readKeyFromStdIn() (string, error) {
	fmt.Print("Enter key: ")
	key, err := i.reader.ReadString('\n')
	return key, err
}

func (i CommandLineInterface) readValueFromStdIn() (string, error) {
	fmt.Print("Enter value: ")
	value, err := i.reader.ReadString('\n')
	return value, err
}

func (i CommandLineInterface) handleCreateCommand() {
	key, err := i.readKeyFromStdIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := i.readValueFromStdIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = i.storage.CreateString(key, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
}

func (i CommandLineInterface) handleReadCommand() {
	key, err := i.readKeyFromStdIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := i.storage.ReadString(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}

func (i CommandLineInterface) handleUpdateCommand() {
	key, err := i.readKeyFromStdIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	value, err := i.readValueFromStdIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = i.storage.UpdateString(key, value)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
}

func (i CommandLineInterface) handleDeleteCommand() {
	key, err := i.readKeyFromStdIn()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = i.storage.DeleteString(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
}

func (i CommandLineInterface) handlePrintCommand() {
	err := i.storage.PrintDump()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (i CommandLineInterface) Run() {

	for {
		fmt.Print("Enter command: ")
		command, _ := i.reader.ReadString('\n')
		cmd := strings.ToLower(strings.TrimSpace(command))

		switch cmd {
		case "create":
			i.handleCreateCommand()
			continue
		case "read":
			i.handleReadCommand()
			continue
		case "update":
			i.handleUpdateCommand()
			continue
		case "delete":
			i.handleDeleteCommand()
			continue
		case "print":
			i.handlePrintCommand()
			continue
		case "exit":
			i.handleExitCommand()
		default:
			fmt.Println("not implemented")
			continue
		}
	}

}

func (i CommandLineInterface) handleExitCommand() {
	fmt.Println("exiting")
	os.Exit(0)
}
