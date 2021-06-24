package main

type Message struct {
	ID   string `json:"id"`
	From int    `json:"from"`
	To   int    `json:"to"`
	Fizz string `json:"fizz"`
	Buzz string `json:"buzz"`
}
