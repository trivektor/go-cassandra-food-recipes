package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "github.com/gocql/gocql"
  "time"
  "fmt"

  "recipes/models"
  "recipes/cassandra"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  recipes := []models.Recipe{}
  recipe := models.Recipe{}
  m := map[string]interface{}{}

  query := "SELECT * FROM recipes LIMIT 10"
  iterable := cassandra.Session.Query(query).Iter()

  for iterable.MapScan(m) {
    recipe.Id = m["id"].(gocql.UUID) // gocql.UUID type assertion
    recipe.Name = m["name"].(string) // string type assertion
    recipe.Description = m["description"].(string)
    recipe.CreatedAt = m["created_at"].(time.Time)

    recipes = append(recipes, recipe)
    m = map[string]interface{}{}
  }

  response, _ := json.Marshal(recipes)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(response)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  name := r.FormValue("name")
  description := r.FormValue("description")

  err := cassandra.Session.Query(
    "INSERT INTO recipes(id, created_at, name, description) VALUES(now(), ?, ?, ?)",
    time.Now(), name, description,
  ).Exec()

  w.Header().Set("Content-Type", "application/json")

  fmt.Println(err)

  if err != nil {
    m := make(map[string]string)
    m["error"] = err.Error()
    response, _ := json.Marshal(m)
    w.WriteHeader(400)
    w.Write(response)
  } else {
    w.WriteHeader(201)
  }
}
