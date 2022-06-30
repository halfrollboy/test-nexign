package spalling_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/halfrollboy/test-nexign/internal/spelling"
)

func TestSpelling(t *testing.T) {
	//Получение данных из подготовленных файлов
	data := getJson("tests/test_data/data.json")
	test := getJson("tests/test_data/test.json")
	testTable := []struct {
		text   []string
		result []string
	}{
		{
			text:   data,
			result: test,
		},
		{
			text:   []string{"новая машына"},
			result: []string{"новая машина"},
		},
	}

	for _, val := range testTable {
		spelling.SpelingText(val.text)
		for i, rowResult := range val.result {
			if val.text[i] != rowResult {
				t.Errorf("String '%s' does not match,with testing row '%s'", val.text[i], rowResult)
			}
		}
	}
}

func BenchmarkSpelling(b *testing.B) {
	data := getJson("tests/test_data/data.json")
	b.ResetTimer()
	spelling.SpelingText(data)
}

//Вспомогательная ф-ци которая берёт данные из файлов
func getJson(name string) []string {

	jsonData, errData := os.Open(name)

	if errData != nil {
		errors.New("Error open test data")
	}

	defer jsonData.Close()

	byteData, _ := ioutil.ReadAll(jsonData)
	var resultData map[string][]string
	json.Unmarshal([]byte(byteData), &resultData)

	textsData := resultData["texts"]
	return textsData
}
