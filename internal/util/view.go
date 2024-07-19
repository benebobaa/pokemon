package util

import (
	"fmt"
	"pokemon_solid/internal/domain"
	"strings"
)

func TableUser(data []domain.User) {

	fmt.Printf("%-5s %-20s %-20s\n", "ID", "Username", "Name")
	fmt.Println(strings.Repeat("-", 55))

	for _, user := range data {
		fmt.Printf("%-5d %-20s %-20s\n", user.ID, user.Username, user.Name)
	}
	fmt.Print("\n\n")
}

func TablePokemon(data []domain.Pokemon) {

	fmt.Printf("%-5s %-20s %-10s %-20s\n", "ID", "Name", "Presence", "Created At")
	fmt.Println(strings.Repeat("-", 55))

	// for _, emp := range data {
	// 	fmt.Printf("%-5d %-20s %-10t %-20s\n", emp.ID, emp.Name, emp.Presence, emp.CreatedAt)
	// }
	fmt.Print("\n\n")
}

func Menu() {
	fmt.Println("[1] View All Users (admin only)")
	fmt.Println("[2] View All Pokemons Available")
	fmt.Println("[3] Delete employee by id")
	fmt.Println("[4] Update employee by id")
	fmt.Println("[5] Exit the program")
}
