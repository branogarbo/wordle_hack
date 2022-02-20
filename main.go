package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	date := time.Date(2022, time.February, 18, 0, 0, 0, 0, time.Now().UTC().Location())

	word, err := GetWordByDate(date)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word)
}
