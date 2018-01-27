package main

import (
	"github.com/go-redis/redis"
	"github.com/mikedata/go-data-source-monitor/api"
	"github.com/mikedata/go-data-source-monitor/config"
	"github.com/mikedata/go-data-source-monitor/cron"
	"github.com/mikedata/go-data-source-monitor/mongo"

	"log"
)

func main() {

	// ---- Config ------
	cfg, err := config.Get()
	if err != nil {
		log.Fatal("Failed to get config.")
	}

	// ---- Mongo ------
	mongo := &mongo.Mongo{
		Collection: cfg.MongoConfig.Collection,
		Database:   cfg.MongoConfig.Database,
		APIURL:     cfg.APIURL,
		URI:        cfg.MongoConfig.Addr,
	}

	session, err := mongo.Init()
	if err != nil {
		log.Fatal("Failed to initialise mongo")
	}

	mongo.Session = session

	log.Printf("Listening to Mongo on: %s", mongo.URI)

	// ---- Redis ------

	rInt := int(cfg.RedisConfig.DB) // need int not int8

	redis := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Addr,
		Password: cfg.RedisConfig.Password,
		DB:       rInt,
	})

	_, err = redis.Ping().Result()
	if err != nil {
		log.Fatal("Initial Redis Connection Failed", err)
	}

	// ---- Start Task Monitoring ---
	log.Printf("Beginning Task Monitoring")
	go cron.Start(mongo, redis)

	// ---- Create and start API ----
	api.CreateAPI(cfg.APIURL, cfg.BindAddr, *mongo)

}
