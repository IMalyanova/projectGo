package main

import "fmt"

func main() {

	MemcachedAddresses := []string {"127.0.0.1:11211"}
	memcacheClient := memcache.New(MemcachedAddresses...)

	mkey := "coursera"

	memcacheClient.Set(&memcache.Item {
		Key:         mkey,
		Value:      []byte("1"),
		Expiration: 3,
	})

	memcacheClient.Increment("habrTag", 1)

	item, err := memcacheClient.Get(mkey)

	if err != nil && err != memcache.ErrCacheMiss {
		fmt.Println("MC error", err)
	}

	fmt.Printf("mc value %#v\n", item)
}



