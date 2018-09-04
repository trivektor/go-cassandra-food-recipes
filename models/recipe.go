package models

import (
  "github.com/gocql/gocql"
)

type Recipe struct {
  Id gocql.UUID `json:"id"`
  Name string `json:"name"`
  Description string `json:"description"`
}
