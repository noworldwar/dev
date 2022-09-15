package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	tgtoken = "2066047122:AAGMxupBSwAvpTUpvnIgBrQA_k9cYtT_OzQ"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// Main funtion starts our server on port 3000
func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}

// This handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	text := strings.ToLower(body.Message.Text)

	switch {
	case strings.Contains(text, "get secret key"):
	case strings.Contains(text, "set sw"):
	case strings.Contains(text, "get whitelist"):
	case strings.Contains(text, "set whitelist"):
	default:
		autoRespond(body.Message.Chat.ID, "Can't recognize the command")
	}

	s := strings.Fields(text)
	fmt.Println(s)

	var err error
	switch {
	case strings.Contains(text, "get secret key"):
		env := s[3]
		lang := s[4]
		agentID := s[5:]
		err = GetSecretKey(env, lang, agentID)
	case strings.Contains(text, "set sw"):
		env := s[2]
		lang := s[3]
		operatorID := s[4]
		endpoint := s[5]
		err = SetSingleWallet(env, lang, operatorID, endpoint)
	case strings.Contains(text, "get whitelist"):
		// err = GetWhiteList()
	case strings.Contains(text, "set whitelist"):
		// err = SetWhitelList()
	default:
		break
	}

	if err != nil {
		fmt.Println("error in function:", err)
		return
	}

	if err := autoRespond(body.Message.Chat.ID, "123"); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	fmt.Println("reply sent")
}

// send reply
func autoRespond(chatID int64, msg string) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   msg,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post("https://api.telegram.org/bot"+tgtoken+"/sendMessage", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

/*
get secret key env lang agentID, agentID...
set sw env lang operatorID endpoint
get whitelist env agent: ip:
set whitelist env agent type(bo|api) ip1,ip2,ip3...
*/

func GetSecretKey(env, lang string, agentID []string) error {
	return nil
}

func SetSingleWallet(env, lang, operatorID, endpoint string) error {
	return nil
}

func GetWhiteList(env, agent, ip string) error {
	return nil
}

func SetWhitelList(env, agent, setType string, ip []string) error {
	return nil
}
