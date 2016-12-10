package bot

import (
    "log"
    "github.com/AntonioLangiu/odgbot/common"
    "gopkg.in/telegram-bot-api.v4"
)

func LoadBot(configuration *common.Configuration) {
	// Start the Bot
	bot, err := tgbotapi.NewBotAPI(configuration.TelegramAPI)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	// TODO: check err

	var out string
	// Handle each request
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			var command string = update.Message.Command()
			// general commands
			// if the message /start is received add the group or user
			// to the db, if it's already present print the help message
			// if help is received print the help message
			log.Printf("Command is %s\n\n",command)
			if command == "start" {
				out = "This bot saves a list of things"
			} else if command == "help" {
				out = "This bot saves a list of things"

				// user commands
				// Check if the user or group is present in the db
				// if not do nothing
				// /add
				// /list
				// /remove N
				// /xxx add with a tag
			} else if command == "list" {

			} else if command == "remove" {

			} else {
			}
		} else {
			// This type of messages could be:
			//   - user added to a gruop
			//   - user removed from a group
			//   - bot removed from a group
			//   . if the bot is used not in a group then every message that doesn't start with /
			//   - what more?
			continue;
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, out)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
