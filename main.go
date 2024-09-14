package main

import (
	"cache/cacher"
	"fmt"
)

func main() {
	cache := cacher.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
