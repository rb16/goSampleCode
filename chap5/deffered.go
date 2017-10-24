package main

import (
	"log"
	"time"
)

func bitCalculation() {
	defer timer("bit calulation")()
	time.Sleep(time.Second * 3)

}

func timer(message string) func() {
	t := time.Now()
	return func() {
		log.Printf("%s took %s", message, time.Since(t))

	}
}

func main() {
	bitCalculation()
}
