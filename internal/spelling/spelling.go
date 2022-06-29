package spelling

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type YandexResponse struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

// основная ф-ция проверки корректности
func SpelingText(rows []string) {
	errorWords := getJsons(formateText(rows))
	findErrors(rows, errorWords)
}

//ф-ия получения данных из сервиса яндекса
func getJsons(str string) [][]YandexResponse {
	resp, err := http.Get("https://speller.yandex.net/services/spellservice.json/checkTexts?text=" + str)
	if err != nil {
		errors.New("No response from request")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // response body is []byte
	var targets [][]YandexResponse
	json.Unmarshal([]byte(body), &targets)
	return targets
}

//форматирование строк для запроса
func formateText(mass []string) string {
	s := strings.Join(mass, "&text=")
	s = strings.Replace(s, " ", "+", -1)
	return s
}

//поиск и замена слов с ощибками
func findErrors(rows []string, errMas [][]YandexResponse) {
	for id, val := range errMas {
		if len(val) == 0 {
			continue
		}
		for _, word := range val {
			rows[id] = strings.Replace(rows[id], word.Word, word.S[0], -1)

		}
	}
}
