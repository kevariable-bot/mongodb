# Memulai belajar MongoDB

## Petunjuk

  - ## Instalasi

    - ### [Linux](Installation.md#linux)
    - ### [Windows](Installation.md#windows)

  - ## Intregrasi MongoDB terhadap Bahasa Pemograman

    - ### [Golang](go_example/README.md)

      - #### [CRUD Golang](contoh/crud)

  - ## DDL (Data Definition Language)

    - [x] [Collection](#collection)

## collection

collection merupakan tahap awal dalam pembuatan sebuah database, caranya hanya dengan perintah

  ```
  use <namaDatabase>
  ```

  > Gunakan camelCase, setiap akhir perintah boleh menggunakan semicolon ataupun tidak

## getCollection

getCollection merupakan sebuah sub collection yaitu sebuah table jika anda pernah menggunakan SQL

  ```
  db.getCollection(<namaCollection>)
  ```

  contoh disini kita akan membuat collection bernama `buku`

  ```
  db.getCollection('buku')
  ```

  > note: pastikan anda sudah `use <namaDatabase>` sebelumnya agar bisa membuat seperti diatas

## insertOne

Mengisi sebuah data terhada collection. Disini jika anda familiar menggunakan JavaScript sudah dipastikan anda akan terbiasa melakukan ini di MongoDB sudah berasa seperti dirumah sendiri ^^

  ```
  db.<namaCollection>.insertOne(<data>)
  ```

  `insertOne` hanya akan membuat satu data.

  ```
  db.buku.insertOne({
    id: 1,
    nama: "Belajar MongoDB maha asyieque"
  })
  ```

## insertMany

Mengisi sebuah data dengan banyak. like JSON

  ```
  db.buku.insertMany([
    {
      id: 1,
      nama: "Belajar sesuatu yang baru"
    },
    {
      id: 2,
      nama: "Belajar sesuatu yang lama"
    },
    ...data seterusnya
  ])
  ```

## updateOne

updateOne merupakan featur untuk merubah dan memanipulasi sebuah ini collection

saya asumsikan mempunyai data seperti ini

```
{ id: 1, nama: "Belajar MongoDB Pertama" },
{ id: 2, nama: "Belajar MongoDB Kedua" }, 
{ id: 3, nama: "Belajar MongoDB Ketiga" } 
```

pengaplikasian

```
db.<collection>.updateOne(<query>, <update>, [option])
```

```
db.buku.updateOne({ 
  nama: /Ketiga/ 
}, 
{ $set: {
   nama: "Belajar MongoDB Keempat" 
   }  
})
```

## updateMany

updateMany merupakan cara mengupdate sebuah data dengan jumlah yang multiple sekaligus dimana data yang sama akan diupdate semua

saya asumsikan data seperti ini

```
{ id: 1, nama: "Belajar MongoDB Pertama" },
{ id: 2, nama: "Belajar MongoDB Kedua" }, 
{ id: 3, nama: "Belajar MongoDB Ketiga" } 
```

contoh

```
db.buku.updateMany({
  nama: /Belajar/
},
{
  $set: {
    nama: "Semua kata Belajar akan terupdate"
  }
})
```

hasil

```
{ id: 1, nama: "Semua kata Belajar akan terupdate" },
{ id: 2, nama: "Semua kata Belajar akan terupdate" }, 
{ id: 3, nama: "Semua kata Belajar akan terupdate" } 
```

## deleteOne

Menghapus single isi data collection

```
db.<collection>.deleteOne(<query>)
```

contoh

```
db.buku.deleteOne({
  nama: /Belajar/
})
```

> akan menghapus karakter yang mirip dan hanya menghapus data yang pertama ia temui saja

## deleteMany

## find

Mencari sebuah data di collection dengan memanfaatkan fieldnya. Jika di SQL itu mirip `WHERE`

  ```
  db.<collection>.find([keyword])
  ```

  > Jika tidak menggunakan keyword maka data akan ditampilkan semua

  contoh disini saya asumsikan mempunyai data seperti ini ketika menuliskan db.buku.find()

  ```
  { "_id" : ObjectId("5ef0edc9794837e7c46cd8df"), "id" : 1, "nama" : "Fundamental MongoDB pertama" }
  { "_id" : ObjectId("5ef0edc9794837e7c46cd666"), "id" : 2, "nama" : "Fundamental MongoDB kedua" }
  ```

  jika ingin mencari id ke 2 maka

  ```
  db.buku.find({
    id: 2
  })
  ```

  atau

  ```
  db.buku.find({
    nama : "Fundamental MongoDB pertama"
  })
  ```

  namun setelah diperhatikan sangat lah susah jika 
  ingin mencari keyword `nama` lalu isinya panjang dan
  harus sama maka karena itu kita butuh menggunakan 
  Regular Expression jika di SQL itu seperti LIKE. Cukup mudah hanya perlu menambahkan / / seperti ini
  akan mencari datanya.

  ```
  db.buku.find({
    nama : /Mongo/
  })
  ```

  akan tampil semua karakter yang ada Mongo.

## aggregate

Aggregasi merupakan cara kita untuk melakukan operasi lebih dari satu dalam satu query, seperti find, sum dan semacamnya.
Dalam agregasi, setiap operasi dalam agregasi merupakan stage, dimana setiap hasil operasi merupakan stage yang akan diteruskan ke operasi selanjutnya (pipelining).
```
db.book.aggregate([
  { "$match": {"nama": "Laravel is Great"}},
  { "$lookup": 
    {
      from: "harga",
      localField: "id",
      foreignField: "id",
      as: "result"
    }
  },
  {"$project": 
     {
     "nama": "$nama",
     "harga": "$result.harga"
     }
   }
])
```

#### Alur pipeline data
Operator $match sama dengan db.find() / SELECT di SQL (stage 1)

| |  Hasil $match diteruskan ke $lookup sebagai input stage 2 ( stage 1 piping ke stage 2)

v v

Operator $lookup sama dengan LEFT-JOIN di SQL (stage 2)

| |  Hasil $lookup diterukasan ke $project sebagai input stage 3 (stage 2 piping ke stage 3)

v v

Operator $project sama dengan PROJECT di SQL (stage 3) <-- hasil akhir