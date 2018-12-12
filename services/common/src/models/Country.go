package models

type Country struct {
	ID           int    `json:"id,omitempty"`
	Code         string `json:"code,omitempty"`
	Name         string `json:"name,omitempty"`
	Continent    string `json:"continent,omitempty"`
	WikipediaUri string `json:"wikipedia_link,omitempty"`
	Keywords     string `json:"keywords,omitempty"`
}