package main

func main() {
	words, err := ReadJSON("./data/word_history.json")
	if err != nil {
		panic(err)
	}

	for _, word := range words {
		println(word)
	}
}
