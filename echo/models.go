package main

type RestResponse struct {
	Result []Update		`json:"result"`
}

type Update struct {
	UpdateId int		`json:"update_id"`
	Message UserMessage	`json:"message"`
}

type UserMessage struct {
	Text string 		`json:"text"`
	Chat Chat 			`json:"chat"`
}

type Chat struct {
	ChatId int 			`json:"id"`
	FirstName string	`json:"first_name"`
}

type BotMessage struct {
	ChatId int			`json:"chat_id"`
	Text string			`json:"text"`
}