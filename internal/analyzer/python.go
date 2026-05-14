package analyzer

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fogibow/guardpatch/pkg/types"
)

type PythonAnalyzer struct{}

func NewPythonAnalyzer() *PythonAnalyzer {
	return &PythonAnalyzer{}
}

func (a *PythonAnalyzer) Analyze(path string) ([]types.Finding, error) {
	var findings []types.Finding

	err := filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(file, ".py") {
			return nil
		}

		content, readErr := os.ReadFile(file)
		if readErr != nil {
			return nil
		}

		lines := strings.Split(string(content), "\n")

		patterns := map[string]*regexp.Regexp{
			"python-command-injection-shell-true": regexp.MustCompile(`subprocess\.(run|call|Popen)\(.*shell\s*=\s*True`),
			"python-unsafe-eval":                 regexp.MustCompile(`\beval\s*\(`),
			"python-unsafe-exec":                 regexp.MustCompile(`\bexec\s*\(`),
			"python-unsafe-pickle":               regexp.MustCompile(`pickle\.loads?\s*\(`),
		}

		for i, line := range lines {
			for id, re := range patterns {
				if re.MatchString(line) {
					findings = append(findings, types.Finding{
						ID:          id,
						Language:    "python",
						File:        file,
						Line:        i + 1,
						Pattern:     id,
						Snippet:     strings.TrimSpace(line),
						Severity:    "high",
						Description: "Potential insecure Python pattern detected.",
					})
				}
			}
		}

		return nil
	})

	return findings, err
}
