package models

type Country struct {
	ID            int    `json:"id" bson:"id"`
	Code          string `json:"code" bson:"code"`
	Name          string `json:"name" bson:"name"`
	Continent     string `json:"continent" bson:"continent"`
	WikipediaLink string `json:"wikipedia_link" bson:"wikipedia_link"`
	Keywords      string `json:"keywords" bson:"keywords"`
}
