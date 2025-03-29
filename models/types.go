package models

type Shelf struct {
	ID        int
	Name      string
	CreatedAt string
}

type Journal struct {
	ID        int
	Name      string
	ShelfID   *int // optional
	CreatedAt string
}

type Activity struct {
	ID        int
	Name      string
	JournalID int
	CreatedAt string
}

type Time struct {
	ID         int
	ActivityID int
	Date       string
	Duration   int
	Note       string
	CreatedAt  string
}

type Label struct {
	ID        int
	Name      string
	CreatedAt string
}

type TimeLabel struct {
	TimeID  int
	LabelID int
}
