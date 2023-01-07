package main

import "github.com/jayanthkrishna/distributed-cache-golang/cache"

func main() {

	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}

	server := NewServer(opts, cache.New())
}
