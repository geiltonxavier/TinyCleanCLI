package report

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// Item captures a cleanup candidate for reporting.
type Item struct {
	Category string
	Path     string
	Reason   string
}

// Options controls reporting output and context.
type Options struct {
	DryRun          bool
	Verbose         bool
	InactiveDays    int
	IncludeApps     bool
	IncludeCaches   bool
	IncludeProjects bool
	GeneratedAt     time.Time
}

// PrintResults renders a human-friendly summary grouped by category.
func PrintResults(items []Item, opts Options) {
	mode := "plan"
	if opts.DryRun {
		mode = "dry-run"
	}

	fmt.Printf("TinyCleanCLI %s\n", strings.ToUpper(mode))
	fmt.Printf("Generated: %s\n", opts.GeneratedAt.Format(time.RFC1123))
	fmt.Printf("Mode: dry-run=%v | days=%d | apps=%v | projects=%v | caches=%v | verbose=%v\n",
		opts.DryRun, opts.InactiveDays, opts.IncludeApps, opts.IncludeProjects, opts.IncludeCaches, opts.Verbose)
	fmt.Println()

	if len(items) == 0 {
		fmt.Println("No candidates found. (Current scanners still use placeholder logic.)")
		return
	}

	grouped := groupByCategory(items)
	categoryOrder := orderedCategories(grouped)
	limit := 8
	if opts.Verbose {
		limit = -1
	}

	total := 0
	for _, cat := range categoryOrder {
		group := grouped[cat]
		total += len(group)
		label := prettyCategory(cat)
		fmt.Printf("%s (%d)\n", label, len(group))

		count := len(group)
		if limit > 0 && count > limit {
			count = limit
		}
		for i := 0; i < count; i++ {
			item := group[i]
			fmt.Printf("  • %s\n", item.Path)
			if item.Reason != "" {
				fmt.Printf("    - %s\n", item.Reason)
			}
		}
		if limit > 0 && len(group) > limit {
			fmt.Printf("  … %d more (use --verbose to see all)\n", len(group)-limit)
		}
		fmt.Println()
	}

	fmt.Printf("Total candidates: %d\n", total)
	if opts.DryRun {
		fmt.Println("Nothing was deleted because dry-run is enabled.")
	}
	fmt.Println("Current scanners are placeholders; refine logic before enabling deletion.")
}

func groupByCategory(items []Item) map[string][]Item {
	grouped := make(map[string][]Item)
	for _, item := range items {
		grouped[item.Category] = append(grouped[item.Category], item)
	}

	for _, group := range grouped {
		sort.Slice(group, func(i, j int) bool {
			return group[i].Path < group[j].Path
		})
	}

	return grouped
}

func orderedCategories(grouped map[string][]Item) []string {
	priority := []string{"cache", "app", "project"}
	var order []string
	seen := make(map[string]struct{})

	for _, p := range priority {
		if _, ok := grouped[p]; ok {
			order = append(order, p)
			seen[p] = struct{}{}
		}
	}

	var others []string
	for cat := range grouped {
		if _, ok := seen[cat]; ok {
			continue
		}
		others = append(others, cat)
	}
	sort.Strings(others)
	return append(order, others...)
}

func prettyCategory(cat string) string {
	switch strings.ToLower(cat) {
	case "cache":
		return "Caches / Logs / Support / Trash"
	case "app":
		return "Applications"
	case "project":
		return "Projects"
	default:
		return strings.Title(cat)
	}
}
