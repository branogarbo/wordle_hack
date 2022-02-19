package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	word, err := GetWordByDate(2022, time.February, 18)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word)
}
