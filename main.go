package main

import (
  "github.com/julienschmidt/httprouter"
  "github.com/gocql/gocql"
  "log"
  "net/http"

  "recipes/controllers"
  "recipes/cassandra"
)

func main() {
  cluster := gocql.NewCluster("127.0.0.1")
  cluster.Keyspace = "recipes_api"
  cassandra.Session, _ = cluster.CreateSession()
  defer cassandra.Session.Close()

  router := httprouter.New()
  router.GET("/", controllers.Index)
  log.Fatal(http.ListenAndServe(":8080", router))
}
