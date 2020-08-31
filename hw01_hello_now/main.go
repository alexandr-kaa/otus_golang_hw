package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// layout:="09 Jan 2006 01:02:01 +0000 UTC"
	currTime := time.Now()
	timeExact, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatalf("Обнаружена ошибка при обращении к ntp server:\n " + err.Error())
	}

	currTime = currTime.Round(0)
	timeExact = timeExact.Round(0)

	// fmt.Println("```text")
	fmt.Println("current time: " + currTime.String())
	fmt.Println("exact time: " + timeExact.String())
	// fmt.Println("```")
}
