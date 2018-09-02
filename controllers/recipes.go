package controllers

import (
  "net/http"
  "github.com/julienschmidt/httprouter"
  "fmt"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprintf(w, "All Recipes")
}
