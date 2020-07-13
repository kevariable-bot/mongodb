package main

import (
   "fmt"
   "os"
   "crud-mongo/config"
   "crud-mongo/controller"
   "crud-mongo/model"
)

func main(){
  config.InitMongo()
  config.selectDb(os.Getenv("DB_NAME"))
}
