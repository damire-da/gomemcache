package main

import (
	gomemcache "cache/in-memory/cache"
	"fmt"
	"time"
)

func main() {
	cache := gomemcache.NewMemCache()
	num := 1

	cache.Set("test", num, time.Minute)

	fmt.Println(cache.Get("test"))

	cache.StartCleanup(time.Second * 10)

	// Периодически проверяем наличие элементов в кэше
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		_, found := cache.Get("test")
		if found {
			fmt.Println("Key1 still exists")
		} else {
			fmt.Println("Key1 removed")
		}
	}
}