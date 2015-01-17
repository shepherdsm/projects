package main

import (
	"bufio"
	"fmt"
	"os"
)

// Age Distribution
// https://www.codeeval.com/open_challenges/152/
func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var age int
		fmt.Sscanf(scanner.Text(), "%d", &age)

		switch {
		case 0 <= age && age <= 2:
			fmt.Println("Still in Mama's arms")
		case 3 <= age && age <= 4:
			fmt.Println("Preschool Maniac")
		case 5 <= age && age <= 11:
			fmt.Println("Elementary school")
		case 12 <= age && age <= 14:
			fmt.Println("Middle school")
		case 15 <= age && age <= 18:
			fmt.Println("High school")
		case 19 <= age && age <= 22:
			fmt.Println("College")
		case 23 <= age && age <= 65:
			fmt.Println("Working for the man")
		case 66 <= age && age <= 100:
			fmt.Println("The Golden Years")
		default:
			fmt.Println("This program is for humans")
		}
	}
}
