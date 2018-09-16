package models

import (
  "github.com/gocql/gocql"
  "time"
)

type Recipe struct {
  Id gocql.UUID `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
  CreatedAt time.Time `json:"created_at"`
}
