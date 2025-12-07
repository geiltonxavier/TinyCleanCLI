package projects

import (
	"fmt"
	"os"
	"path/filepath"
)

// ProjectCandidate describes a project directory that might be removed.
type ProjectCandidate struct {
	Path   string
	Reason string
}

// FindInactiveProjects walks provided roots and produces placeholder suggestions.
// Future work: inspect git remotes, modification times, and prompt for deletion.
func FindInactiveProjects(roots []string, inactiveDays int) []ProjectCandidate {
	var candidates []ProjectCandidate

	for _, root := range roots {
		info, err := os.Stat(root)
		if err != nil || !info.IsDir() {
			continue
		}

		entries, err := os.ReadDir(root)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}

			projectPath := filepath.Join(root, entry.Name())
			if isGitRepo(projectPath) {
				candidates = append(candidates, ProjectCandidate{
					Path:   projectPath,
					Reason: fmt.Sprintf("placeholder: git repo; check clean + unused for %d days", inactiveDays),
				})
				continue
			}

			candidates = append(candidates, ProjectCandidate{
				Path:   projectPath,
				Reason: fmt.Sprintf("placeholder: non-git; check last modified > %d days", inactiveDays),
			})
		}
	}

	return candidates
}

func isGitRepo(path string) bool {
	gitPath := filepath.Join(path, ".git")
	info, err := os.Stat(gitPath)
	if err != nil {
		return false
	}
	return info.IsDir()
}
