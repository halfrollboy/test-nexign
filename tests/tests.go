package greetings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/halfrollboy/test-nexign/internal/spelling"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSpelling(t *testing.T) {
	data, test := getJson()
	spelling.SpelingText(data)
	for i, row := range test {
		if data[i] != row {
			t.Errorf("String does not match,with testing row")
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

	jsonData, err := os.Open("data.json")
	jsonTest, err := os.Open("test.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
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
