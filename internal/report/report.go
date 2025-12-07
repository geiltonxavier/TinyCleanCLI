package report

import (
	"fmt"
	"time"
)

// Item captures a cleanup candidate for reporting.
type Item struct {
	Category string
	Path     string
	Reason   string
}

// PrintResults renders a human-friendly summary.
func PrintResults(items []Item, dryRun bool, generatedAt time.Time) {
	header := "Planned cleanup (dry run)"
	if !dryRun {
		header = "Planned cleanup"
	}
	fmt.Println(header)
	fmt.Printf("Generated at: %s\n", generatedAt.Format(time.RFC1123))
	fmt.Println()

	if len(items) == 0 {
		fmt.Println("No candidates found yet.")
		fmt.Println("Note: logic is placeholder-only; implement scanners to see real results.")
		return
	}

	for _, item := range items {
		fmt.Printf("[%s] %s\n", item.Category, item.Path)
		if item.Reason != "" {
			fmt.Printf("  - %s\n", item.Reason)
		}
	}

	fmt.Println()
	fmt.Printf("Total candidates: %d\n", len(items))
	fmt.Println("Nothing was deleted. Implement deletion logic after reviewing the list.")
}
