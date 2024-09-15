package main

import (
	"cache/cacher"
	"fmt"
	"log"
	"time"
)

func main() {
	cache := cacher.New()

	cache.Set("userId", 42, time.Second*5)

	userId, err := cache.Get("userId")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)

	time.Sleep(time.Second * 6)

	userId, err = cache.Get("userId")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userId)
}
