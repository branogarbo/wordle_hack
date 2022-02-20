package main

import (
	"fmt"
	"time"
)

func main() {
	word, _ := GetWordByDate(time.Now())
	fmt.Println(word)

	word, _ = GetWordByString("2022-02-18")
	fmt.Println(word)

	word, _ = GetWordByInteger(2022, 2, 18)
	fmt.Println(word)
}
