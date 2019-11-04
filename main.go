package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
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
	rescInterface := []interface{}{}
	stmt := "INSERT INTO sample(id, name, createdAt) VALUES"

	for _, value := range resc {
		stmt += "(?,?,?),"
		rescInterface = append(rescInterface, value.id)
		rescInterface = append(rescInterface, value.name)
		rescInterface = append(rescInterface, value.createAt)
	}

	stmt = strings.TrimRight(stmt, ",")

	_, err := db.Exec(stmt, rescInterface...)
	fmt.Println(err)
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
