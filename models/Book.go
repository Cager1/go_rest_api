package models

type Book struct {
	ID       string   `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Author   string   `json:"author,omitempty"`
	Sciences []string `json:"sciences,omitempty"`
}
