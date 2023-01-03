package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TelegramBot struct {
	Token     string
	MessageId string
	URL       string
}

func NewTelegramBot(token, messageId string) *TelegramBot {
	bot := &TelegramBot{
		Token:     token,
		MessageId: messageId,
		URL:       fmt.Sprintf("https://api.telegram.org/bot%s", token),
	}

	return bot
}

func (b *TelegramBot) SendMessage(text string) (bool, error) {

	url := fmt.Sprintf("%s/sendMessage", b.URL)

	body, _ := json.Marshal(map[string]string{
		"chat_id": b.MessageId,
		"text":    text,
	})

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer response.Body.Close()

	bData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	fmt.Println(string(bData))

	return true, nil
}
