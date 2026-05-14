package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fogibow/guardpatch/internal/analyzer"
	"github.com/fogibow/guardpatch/pkg/types"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "analyze":
		runAnalyze(os.Args[2:])
	default:
		printHelp()
		os.Exit(1)
	}
}

func runAnalyze(args []string) {
	fs := flag.NewFlagSet("analyze", flag.ExitOnError)
	target := fs.String("target", ".", "Target source directory")
	_ = fs.Parse(args)

	findings := analyzeTarget(*target)
	printFindings(findings)
}

func analyzeTarget(target string) []types.Finding {
	var all []types.Finding

	py := analyzer.NewPythonAnalyzer()
	pyFindings, err := py.Analyze(target)
	if err == nil {
		all = append(all, pyFindings...)
	}

	c := analyzer.NewCAnalyzer()
	cFindings, err := c.Analyze(target)
	if err == nil {
		all = append(all, cFindings...)
	}

	return all
}

func printFindings(findings []types.Finding) {
	if len(findings) == 0 {
		fmt.Println("No findings detected.")
		return
	}

	fmt.Printf("Detected %d finding(s):\n\n", len(findings))

	for _, f := range findings {
		fmt.Printf("- %s\n", f.Pattern)
		fmt.Printf("  File: %s:%d\n", f.File, f.Line)
		fmt.Printf("  Severity: %s\n", f.Severity)
		fmt.Printf("  Code: %s\n\n", f.Snippet)
	}
}

func printHelp() {
	fmt.Println("GuardPatch: Defensive AI Patch Assistant")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  guardpatch analyze --target ./path")
}
