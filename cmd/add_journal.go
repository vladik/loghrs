package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func AddJournal(db *sql.DB) {
	// Prompt for journal name
	namePrompt := promptui.Prompt{
		Label: "Enter journal name",
	}
	journalName, err := namePrompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Load shelves (optional)
	shelves := []string{"(None)"}
	shelfIDs := []int{0}

	rows, err := db.Query("SELECT id, name FROM shelves ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		shelfIDs = append(shelfIDs, id)
		shelves = append(shelves, name)
	}

	shelves = append(shelves, "+ Create new shelf")
	shelfIDs = append(shelfIDs, -1) // placeholder for creation

	// Prompt for shelf selection
	prompt := promptui.Select{
		Label: "Assign to shelf?",
		Items: shelves,
	}
	i, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	var shelfID interface{} = nil
	if shelfIDs[i] == -1 {
		// Create new shelf
		createPrompt := promptui.Prompt{
			Label: "Enter new shelf name",
		}
		newShelfName, err := createPrompt.Run()
		if err != nil {
			log.Fatal(err)
		}
		res, err := db.Exec("INSERT INTO shelves (name) VALUES (?)", newShelfName)
		if err != nil {
			log.Fatal(err)
		}
		newID, _ := res.LastInsertId()
		shelfID = int(newID)
	} else if shelfIDs[i] != 0 {
		shelfID = shelfIDs[i]
	}

	// Insert journal
	_, err = db.Exec("INSERT INTO journals (name, shelf_id) VALUES (?, ?)", journalName, shelfID)
	if err != nil {
		log.Fatalf("Failed to add journal: %v", err)
	}

	fmt.Println("Journal added successfully.")
}