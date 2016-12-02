# OdgBot

This Telegram bot will be used to create a list of items in a group and
visualize them later via the bot itself or giving a list to a file containing
the list.

## Specification

The normal use of this bot should be to add messages to a list during a week and
discuss them during the weekly meeting asking them to the bot.

After they are discussed it should be possible to clear the list and start again.

The first implementaion must be able to handle a single list of items.

### Commands:

This **bot** will work without the possibility of seeing all the messages in a group
[(privacy mode)](https://core.telegram.org/bots#privacy-mode)
and so to write to him you need to send a message that starts with **/**

- **/start**    starts the bot and prints the help message
- **/help**     prints the help message
- **/add**      add an item to the list
- **/list**     show the list
- **/remove N** remove the item with index N  
- **/clear**    clear the whole list

The commands should be case insensitive

## Testing 

To test the bot you need to ask Telegram for a bot API token. Then you should copy the 
file `config/config_example.json` into `config/config.json` and modify it adding the 
received token.

Then just type:

    go run main.go

and the bot should be running. Check the output to see the received request when
you send a message to it.
