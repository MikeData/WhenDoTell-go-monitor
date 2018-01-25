package main

import (
  "github.com/mikedata/go-data-source-monitor/api"
  "github.com/mikedata/go-data-source-monitor/config"
  "github.com/mikedata/go-data-source-monitor/mongo"
  "github.com/mikedata/go-data-source-monitor/cron"

  "log"
)


func main() {

  // ---- Config & Mongo ----
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

  log.Printf("Listening to Mongo on: %s", mongo.URI)

  // ---- Start Task Monitoring ---
  log.Printf("Beginning Task Montioring")
  go cron.Start(mongo)

  // ---- Create and start API ----
  api.CreateAPI(cfg.APIURL, cfg.BindAddr, *mongo)

}
