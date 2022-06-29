package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halfrollboy/test-nexign/internal/spelling"
)

func CheckCorrect(c *gin.Context) {
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		errors.New("Not valid json")
	}
	var result map[string][]string
	json.Unmarshal([]byte(jsonDataBytes), &result)
	_, ok := result["texts"]

	if ok {
		spelling.SpelingText(result["texts"])
		c.IndentedJSON(http.StatusOK, result)
	} else {
		errors.New("Not valid json, need texts")
	}

}
