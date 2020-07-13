package config

import (
   "fmt"
   "os"
   "go.mongodb.org/mongo-driver/mongo"
)

var clientInstance *mongo.Client
var dbInstance *mongo.Database

func InitMongo(){
  clientInstance, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
  // os.Getenv("MONGODB_URL") bisa diganti string URL
  if err != nil {
		panic(err)
	}

	if err = clientInstance.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}
  
  fmt.Println("[*] Driver terhubung")
}

func selectDb(dbName string){
  dbInstance = clientInstance.Database(dbName)
}
