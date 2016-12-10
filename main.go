package main

import (
    "github.com/AntonioLangiu/odgbot/bot"
    "github.com/AntonioLangiu/odgbot/common"
)

func main() {
    configuration := common.LoadConfiguration()
    db := bot.LoadDb(configuration)
    defer db.Close()
    bot.LoadBot(configuration)
}
