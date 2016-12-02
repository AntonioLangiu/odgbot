# OdgBot

This Telegram bot will be used to create a list of items in a group and
visualize them later.

The normal use of this bot should be to add messages to a list during a week and
discuss them during the weekly meeting asking them to the bot.

After they are discussed it should be possible to clear the list and start again

## Testing 
To test the bot you need to ask Telegram for a bot API token. Then you should copy the 
file `config/config_example.json` into `config/config.json` and modify it adding the 
received token.

Then just type:

    go run main.go

and the bot should be running. Check the output to see the received request when
you send a message to it.
