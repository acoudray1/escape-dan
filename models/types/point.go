package models

import (
    "fmt"
    "net/http"
)

type Point struct {
    X float64 `json:"position.x"`
    Y float64 `json:"position.y"`
}