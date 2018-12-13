package models

type Airport struct {
	ID               int    `json:"id" bson:"id"`
	Ident            string `json:"ident" bson:"ident"`
	Type             string `json:"type" bson:"type"`
	Name             string `json:"name" bson:"name"`
	LatitudeDeg      string `json:"latitude_deg" bson:"latitude_deg"`
	LongitudeDeg     string `json:"longitude_deg" bson:"longitude_deg"`
	ElevationFt      string `json:"elevation_ft" bson:"elevation_ft"`
	Continent        string `json:"continent" bson:"continent"`
	IsoCountry       string `json:"iso_country" bson:"iso_country"`
	IsoRegion        string `json:"iso_region" bson:"iso_region"`
	Municipality     string `json:"municipality" bson:"municipality"`
	ScheduledService string `json:"scheduled_service" bson:"scheduled_service"`
	GpsCode          string `json:"gps_code" bson:"gps_code"`
	IataCode         string `json:"iata_code" bson:"iata_code"`
	LocalCode        string `json:"local_code" bson:"local_code"`
	HomeLink         string `json:"home_link" bson:"home_link"`
	WikipediaLink    string `json:"wikipedia_link" bson:"wikipedia_link"`
	Keywords         string `json:"keywords" bson:"keywords"`
}
