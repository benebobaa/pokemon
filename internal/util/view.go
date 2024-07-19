package util

import (
	"fmt"
	"pokemon_solid/internal/domain"
	"strings"
)

func TableUser(data []domain.User) {

	fmt.Printf("%-5s %-20s %-20s\n", "ID", "Username", "Name")
	fmt.Println(strings.Repeat("-", 36))

	for _, user := range data {
		fmt.Printf("%-5d %-20s %-20s\n", user.ID, user.Username, user.Name)
	}
	fmt.Print("\n\n")
}

func TablePokemon(data []domain.Pokemon) {
	fmt.Print("\n\n")
	fmt.Printf("%-5s %-20s %-10s %-10s %-10s %-20s\n", "ID", "Name", "Type", "CatchRate", "IsRare", "Registered Date")
	fmt.Println(strings.Repeat("-", 75))

	for _, p := range data {
		fmt.Printf("%-5d %-20s %-10s %-10d %-10t %-20s\n", p.ID, p.Name, p.Type, p.CatchRate, p.IsRare, p.RegisteredDate)
	}
	fmt.Print("\n\n")
}

func TableCollection(data []domain.Collections) {
	fmt.Print("\n\n")
	fmt.Printf("%-5s %-20s %-20s %-10s %-10s %-20s\n", "ID", "User Name", "Pokemon Name", "Type", "Is Rare", "Created At")
	fmt.Println(strings.Repeat("-", 85))

	for _, c := range data {
		fmt.Printf("%-5d %-20s %-20s %-10s %-10t %-20s\n", c.ID, c.User.Username, c.Pokemon.Name, c.Pokemon.Type, c.Pokemon.IsRare, c.CreatedAt)
	}
	fmt.Print("\n\n")
}

func Menu() {
	fmt.Println("[1] View All Users (admin only)")
	fmt.Println("[2] View All Pokemons Available")
	fmt.Println("[3] Try Catch a Pokemon")
	fmt.Println("[4] Check your pokemon collections")
	fmt.Println("[5] Release your collection pokemon")
	fmt.Println("[6] Logout user")
}

func InitMenu() {
	fmt.Println("[1] Access with username")
	fmt.Println("[2] Exit the program")
}
