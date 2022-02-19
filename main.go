package main

import (
	"fmt"
	"log"
)

func main() {
	word, err := GetWordByDate(2, 18, 2022)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word)
}
