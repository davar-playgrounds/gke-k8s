package models

type Runway struct {
	ID                     int     `json:"id,omitempty"`
	Ref                    int     `json:"airport_ref,omitempty"`
	Ident                  string  `json:"airport_ident,omitempty"`
	LengthFt               int     `json:"length_ft,omitempty"`
	WidthFt                int     `json:"width_ft,omitempty"`
	Surface                string  `json:"surface,omitempty"`
	Lighted                bool    `json:"lighted,omitempty"`
	Closed                 bool    `json:"closed,omitempty"`
	LeIdent                string  `json:"le_ident,omitempty"`
	LeLatitudeDeg          float32 `json:"le_latitude_deg,omitempty"`
	LeLongitudeDeg         float32 `json:"le_longitude_deg,omitempty"`
	LeElevationFt          int     `json:"le_elevation_ft,omitempty"`
	LeHeadingDegT          int     `json:"le_heading_degT,omitempty"`
	LeDisplacedThresholdFt int     `json:"le_displaced_threshold_ft,omitempty"`
	HeIdent                string  `json:"he_ident,omitempty"`
	HeLatitudeDeg          float32 `json:"he_latitude_deg,omitempty"`
	HeLongitudeDeg         float32 `json:"he_longitude_deg,omitempty"`
	HeElevationFt          int     `json:"he_elevation_ft,omitempty"`
	HeHeadingDegT          int     `json:"he_heading_degT,omitempty"`
	HeDisplacedThresholdFt int     `json:"he_displaced_threshold_ft,omitempty"`
}
