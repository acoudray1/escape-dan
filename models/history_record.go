package models

import (
    "net/http"
)

type HistoryRecord struct {
    UserId int `json:"user_id"`
    LocationId int `json:"location_id"`
    RecordDate Timestamp `json:"record_date"`
    DangerType Danger `json:"danger_type"`
    NumberOfValidations int `json:"nb_of_validations"`
}

type HistoryRecordsList struct {
    HistoryRecords []HistoryRecord `"json:history_records"`
}

func (i *HistoryRecord) Bind(r *http.Request) error {
    // if i.Name == "" {
    //     return fmt.Errorf("name is a required field")
    // }
    return nil
}

func (*HistoryRecordsList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (*HistoryRecord) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}