package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	exactTime, err := ntp.Time("0.pool.ntp.org")

	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("current time: %s\n", time.Now().Round(0))
	fmt.Printf("exact time: %v\n", exactTime.Round(0))
}
