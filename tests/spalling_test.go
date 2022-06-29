package spalling_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/halfrollboy/test-nexign/internal/spelling"
)

func TestSpelling(t *testing.T) {
	//Получение данных из подготовленных файлов
	data, test := getJson()
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

	// data, test := getJson()
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
	data, _ := getJson()
	b.ResetTimer()
	spelling.SpelingText(data)
}

//Вспомогательная ф-ци которая берёт данные из файлов
func getJson() ([]string, []string) {

	jsonData, err := os.Open("tests/test_data/data.json")
	jsonTest, err := os.Open("tests/test_data/test.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonData.Close()
	defer jsonTest.Close()

	byteData, _ := ioutil.ReadAll(jsonData)
	byteTest, _ := ioutil.ReadAll(jsonTest)

	var resultData map[string][]string
	var resultTest []string
	json.Unmarshal([]byte(byteData), &resultData)
	json.Unmarshal([]byte(byteTest), &resultTest)

	textsData := resultData["texts"]
	textsTest := resultTest
	return textsData, textsTest
}
