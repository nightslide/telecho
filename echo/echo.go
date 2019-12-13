package main

import (
	"net/http"
	"os"
	"log"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
	"strconv"
)

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	apiUrl := "https://api.telegram.org/bot"
	botUrl := apiUrl + botToken
	offset := 0
	for ;; {
		updates, err := getUpdates(botUrl, offset)
		if err != nil {
			log.Fatal(err.Error())
		}
		for _, update := range updates {
			err = respond(botUrl, update)
			fmt.Println(update)
			if err != nil {
				log.Print("Smth went wrong: ", err.Error())
			}
			offset = update.UpdateId + 1
		}
	}
}

func getUpdates(botUrl string, offset int) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates?offset=" + strconv.Itoa(offset))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var updates RestResponse
	err = json.Unmarshal(body, &updates)
	if err != nil {
		return nil, err
	}
	return updates.Result, nil
}

func respond(botUrl string, update Update) (error) {
	var response BotMessage
	response.ChatId = update.Message.Chat.ChatId
	response.Text = update.Message.Text
	buf, err := json.Marshal(response)
	if err != nil {
		return err
	}
	_, err = http.Post(botUrl + "/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	return nil
}