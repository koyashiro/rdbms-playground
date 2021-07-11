package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/docker/docker/client"
	"github.com/koyashiro/postgres-playground/backend/runtime"

	_ "github.com/lib/pq"
)

func main() {
	// set lib logger
	runtime.Logger = log.New(os.Stdout, "playground: ", log.LstdFlags)

	dbm := runtime.NewDBManage(client.DefaultDockerHost)

	fmt.Println("try create")

	db, err := dbm.Create("example")
	if err != nil {
		panic(err)
	}

	fmt.Println(db.ID, db.Hash, db.Status, db.Port)

	time.Sleep(2 * time.Second)

	dbd, err := sql.Open("postgres", "user=playground password=password dbname=playground sslmode=disable port="+strconv.Itoa(db.Port))
	if err != nil {
		panic(err)
	}

	fmt.Println("try query execute")

	rows, err := dbd.Query("SELECT TRUE AS Result;")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = rows.Close()
	}()

	var result string
	for rows.Next() {
		switch err = rows.Scan(&result); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%s)\n", result)
		default:
			panic(err)
		}
	}

	fmt.Println("try destroy")
	err = dbm.Destroy("example")

	if err != nil {
		panic(err)
	}

	// cleanup container
	// docker rm -f `docker ps -f "label=type=playground" -aq`
}
