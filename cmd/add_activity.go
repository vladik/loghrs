package cmd

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
)

func AddActivity(db *sql.DB) {
	// Load journals
	rows, err := db.Query("SELECT id, name FROM journals ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var journalIDs []int
	var journalNames []string
	for rows.Next() {
		var id int
		var name string
		r := rows.Scan(&id, &name)
		if r != nil {
			log.Fatal(r)
		}
		journalIDs = append(journalIDs, id)
		journalNames = append(journalNames, name)
	}

	if len(journalNames) == 0 {
		fmt.Println("No journals found. Please create a journal first.")
		return
	}

	// Prompt for journal
	prompt := promptui.Select{
		Label: "Select a journal",
		Items: journalNames,
	}
	i, _, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	journalID := journalIDs[i]

	// Load existing activities for this journal
	rows, err = db.Query("SELECT id, name FROM activities WHERE journal_id = ? ORDER BY name", journalID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	activityNames := []string{"+ Create new activity"}
	for rows.Next() {
		var name string
		r := rows.Scan(new(int), &name)
		if r != nil {
			log.Fatal(r)
		}
		activityNames = append(activityNames, name)
	}

	// Prompt to select or create
	prompt = promptui.Select{
		Label: "Choose or create an activity",
		Items: activityNames,
	}
	selected, selectedName, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	var activityName string
	if selected == 0 {
		// Create new
		namePrompt := promptui.Prompt{
			Label: "Enter activity name",
		}
		activityName, err = namePrompt.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		activityName = selectedName
	}

	// Insert if it's a new one
	if selected == 0 {
		_, err = db.Exec("INSERT INTO activities (name, journal_id) VALUES (?, ?)", activityName, journalID)
		if err != nil {
			log.Fatalf("Failed to add activity: %v", err)
		}
		fmt.Println("Activity added successfully.")
	} else {
		fmt.Printf("Selected existing activity: %s\n", activityName)
	}
}