package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/entity"
	"github.com/thiagodebastos/gofixit/domain/valueobject"
)

func createIssue() entity.Issue {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create a new issue title: ")

	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Add a description: ")

	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	i, _ := entity.NewIssue(uuid.New(), title, description, valueobject.StatusOpen)

	print(i)

	return nil
}

func main() {
	printAligned := func(label string, value interface{}) {
		fmt.Printf("%-12s: %v\n", label, value)
	}

	printIssue := func(i entity.Issue) {
		printAligned("ID", i.ID())
		printAligned("Title", i.Title())
		printAligned("Description", i.Description())
		printAligned("Status", i.Status())
		fmt.Printf("\n")
	}

	myIssue := createIssue()

	printIssue(myIssue)
}
