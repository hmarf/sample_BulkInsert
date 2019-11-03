package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type DBStruct struct {
	id       string
	name     string
	createAt time.Time
}

func BulkInsert(resc []DBStruct, db *sql.DB) {
	ins, err := db.Prepare("INSERT INTO sample(id,name,createdAt) VALUES(?,?,?)")
	if err != nil {
		log.Println(err)
	}
	for _, value := range resc {
		fmt.Println(value)
		ins.Exec(value.id, value.name, value.createAt)
		fmt.Println("ok")
	}
	return
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(0.0.0.0)/sampleDB?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u, err := uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
			return
		}
		u1 := u.String()

		u, err = uuid.NewRandom()
		if err != nil {
			fmt.Println(err)
			return
		}
		u2 := u.String()

		aa := []DBStruct{
			{
				id:       u1,
				name:     "name",
				createAt: time.Now(),
			},
			{
				id:       u2,
				name:     "name",
				createAt: time.Now(),
			},
		}

		go BulkInsert(aa, db)

	})

	// start server
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
