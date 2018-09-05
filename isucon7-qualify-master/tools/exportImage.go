package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Image image table entity
type Image struct {
	Id   string
	Name string
	Data []byte
}

func main() {
	db, err := sqlx.Connect("mysql", "isucon:isucon@(127.0.0.1:3306)/isubata")
	if err != nil {
		log.Fatalln(err)
	}
	url := "..//webapp/public/image/"
	image := []Image{}
	errr := db.Select(&image, "SELECT * FROM image")
	if errr != nil {
	}
	fmt.Print(len(image))
	for i := 0; i < len(image); i++ {
		file, _ := os.Create((url + image[i].Name)
		file.Write(image[i].Data)
	}
}
