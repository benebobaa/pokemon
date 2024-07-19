package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokemon_solid/internal/util"
)

const passwordAdmin = "12345"

func main() {
	var username string

	// init
	userHandler, pokemonHandler := initHandler()
	log.Println("success init")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n\nYour username: ")
	username = util.InputCli(scanner)
	ok := userHandler.Access(username)

	if ok {

		fmt.Print("\n\n")
		util.Menu()
		for {
			fmt.Print(username, " > ")
			input := util.InputCli(scanner)

			switch input {
			case "1":
				fmt.Print("Password: ")
				scanner.Scan()
				password := scanner.Text()

				if password != passwordAdmin {
					fmt.Println("\nInvalid password admin")
					continue
				}

				userHandler.FindAll()

			case "2":
				pokemonHandler.FindAll()
			case "3":
				fmt.Println("3")
			case "6":
				fmt.Println("Thankyou!", username)
				return
			default:
				fmt.Print("\nWrong input menu!\n")
				util.Menu()
			}
		}
	} else {
		fmt.Println("Please try again")
	}
}
