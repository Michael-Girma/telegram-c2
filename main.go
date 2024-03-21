package main

import (
	"log"
	"sync"
	"telegram-c2/internal/pkg/config"
	"telegram-c2/internal/pkg/services/c2"
)

func main() {
	var config = config.NewConfig()
	var server = c2.NewTGC2(config)
	var listeners sync.WaitGroup

	listeners.Add(1)
	go func() {
		defer listeners.Done()
		err := server.ListenForNewAgents()
		if err != nil {
			log.Fatalln("Can't listen for new agents, exiting.")
		}
	}()

	listeners.Wait()
}
