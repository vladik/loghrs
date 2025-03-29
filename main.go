package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"loghrs/cmd"
	"loghrs/db"
)

func main() {
	// Determine DB path
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, ".loghrs", "loghrs.db")
	_ = os.MkdirAll(filepath.Dir(dbPath), 0755)

	// Open DB
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Ensure schema
	if err := db.Init(conn); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Handle command
	if len(os.Args) < 2 {
		fmt.Println("Usage: loghrs <command>")
		fmt.Println("Commands: add-activity")
		return
	}

	switch os.Args[1] {
	case "add-shelf":
		cmd.AddShelf(conn)	
	case "add-journal":
		cmd.AddJournal(conn)
	case "add-activity":
		cmd.AddActivity(conn)
	case "add-time":
		cmd.AddTime(conn)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}


