package analyzer

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fogibow/guardpatch/pkg/types"
)

type CAnalyzer struct{}

func NewCAnalyzer() *CAnalyzer {
	return &CAnalyzer{}
}

func (a *CAnalyzer) Analyze(path string) ([]types.Finding, error) {
	var findings []types.Finding

	err := filepath.Walk(path, func(file string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		if !(strings.HasSuffix(file, ".c") || strings.HasSuffix(file, ".h")) {
			return nil
		}

		content, readErr := os.ReadFile(file)
		if readErr != nil {
			return nil
		}

		lines := strings.Split(string(content), "\n")

		patterns := map[string]*regexp.Regexp{
			"c-unsafe-strcpy":  regexp.MustCompile(`\bstrcpy\s*\(`),
			"c-unsafe-strcat":  regexp.MustCompile(`\bstrcat\s*\(`),
			"c-unsafe-gets":    regexp.MustCompile(`\bgets\s*\(`),
			"c-unsafe-sprintf": regexp.MustCompile(`\bsprintf\s*\(`),
		}

		for i, line := range lines {
			for id, re := range patterns {
				if re.MatchString(line) {
					findings = append(findings, types.Finding{
						ID:          id,
						Language:    "c",
						File:        file,
						Line:        i + 1,
						Pattern:     id,
						Snippet:     strings.TrimSpace(line),
						Severity:    "high",
						Description: "Potential unsafe C function detected.",
					})
				}
			}
		}

		return nil
	})

	return findings, err
}
