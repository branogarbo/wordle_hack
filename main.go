package main

import (
	"fmt"
	"log"
)

func main() {
	word, err := GetWordByDate("2022-02-18")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word)
}
