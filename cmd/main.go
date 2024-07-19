package main

import (
	"bufio"
	"fmt"
	"os"
	"pokemon_solid/internal/handler"
	"pokemon_solid/internal/repository"
	"pokemon_solid/internal/usecase"
	"pokemon_solid/internal/util"
)

const passwordAdmin = "12345"

func main() {
	var username string

	// init
	ur := repository.NewUserRepository()
	uc := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uc)
	generateUsers(uh)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\n\nYour username: ")
	username = util.InputCli(scanner)
	ok := uh.Access(username)

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

				uh.FindAll()

			case "2":
				fmt.Println("Test")

			case "5":
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
