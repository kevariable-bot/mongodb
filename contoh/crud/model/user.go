package model

import (
	"crud-mongo/config"
	"context"
)

type User struct {
	NamaDepan	string
	NamaBelakang	string
	NamaPengguna	string
}


/*

Struct ini akan dikonversi dari struct
menjadi BSON.

Representasi dalam BSON

	{
		"NamaDepan": "",
		"NamaBelakang": "",
		"NamaPengguna": ""
	}
*/

func(u User) Insert()(*InsertOneResult,error){
	return config.dbInstance.collection("users").InsertOne(context.TODO(),u)
	// config.dbInstance dijabarkan untuk mudah dikenali
	// gunakan abstraksi lanjut untuk DX yang lebih baik
}

func(u User) Read(namaPengguna string)([]*User,error){
	hasil,err := config.dbInstance.collection("users").Find(
		context.TODO(),
		bson.D{{"NamaPengguna": namaPengguna}},
	)
	if err != nil {
		return nil,err
	}

	var hasil []*User

	for _,elemen := range hasil{
		user := new(User)
		err = elemen.Decode(&user)
		if err != nil{
			return nil,err
		}
		hasil = append(hasil,&user)
	}
	return hasil,nil
}

// WIP
