package cleaner

import (
	"os"
	"path/filepath"
)

// Target describes a cache/log directory that could be cleared.
type Target struct {
	Path   string
	Reason string
}

// CacheTargets returns common macOS cache-like locations and their immediate contents.
func CacheTargets() []Target {
	rawPaths := []string{
		"~/Library/Caches",
		"/Library/Caches",
		"~/Library/Logs",
		"~/Library/Application Support",
		"~/Library/Developer/Xcode/DerivedData",
		"~/.Trash",
	}

	var targets []Target
	for _, p := range rawPaths {
		expanded := expandUser(p)
		info, err := os.Stat(expanded)
		if err != nil || !info.IsDir() {
			continue
		}

		entries, err := os.ReadDir(expanded)
		if err != nil {
			continue
		}

		if len(entries) == 0 {
			targets = append(targets, Target{
				Path:   expanded,
				Reason: "empty directory",
			})
			continue
		}

		for _, entry := range entries {
			targets = append(targets, Target{
				Path:   filepath.Join(expanded, entry.Name()),
				Reason: "cache/log cleanup",
			})
		}
	}

	return targets
}

func expandUser(p string) string {
	if len(p) > 0 && p[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return p
		}
		return filepath.Join(home, p[1:])
	}
	return p
}
