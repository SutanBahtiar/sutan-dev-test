package main

import (
	"encoding/json"
	"strings"
)

func SetClientMessage(message Message) string {
	id := message.ID
	from := message.From
	to := message.To
	fizz := message.Fizz
	buzz := message.Buzz

	if len(id) == 0 {
		return "Error parse message"
	}

	if from > to {
		return "'From' greater from 'to'"
	}

	if from == to {
		return "'From' equals with 'to'"
	}

	var sb strings.Builder
	sb.WriteString("{")
	sb.WriteString("\"" + id + "\"")

	/*
		{"id":"one","from":0,"to":15,"fizz":"zzif","buzz":"zzub"}\n
		{"one":["zzifzzub",1,2,"zzif",4,"zzub","zzif",7,8,"zzif","zzub",11,"zzif",13,14,"zzifzzub"]}\n

		return fizz if the integer is divisible by 3,
		return buzz if the integer is divisible by 5,
		and return fizz+buzz if the integer is divisible by 3 and 5.
	*/

	arr := make([]interface{}, to-from+1)
	f := from
	for i := 0; i <= to-from; i++ {
		var fizzbuzz string
		if f%3 == 0 {
			fizzbuzz = fizz
		}
		if f%5 == 0 {
			fizzbuzz += buzz
		}

		if len(fizzbuzz) > 0 {
			arr[i] = fizzbuzz
		} else {
			arr[i] = f
		}
		f++
	}

	// convert to json
	m, _ := json.Marshal(arr)

	sb.WriteString(string(m))
	sb.WriteString("}")
	sb.WriteString("\\n")

	return sb.String()
}
