package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// get last argument (the interstellar argument json)
	last := len(os.Args) - 1
	// Get and decode the json body passed by arguments
	byt := []byte(os.Args[last])
	var dat map[string]map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Get configuration (path for sqlite db)
	db_path, err := client.Get("config:" + dat["headers"]["host"].(string) + ":db_path").Result()
	param := dat["querystring"]["title"]
	if err == redis.Nil {
		fmt.Println("key2 does not exists", dat["headers"]["host"].(string))
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
		_, err = stmt.Exec(param)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Deleted")
	}

}
