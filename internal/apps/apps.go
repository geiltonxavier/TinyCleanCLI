package apps

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// StaleApp represents an app candidate for removal.
type StaleApp struct {
	Path   string
	Reason string
}

// FindStaleApps is a placeholder that lists top-level apps in /Applications.
// Future work: read last-opened metadata to decide staleness.
func FindStaleApps(inactiveDays int) []StaleApp {
	appDir := "/Applications"
	entries, err := os.ReadDir(appDir)
	if err != nil {
		return nil
	}

	var apps []StaleApp
	for _, entry := range entries {
		if entry.IsDir() && strings.HasSuffix(entry.Name(), ".app") {
			apps = append(apps, StaleApp{
				Path:   filepath.Join(appDir, entry.Name()),
				Reason: fmt.Sprintf("placeholder: check if unused for %d days", inactiveDays),
			})
		}
	}
	return apps
}
