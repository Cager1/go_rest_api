package models

type Book struct {
	ID       int       `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Author   string    `json:"author,omitempty"`
	Sciences []Science `json:"sciences,omitempty"`
}
