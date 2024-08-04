package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/entity"
)

func main() {
	printAligned := func(label string, value interface{}) {
		fmt.Printf("%-12s: %v\n", label, value)
	}

	printIssue := func(i entity.Issue) {
		printAligned("ID", i.ID())
		printAligned("Title", i.Title)
		printAligned("Description", i.Description)
		printAligned("Status", i.Status)
		printAligned("CreatedAt", i.CreatedAt)
		printAligned("UpdatedAt", i.UpdatedAt)
		fmt.Printf("\n")
	}

	myIssue := entity.NewIssue(
		uuid.New(),
		"make wireframes",
		"Create UI wireframes based on UX design spec",
	)

	myIssue.UpdateStatus(entity.StatusDoing)

	printIssue(myIssue)

	myIssue.UpdateStatus(entity.StatusDone)

	printIssue(myIssue)
}
