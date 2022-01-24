package main

import (
	"strings"

	_ "github.com/lib/pq"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {

	// confighelper.InitViper()
	e := echo.New()
	e.POST("/getWordCountService", GetWordCountService)

	e.Logger.Fatal(e.Start(":4000"))
}

//To get the word count
func GetWordCountService(c echo.Context) error {

	body, _ := GetRequestBodyJson(c)
	InputString := body.Get("inputString").String()

	stripped := string(InputString)
	// stripped := "AA BB AA BB AA CC"
	totalCount := 0
	list := []interface{}{}

	for index, element := range wordCount(stripped) {
		totalCount++
		ldDetails, _ := sjson.Set("{}", "word", index)
		ldDetails, _ = sjson.Set(ldDetails, "count", element)
		ldData := gjson.Parse(ldDetails)
		list = append(list, ldData.Value())

		// fmt.Println(index, "=>", element)
	}
	fmt.Println(":::::::::::::::::::::::", list)
	return c.JSON(http.StatusOK, list)
}

func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		_, ok := counts[word]
		if ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

//Read the URL and retun the json string
func GetRequestBodyJson(c echo.Context) (body gjson.Result, err error) {
	bb, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		fmt.Println("Error Reading Request Body")
		return gjson.Result{}, err
	}

	return gjson.ParseBytes(bb), nil
}
