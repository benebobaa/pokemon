package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/domain/request"
	"pokemon_solid/internal/util"
)

const passwordAdmin = "12345"

func main() {
	var user domain.User

	// init
	userHandler, pokemonHandler, collectionsHandler := initHandler()
	log.Println("success init")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n\nYour username: ")
	user.Username = util.InputCli(scanner)
	ok := userHandler.Access(&user)

	fmt.Println("Welcome! ", user.Name)
	if ok {

		fmt.Print("\n\n")
		util.Menu()
		for {
			fmt.Print(user.Username, " > ")
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
				fmt.Println("Letsgo! catch a pokemon")

				fmt.Print("Input pokemon id: ")
				pokemonId, err := util.InputCliNumber(scanner)

				if err != nil {
					fmt.Println("Pokemon id must a number!")
				}

				catchRequest := request.CatchRequest{
					UserID:    user.ID,
					PokemonID: *pokemonId,
				}
				collectionsHandler.Catch(catchRequest)
			case "4":
				collectionsHandler.FindAll()
			case "6":
				fmt.Println("Thankyou!", user.Username)
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
