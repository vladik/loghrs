package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func AddShelf(db *sql.DB) {
	// Prompt for shelf name
	namePrompt := promptui.Prompt{
		Label: "Enter shelf name",
	}
	shelfName, err := namePrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Insert shelf
	_, err = db.Exec("INSERT INTO shelves (name) VALUES (?)", shelfName)
	if err != nil {
		log.Fatalf("Failed to add shelf: %v", err)
	}

	fmt.Println("Shelf added successfully.")
}