package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/thiagodebastos/gofixit/domain/entity"
	"github.com/thiagodebastos/gofixit/domain/valueobject"
	"github.com/thiagodebastos/gofixit/infra/persistence/sqlite"
)

func createIssue() entity.Issue {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Create a new issue title: ")

	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Add a description: ")

	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Add a priority [lowest, low, medium, high, highest]: ")
	priority, _ := reader.ReadString('\n')
	priority = strings.TrimSpace(priority)
	newPriority, priorityOk := valueobject.PriorityFromString(priority)

	for !priorityOk {
		fmt.Print("Invalid priority, add a priority [lowest, low, medium, high, highest]: ")
		priority, _ := reader.ReadString('\n')
		priority = strings.TrimSpace(priority)
		newPriority, priorityOk = valueobject.PriorityFromString(priority)
	}

	i, _ := entity.CreateIssue(
		uuid.New(),
		title,
		description,
		valueobject.StatusOpen,
		newPriority,
	)

	return i
}

func main() {
	// Open a connection to an in-memory SQLite database.
	conn, err := sqlite.OpenConn(":memory:", sqlite.OpenReadWrite)
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}
	defer conn.Close()

	// Set up your repositories
	userRepo := sqlite.NewUserRepository(conn)
	issueRepo := sqlite.NewIssueRepository(conn)

	// Example of using the repositories
	err = userRepo.CreateUser()
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

	user, err := userRepo.GetUser(1)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}
	log.Printf("User: %v", user)

	printAligned := func(label string, value interface{}) {
		fmt.Printf("%-12s: %v\n", label, value)
	}

	printIssue := func(i entity.Issue) {
		printAligned("ID", i.ID())
		printAligned("Title", i.Title())
		printAligned("Description", i.Description())
		printAligned("Status", i.Status().ToString())
		printAligned("Priority", i.Priority().String())
		fmt.Printf("\n")
	}

	myIssue := createIssue()

	printIssue(myIssue)
}
