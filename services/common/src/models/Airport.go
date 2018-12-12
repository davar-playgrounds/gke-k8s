package models

type Airport struct {
	ID               int    `json:"id,omitempty"`
	Ident            string `json:"ident,omitempty"`
	Type             string `json:"type,omitempty"`
	Name             string `json:"name,omitempty"`
	LatitudeDeg      string `json:"latitude_deg,omitempty"`
	LongitudeDeg     string `json:"longitude_deg,omitempty"`
	ElevationFt      string `json:"elevation_ft,omitempty"`
	Continent        string `json:"continent,omitempty"`
	IsoCountry       string `json:"iso_country,omitempty"`
	IsoRegion        string `json:"iso_region,omitempty"`
	Municipality     string `json:"municipality,omitempty"`
	ScheduledService string `json:"scheduled_service,omitempty"`
	GpsCode          string `json:"gps_code,omitempty"`
	IataCode         string `json:"iata_code,omitempty"`
	LocalCode        string `json:"local_code,omitempty"`
	HomeLink         string `json:"home_link,omitempty"`
	WikipediaLink    string `json:"wikipedia_link,omitempty"`
	Keywords         string `json:"keywords,omitempty"`
}
