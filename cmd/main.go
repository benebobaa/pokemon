package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"pokemon_solid/internal/domain"
	"pokemon_solid/internal/domain/request"
	"pokemon_solid/internal/util"
	"time"
)

const passwordAdmin = "12345"

func main() {
	var user domain.User

	// init
	userHandler, pokemonHandler, collectionsHandler := initHandler()
	log.Println("success init")

	scanner := bufio.NewScanner(os.Stdin)
	util.InitMenu()

	for {
		fmt.Print("Input: ")
		input := util.InputCli(scanner)

		switch input {
		case "1":
			fmt.Print("\n\nYour username: ")
			user.Username = util.InputCli(scanner)
			ok := userHandler.Access(&user)

			if ok {

				fmt.Println("Welcome! ", user.Name)

				fmt.Print("\n\n")
				util.Menu()

				for {
					fmt.Print(user.Username, " > ")
					input := util.InputCli(scanner)

					if input == "7" {
						fmt.Println("Thankyou!", user.Username)
						break
					}

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
							continue
						}

						catchRequest := request.CatchRequest{
							UserID:    user.ID,
							PokemonID: *pokemonId,
						}
						collectionsHandler.Catch(catchRequest)
					case "4":
						fmt.Print("Password: ")
						scanner.Scan()
						password := scanner.Text()

						if password != passwordAdmin {
							fmt.Println("\nInvalid password admin")
							continue
						}

						collectionsHandler.FindAll()
					case "5":
						collectionsHandler.FindAllByUserID(user.ID)
					case "6":
						fmt.Print("Input collection id: ")
						collectionId, err := util.InputCliNumber(scanner)

						if err != nil {
							fmt.Println("Collection id must a number!")
							continue
						}

						releaseRequest := request.ReleaseRequest{
							UserID:       user.ID,
							CollectionID: *collectionId,
						}

						log.Println("release req :: ", releaseRequest)

						collectionsHandler.Release(releaseRequest)
					default:
						fmt.Print("\nWrong input menu!\n")
						util.Menu()
					}
				}
			} else {
				fmt.Println("Please try again")
			}
		case "2":
			fmt.Print("Exiting")
			for i := 0; i < 5; i++ {
				fmt.Print(" . ")
				time.Sleep(300 * time.Millisecond)
			}
			fmt.Print("\n\n")
			return
		default:
			fmt.Println("Your input wrong! ")
			util.InitMenu()
		}
	}
}
