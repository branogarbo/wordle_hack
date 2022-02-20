package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	date := time.Now().Add(time.Hour * -24)

	word, err := GetWordByDate(date)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word)
}
