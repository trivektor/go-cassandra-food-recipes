package main

import (
  "github.com/julienschmidt/httprouter"
  "log"
  "net/http"
  "recipes/controllers"
)

func main() {
  router := httprouter.New()
  router.GET("/", controllers.Index)
  log.Fatal(http.ListenAndServe(":8080", router))
}
