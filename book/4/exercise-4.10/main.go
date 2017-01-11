package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//!+
func main() {
	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	oneMonth := now.AddDate(0, -1, 0)
	oneYear := now.AddDate(-1, 0, 0)

	fmt.Printf("A month ago:\n")
	for _, item := range result.Items {
		if item.CreatedAt.Before(oneMonth) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Printf("A year ago:\n")
	for _, item := range result.Items {
		if item.CreatedAt.Before(oneYear) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Printf("After a year :\n")
	for _, item := range result.Items {
		if item.CreatedAt.After(oneYear) {
			fmt.Printf("#%-5d %9.9s %.55s\n",
				item.Number, item.User.Login, item.Title)
		}
	}
}

func filterTime(issue Issue, timeToMatch time.Time, timeFilter func(current, filteredTime time.Time) bool) bool {
	time := issue.CreatedAt
	return timeFilter(time, timeToMatch)
}
