package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

func GetWordByDate(month, day, year int) (string, error) {
	words, err := ReadJSON("./data/word_list_ordered.json")
	if err != nil {
		return "", err
	}

	wdm := PopWordDateMap(words)

	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Now().UTC().Location())
	if err != nil {
		return "", err
	}

	word := wdm[date]

	if word == "" {
		return "", errors.New("the passed date is not within wordle's date range")
	}

	return word, nil
}

func PopWordDateMap(words Words) word_date_map {
	var (
		dodgeIndex    int
		word_date_map = make(word_date_map)
		dodgeDay      = time.Date(2022, time.February, 18, 0, 0, 0, 0, time.Now().UTC().Location())
	)

	for i, w := range words {
		if w == "dodge" {
			dodgeIndex = i
			break
		}
	}

	for wordIndex, word := range words {
		var (
			dayOffset = wordIndex - dodgeIndex
			wordDay   = dodgeDay.AddDate(0, 0, dayOffset)
		)

		word_date_map[wordDay] = word
	}

	return word_date_map
}

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
