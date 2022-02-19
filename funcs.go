package main

import (
	"encoding/json"
	"os"
)

func ReadJSON(path string) (Words, error) {
	var wordsOut Words

	byteOut, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteOut, &wordsOut)
	if err != nil {
		return nil, err
	}

	return wordsOut, nil
}
