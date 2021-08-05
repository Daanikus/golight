package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/daanikus/golight/effects"
	"github.com/daanikus/golight/lights"
)

func main() {
	l, err := lights.New(32)
	if err != nil {
		log.Fatal(err)
	}

	cleanup := func() {
		l.Off()
		l.Close()
	}
	defer cleanup()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	palette := lights.Palettes["synthwave"]

	effects.Stream(l, &palette, 10*time.Second)
}
