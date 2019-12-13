package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	apiUrl := "https://api.telegram.org/bot"
	botUrl := apiUrl + botToken
	resp, err := http.Get(botUrl + "/getMe")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
}
