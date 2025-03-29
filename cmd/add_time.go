package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
)

func AddTime(db *sql.DB) {
	// Load journals
	jRows, err := db.Query("SELECT id, name FROM journals ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer jRows.Close()

	var journalIDs []int
	var journalNames []string
	for jRows.Next() {
		var id int
		var name string
		jRows.Scan(&id, &name)
		journalIDs = append(journalIDs, id)
		journalNames = append(journalNames, name)
	}

	if len(journalNames) == 0 {
		fmt.Println("No journals found. Please create one first.")
		return
	}

	// Select journal
	journalPrompt := promptui.Select{
		Label: "Select a journal",
		Items: journalNames,
	}
	ji, _, err := journalPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	journalID := journalIDs[ji]

	// Load activities for the journal
	aRows, err := db.Query("SELECT id, name FROM activities WHERE journal_id = ? ORDER BY name", journalID)
	if err != nil {
		log.Fatal(err)
	}
	defer aRows.Close()

	var activityIDs []int
	var activityNames []string
	for aRows.Next() {
		var id int
		var name string
		aRows.Scan(&id, &name)
		activityIDs = append(activityIDs, id)
		activityNames = append(activityNames, name)
	}

	if len(activityNames) == 0 {
		fmt.Println("No activities found for this journal. Please add one first.")
		return
	}

	// Select activity
	activityPrompt := promptui.Select{
		Label: "Select an activity",
		Items: activityNames,
	}
	ai, _, err := activityPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	activityID := activityIDs[ai]

	// Duration input
	durPrompt := promptui.Prompt{
		Label: "Enter duration (e.g. 1h30m)",
	}
	durStr, err := durPrompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	duration, err := time.ParseDuration(durStr)
	if err != nil {
		log.Fatalf("Invalid duration: %v", err)
	}
	durationMinutes := int(duration.Minutes())

	// Note input
	notePrompt := promptui.Prompt{
		Label: "Add note (optional)",
		AllowEdit: true,
		Validate:  func(s string) error { return nil },
	}
	note, _ := notePrompt.Run()

	// Store current date (not time)
	today := time.Now().Format("2006-01-02")

	// Insert time entry
	_, err = db.Exec("INSERT INTO times (activity_id, date, duration, note) VALUES (?, ?, ?, ?)",
		activityID, today, durationMinutes, strings.TrimSpace(note))
	if err != nil {
		log.Fatalf("Failed to add time entry: %v", err)
	}

	fmt.Println("Time logged successfully.")
}