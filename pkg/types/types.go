package types

type Finding struct {
	ID          string `json:"id"`
	Language    string `json:"language"`
	File        string `json:"file"`
	Line        int    `json:"line"`
	Pattern     string `json:"pattern"`
	Snippet     string `json:"snippet"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
}
