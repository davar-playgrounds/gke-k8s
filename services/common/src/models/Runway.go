package models

type Runway struct {
	ID                     int     `json:"id" bson:"id"`
	Ref                    int     `json:"airport_ref" bson:"airport_ref"`
	Ident                  string  `json:"airport_ident" bson:"airport_ident"`
	LengthFt               int     `json:"length_ft" bson:"length_ft"`
	WidthFt                int     `json:"width_ft" bson:"width_ft"`
	Surface                string  `json:"surface" bson:"surface"`
	Lighted                bool    `json:"lighted" bson:"lighted"`
	Closed                 bool    `json:"closed" bson:"closed"`
	LeIdent                string  `json:"le_ident" bson:"le_ident"`
	LeLatitudeDeg          float32 `json:"le_latitude_deg" bson:"le_latitude_deg"`
	LeLongitudeDeg         float32 `json:"le_longitude_deg" bson:"le_longitude_deg"`
	LeElevationFt          int     `json:"le_elevation_ft" bson:"le_elevation_ft"`
	LeHeadingDegT          int     `json:"le_heading_degT" bson:"le_heading_degT"`
	LeDisplacedThresholdFt int     `json:"le_displaced_threshold_ft" bson:"le_displaced_threshold_ft"`
	HeIdent                string  `json:"he_ident" bson:"he_ident"`
	HeLatitudeDeg          float32 `json:"he_latitude_deg" bson:"he_latitude_deg"`
	HeLongitudeDeg         float32 `json:"he_longitude_deg" bson:"he_longitude_deg"`
	HeElevationFt          int     `json:"he_elevation_ft" bson:"he_elevation_ft"`
	HeHeadingDegT          int     `json:"he_heading_degT" bson:"he_heading_degT"`
	HeDisplacedThresholdFt int     `json:"he_displaced_threshold_ft" bson:"he_displaced_threshold_ft"`
}
