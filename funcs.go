package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

func GetWordByDate(dateString string) (string, error) {
	words, err := ReadJSON("./data/word_list_ordered.json")
	if err != nil {
		return "", err
	}

	wdm := PopWordDateMap(words)

	date, err := time.Parse("2006-01-02T15:04:05.000Z", dateString+"T00:00:00.000Z")
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
		word_date_map = make(word_date_map)
		dodgeIndex    = 244
		dodgeDay      = time.Date(2022, time.February, 18, 0, 0, 0, 0, time.Now().UTC().Location())
	)

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
