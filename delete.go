package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

func main() {
        //argsWithoutProg := os.Args[1:]
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Get configuration (path for sqlite db)
	db_path, err := client.Get("config:" + os.Args[1] + ":db_path").Result()
	clean1 := strings.Replace(os.Args[2], "{", "", -1)
	clean2 := strings.Replace(clean1, "}", "", -1)
	param := strings.Split(clean2, ":")
	if err == redis.Nil {
		fmt.Println("key2 does not exists", os.Args[1])
	} else if err != nil {
		panic(err)
	} else {
		db, err := sql.Open("sqlite3", db_path)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()

		stmt, err := db.Prepare("delete from posts where title = ?")
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(param[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Deleted")
	}

}
