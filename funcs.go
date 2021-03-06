package wordle_hack

import (
	"errors"
	"time"
)

func GetWordByDate(date time.Time) (string, error) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)

	wdm := popWordDateMap(words)
	word := wdm[date]

	if word == "" {
		return "", errors.New("the passed date is not within wordle's date range")
	}

	return word, nil
}

func GetWordByString(dateString string) (string, error) {
	date, err := time.Parse("2006-01-02T15:04:05.000Z", dateString+"T00:00:00.000Z")
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

func GetWordByInteger(year, month, day int) (string, error) {
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	wdm := popWordDateMap(words)
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
		dodgeDay      = time.Date(2022, time.February, 18, 0, 0, 0, 0, time.UTC)
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
