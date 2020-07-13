# MongoDB dengan Go

Bagian ini membahas tata cara penggunaan MongoDB mengunakan mongo-go-driver.

## Daftar Konten




## Tata cara pemasangan dan penggunaan 
Ada anjuran untuk driver resmi sebagainya menjadi bagian dari komponen modul Go,
maka untuk menginstall ini diharapkan berada di dalam working dir modul Go.

 1. Jalankan pada terminal.
 
 `
$ go get go.mongodb.org/mongo-driver/mongo
 `
 
 2. Kemudian import modul.
 ```go
 package main
 
 import (
    "fmt"
    "os"
    "go.mongodb.org/mongo-driver/mongo"
 )
 
 func main(){}
 ```

3. Buat fungsi untuk menghubungkan mongo-go-driver dengan server MongoDB.
```go

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

```

4. Panggil di entry function.

```go
func main(){
  InitMongo()
  selectDb(os.Getenv("DB_NAME"))
}
```

5. Eksekusi dengan perintah.

`
$ MONGODB_URL=mongodb://<alamat server mongo>:27017 DB_NAME=<nama database> go run main.go
`

