package controllers

import (
  "net/http"
  "encoding/json"
  "github.com/julienschmidt/httprouter"
  "github.com/gocql/gocql"

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
    recipe.Id = m["id"].(gocql.UUID)
    recipe.Name = m["name"].(string)
    recipe.Description = m["description"].(string)

    recipes = append(recipes, recipe)
    m = map[string]interface{}{}
  }

  response, _ := json.Marshal(recipes)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(200)
  w.Write(response)
}
