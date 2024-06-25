package main

import (
	"log"
	"os"
)

func main() {
	d := os.Getenv("RG-ASEPP")
	log.Println(d)
}
