package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]
		switch command {
		case "exit":
			fmt.Println("Closing the Pokedex... Goodbye!")
			os.Exit(0)

		case "help":
			fmt.Println("Welcome to the Pokedex!")
			fmt.Println("Usage: ")
			fmt.Println("help: Displays a help message")
			fmt.Println("exit: Exit the Pokedex")
			fmt.Println()

		default:
			fmt.Printf("Your command was: %v\n", command)
		}
	}
}

type cliCommand struct {

}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered);
	return words	
}