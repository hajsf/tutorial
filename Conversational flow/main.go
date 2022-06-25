package main

import (
	"fmt"

	"Conversation/flow"
)

func (u *UserData) String() string {
	return fmt.Sprintf("UserData{login: %q, password: %q}", u.login, u.password)
}

//bot emulator
func newBot(humanChan chan string) *flow.Flow {
	var awaitCommand *flow.Step
	var askEmail *flow.Step
	var askPassword *flow.Step

	//wait for command from human
	awaitCommand =
		flow.OnReply(func(msg flow.Message, data flow.Data) *flow.NextStep {
			switch msg {
			case "register":
				return flow.Goto(askEmail).Using(&UserData{})
			case "quit":
				return flow.End()
			}
			return flow.DefaultHandler()(msg, data)
		})

	//ask human for email
	askEmail =
		flow.Ask(func(data flow.Data) {
			humanChan <- "please enter your email"
		}).OnReply(func(msg flow.Message, data flow.Data) *flow.NextStep {
			email := msg.(string)
			humanData := data.(*UserData)
			humanData.login = email
			return flow.Goto(askPassword)
		})

	//ask human for password
	askPassword =
		flow.Ask(func(data flow.Data) {
			humanChan <- "please enter your password"
		}).OnReply(func(msg flow.Message, data flow.Data) *flow.NextStep {
			password := msg.(string)
			humanData := data.(*UserData)
			humanData.password = password
			return flow.End().Using(humanData)
		})

	return flow.New(awaitCommand)
}

//human emulator
func newHuman(botChan chan string) *flow.Flow {
	var askRegister *flow.Step
	var sendEmail *flow.Step
	var sendPassword *flow.Step

	//send regiter command to bot and process response
	askRegister =
		flow.Ask(func(data flow.Data) {
			botChan <- "register"
		}).OnReply(func(msg flow.Message, data flow.Data) *flow.NextStep {
			switch msg {
			case "please enter your email":
				return flow.Goto(sendEmail)
			}
			return flow.DefaultHandler()(msg, data)
		})

	//send email to bot and process response
	sendEmail =
		flow.Ask(func(data flow.Data) {
			email := GetInput("please enter your email")
			botChan <- email
		}).OnReply(func(msg flow.Message, data flow.Data) *flow.NextStep {
			switch msg {
			case "please enter your password":
				return flow.Goto(sendPassword)
			}
			return flow.DefaultHandler()(msg, data)
		})

	//just send a password to the bot and stop the flow
	sendPassword =
		flow.Ask(func(data flow.Data) {
			botChan <- GetInput("please enter your password")
		})

	return flow.New(askRegister)
}

func main() {
	humanChan := make(chan string)
	botChan := make(chan string)

	//bot conversation flow
	bot := newBot(humanChan)
	//human conversation flow
	human := newHuman(botChan)

	go func() {
		for {
			select {
			//receive massage from human and redirect it to bot
			case toBot := <-botChan:
				bot.Send(toBot)
			//receive massage from bot and redirect it to human
			case toHuman := <-humanChan:
				human.Send(toHuman)
			}
		}
	}()

	human.Start()
	//lock until the end of bot flow
	collectedData := <-bot.Start()
	fmt.Println(collectedData)
}
