package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestHandleMessage(t *testing.T) {

	/*
		{"id":"one","from":0,"to":15,"fizz":"zzif","buzz":"zzub"}\n
		{"one":["zzifzzub",1,2,"zzif",4,"zzub","zzif",7,8,"zzif","zzub",11,"zzif",13,14,"zzifzzub"]}\n
	*/

	msg := Message{
		ID:   "two",
		From: 6,
		To:   10,
		Fizz: "/n/n",
		Buzz: "/n2/n2",
	}

	resp := SetClientMessage(msg)
	t.Log(string(resp))
}

func TestJsonMarshal(t *testing.T) {
	msg := "{\"id\":\"one\",\"from\":0,\"to\":15,\"fizz\":\"zzif\",\"buzz\":\"zzub\"}\\n"
	t.Log("Message:", msg)

	if strings.Contains(msg, "\\n") {
		msg = strings.ReplaceAll(msg, "\\n", "")
	}

	var message Message
	json.Unmarshal([]byte(msg), &message)
	t.Log("Marshal:", message)
}

func TestRemoveChar(t *testing.T) {
	msg := "jakarta timur\\n"
	t.Log("Message:", msg)

	msg = string(msg[len(msg)-2:])
	t.Log("Message:", msg)
}
