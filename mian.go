package main

import (
	"flag"
	"log"
)

var hubd = flag.String("who", "world", "Say hello to who")

func main() {
	flag.Parse()
	log.Println("Hello,", *hubd)
}
