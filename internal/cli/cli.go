package cli

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/geiltonxavier/TinyCleanCLI/internal/apps"
	"github.com/geiltonxavier/TinyCleanCLI/internal/cleaner"
	"github.com/geiltonxavier/TinyCleanCLI/internal/projects"
	"github.com/geiltonxavier/TinyCleanCLI/internal/report"
)

// Options controls the scan behavior.
type Options struct {
	DryRun          bool
	IncludeApps     bool
	IncludeProjects bool
	IncludeCaches   bool
	InactiveDays    int
	ProjectPaths    []string
}

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSlice) Set(value string) error {
	if value == "" {
		return errors.New("empty path")
	}
	*s = append(*s, value)
	return nil
}

// Execute is the public entrypoint for the CLI.
func Execute(args []string) error {
	if len(args) < 2 {
		printRootUsage()
		return nil
	}

	switch args[1] {
	case "scan":
		return runScan(args[2:])
	case "-h", "--help", "help":
		printRootUsage()
		return nil
	default:
		printRootUsage()
		return fmt.Errorf("unknown command %q", args[1])
	}
}

func printRootUsage() {
	fmt.Println("TinyCleanCLI - macOS cleanup assistant")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  tinycleancli scan [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  scan    Inspect caches, apps, and projects for cleanup candidates")
	fmt.Println()
	fmt.Println("Use \"tinycleancli scan -h\" for flags.")
}

func runScan(args []string) error {
	var opts Options
	var projectRoots stringSlice

	// Default search roots for projects can be overridden.
	projectRoots = append(projectRoots,
		filepath.Join(os.Getenv("HOME"), "Projects"),
		filepath.Join(os.Getenv("HOME"), "projects"),
		filepath.Join(os.Getenv("HOME"), "code"),
	)

	fs := flag.NewFlagSet("scan", flag.ContinueOnError)
	fs.BoolVar(&opts.DryRun, "dry-run", false, "simulate cleanup actions without deleting anything")
	fs.BoolVar(&opts.IncludeApps, "apps", true, "include unused app detection")
	fs.BoolVar(&opts.IncludeProjects, "projects", true, "include inactive projects detection")
	fs.BoolVar(&opts.IncludeCaches, "caches", true, "include system cache and logs cleanup")
	fs.IntVar(&opts.InactiveDays, "days", 30, "consider items inactive after this many days")
	fs.Var(&projectRoots, "projects-path", "path to search for projects (repeatable)")
	fs.SetOutput(os.Stdout)

	if err := fs.Parse(args); err != nil {
		return err
	}

	opts.ProjectPaths = dedupeStrings(projectRoots)

	var results []report.Item

	if opts.IncludeCaches {
		for _, c := range cleaner.CacheTargets() {
			results = append(results, report.Item{
				Category: "cache",
				Path:     c.Path,
				Reason:   c.Reason,
			})
		}
	}

	if opts.IncludeApps {
		for _, a := range apps.FindStaleApps(opts.InactiveDays) {
			results = append(results, report.Item{
				Category: "app",
				Path:     a.Path,
				Reason:   a.Reason,
			})
		}
	}

	if opts.IncludeProjects {
		for _, p := range projects.FindInactiveProjects(opts.ProjectPaths, opts.InactiveDays) {
			results = append(results, report.Item{
				Category: "project",
				Path:     p.Path,
				Reason:   p.Reason,
			})
		}
	}

	report.PrintResults(results, opts.DryRun, time.Now())
	return nil
}

func dedupeStrings(values []string) []string {
	seen := make(map[string]struct{})
	var cleaned []string
	for _, v := range values {
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			continue
		}
		if _, ok := seen[trimmed]; ok {
			continue
		}
		seen[trimmed] = struct{}{}
		cleaned = append(cleaned, trimmed)
	}
	return cleaned
}
