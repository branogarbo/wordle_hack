package wordle_hack

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

func GetWordByDate(date time.Time) (string, error) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Now().UTC().Location())

	words, err := readJSON("./data/word_list_ordered.json")
	if err != nil {
		return "", err
	}

	wdm := popWordDateMap(words)
	word := wdm[date]

	if word == "" {
		return "", errors.New("the passed date is not within wordle's date range")
	}

	return word, nil
}

func GetWordByString(dateString string) (string, error) {
	words, err := readJSON("./data/word_list_ordered.json")
	if err != nil {
		return "", err
	}

	wdm := popWordDateMap(words)

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

func GetWordByInteger(year, month, day int) (string, error) {
	words, err := readJSON("./data/word_list_ordered.json")
	if err != nil {
		return "", err
	}

	wdm := popWordDateMap(words)

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

func popWordDateMap(words Words) WordDateMap {
	var (
		dodgeIndex    int
		word_date_map = make(WordDateMap)
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

func readJSON(path string) (Words, error) {
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
