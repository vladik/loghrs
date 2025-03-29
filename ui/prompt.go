package ui

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func SelectJournalPrompt(db *sql.DB) {
	rows, err := db.Query("SELECT id, name FROM journals ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var ids []int
	var names []string
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		ids = append(ids, id)
		names = append(names, name)
	}

	if len(names) == 0 {
		fmt.Println("No journals found.")
		return
	}

	prompt := promptui.Select{Label: "Select a journal", Items: names}
	i, result, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Selected journal: %s (ID: %d)\n", result, ids[i])
}