package main

import (
  "github.com/mikedata/go-data-source-monitor/api"
  "github.com/mikedata/go-data-source-monitor/config"
  "github.com/mikedata/go-data-source-monitor/mongo"

  "log"
)


func main() {

  cfg, err := config.Get()
	if err != nil {
		log.Fatal("Failed to get config.")
	}

  mongo := &mongo.Mongo{
    Collection:  cfg.MongoConfig.Collection,
    Database:    cfg.MongoConfig.Database,
    APIURL:      cfg.APIURL,
    URI:         cfg.MongoConfig.BindAddr,
  }

  session, err := mongo.Init()
  if err != nil {
    log.Fatal("Failed to initialise mongo")
  }

  mongo.Session = session

  log.Printf("Attempting to connect to Mongo on: %s", mongo.URI)

  api.CreateAPI(cfg.APIURL, cfg.BindAddr, *mongo)

}
