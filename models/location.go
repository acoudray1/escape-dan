package models

import (
    "net/http"
)

type Location struct {
    Id int `json:"id"`
    Address string `json:"address"`
    City string `json:"city"`
	PostalCode string `json:"postal_code"`
    PositionLat float64 `json:"position.x"`
    PositionLong float64 `json:"position.y"`
	IsAPlace bool `json:"is_a_place"`
	LastRecordDate Timestamp `json:"last_record_date"`
}

type LocationsList struct {
    Locations []Location `"json:locations"`
}

func (i *Location) Bind(r *http.Request) error {
    // if i.Name == "" {
    //     return fmt.Errorf("name is a required field")
    // }
    return nil
}

func (*LocationsList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (*Location) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}