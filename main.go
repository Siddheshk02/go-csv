package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	// Open the CSV file
	file, err := os.Open("people.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Convert the CSV records to a slice of Person structs
	people := make([]Person, 0, len(records)-1)
	for i, record := range records {
		if i == 0 {
			// Skip header row
			continue
		}
		age := 0
		fmt.Sscanf(record[1], "%d", &age)
		people = append(people, Person{
			Name:  record[0],
			Age:   age,
			Email: record[2],
		})
	}

	// Output the data in a formatted table
	fmt.Printf("| %-20s | %-5s | %-30s |\n", "Name", "Age", "Email")
	fmt.Println("|----------------------|-------|--------------------------------|")
	for _, person := range people {
		fmt.Printf("| %-20s | %-5d | %-30s |\n", person.Name, person.Age, person.Email)
	}
}
