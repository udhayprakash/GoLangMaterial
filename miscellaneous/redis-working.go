package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

/*
Purpose: Redis
	- in-memory data structure store
	- used as a database, cache, and message broker.
	- It supports data structures such as strings,
	  hashes, lists, sets, sorted sets with range queries,
	  bitmaps, hyperloglogs, geospatial indexes with
	  radius queries and streams.

	- https://redis.io/
	- go get github.com/gomodule/redigo/redis
*/

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatal(err) // handle error
	}
	defer c.Close()

	_, err = conn.Do(
		"HMSET",
		"book:1",
		"title",
		"Sherlock Holmes",
		"creator",
		"AC Doyle",
		"category",
		"detective",
		"price of book",
		59.99,
	)

	title, err := redis.String(conn.Do("HGET", "book:1", "title"))
	if err != nil {
		log.Fatal(err) // handle error
	}
	fmt.Println("Book Title:", title)
}
