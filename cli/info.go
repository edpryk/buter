package cli

import "fmt"

const author = "Anonym"
const email = "Anonym"
const version = "0.0.2"

func PrintInfo() {
	fmt.Println()
	fmt.Printf("%-10s %s\n", "Version:", version)
	fmt.Printf("%-10s %s\n", "Contact:", email)
	fmt.Printf("%-10s %s \n", "Author:", author)
	fmt.Println()
}
