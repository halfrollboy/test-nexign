package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/halfrollboy/test-nexign/internal/spelling"
)

type SpelingRequest struct {
	texts map[string][]string
}

// GetLetters godoc
// Tags spelling
// @Summary      Page with spalling
// @Description  Fetch rows and spelling
// @Accept       json
// @Produce      json
// @Param input body SpelingRequest true "tests for spelling"
// @Success      200  {object}  string
// @Failure      400  {object}  string
// @Failure      422  {object}  string
// @Failure      500  {object}  string
// @Router       /letters [post]
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
