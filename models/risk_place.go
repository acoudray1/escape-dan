package models

import (
    "net/http"
)

type RiskPlace struct {
    Id int `json:"id"`
    LocationId int `json:"location_id"`
    NuberOfRecords int `json:"number_of_records"`
}

type RiskPlacesList struct {
    RiskPlaces []RiskPlace `"json:risk_places"`
}

func (i *RiskPlace) Bind(r *http.Request) error {
    // if i.Name == "" {
    //     return fmt.Errorf("name is a required field")
    // }
    return nil
}

func (*RiskPlacesList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (*RiskPlace) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}