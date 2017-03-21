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

		tx, err := db.Begin()
		if err != nil {
			fmt.Println(err)
		}
		stmt, err := tx.Prepare("insert into posts(title, text) values(?, ?)")
		if err != nil {
			fmt.Println(err)
		}
		defer stmt.Close()
		clean1 := strings.Replace(os.Args[2], "\"", "", -1)
		split1 := strings.Split(clean1, ",")
		title := strings.Replace(split1[0], "{title:", "", -1)
		clean2 := strings.Replace(split1[1], "text:", "", -1)
		text := strings.Replace(clean2, "}", "", -1)
		_, err = stmt.Exec(title, text)
		if err != nil {
			fmt.Println(err)
		}
		tx.Commit()
		fmt.Println("Post added")
	}

}
