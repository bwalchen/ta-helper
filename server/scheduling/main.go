package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alabama/final-project-alabama/server/scheduling/handlers"
	"github.com/alabama/final-project-alabama/server/scheduling/questions"
	"github.com/go-redis/redis"
)

type ServiceEvent struct {
	ServiceName   string    `json:"name"`
	PathPattern   string    `json:"pathPattern"`
	Address       string    `json:"address"`
	LastHeartbeat time.Time `json:"lastHeartbeat"`
	Priviledged   bool      `json:"priviledged"`
}

//main is the main entry point for the server
func main() {
	addr := os.Getenv("ADDR")
	redisAddr := os.Getenv("REDISADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	redisdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for range ticker.C {
			event := &ServiceEvent{"scheduling", "/v1/scheduling", "scheduling:80", time.Now(), true}
			jsonString, err := json.Marshal(event)
			if err != nil {
				log.Fatal(err)
			}
			_, err = redisdb.RPush("ServiceEvents", jsonString).Result()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	mongoDBName := "tahelper"

	fmt.Println("Beginning...")
	MongoConnection, err := questions.NewSession("localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connecto to Mongo DB: %v \n", err)
	}
	fmt.Println("Successfully connected to Mongo!")

	// Context
	// ctx := models.Context{MongoConnection}
	// get users collection

	questionCollection := questions.QuestionCollection{MongoConnection.GetCollection(mongoDBName, "questions")}
	officeHoursCollection := questions.OfficeHourCollection{MongoConnection.GetCollection(mongoDBName, "officeHours")}

	ctx := handlers.Context{
		QuestionCollection:   questionCollection,
		OfficeHourCollection: officeHoursCollection,
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/scheduling", handlers.EnsureAuth(ctx.QuestionHandler))
	log.Printf("server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
