# Telegram generator

The Telegram generator creates a Telegram service URL by talking to your bot: it validates the bot
token, fetches the bot info and then listens for incoming messages so you can pick the chats to
include by simply sending a message to the bot (or adding it to a group or channel).

## Usage

```shell
$ shoutrrr generate telegram
```

1. Enter your bot token (if you don't have a bot yet, the generator links to [@BotFather](https://t.me/botfather?start)).
2. The generator fetches the bot info and starts waiting for messages.
3. Send a message to the bot in each chat you want to add — direct messages, groups or channels.
4. When you have added all chats, answer **No** when asked whether to add more.
5. The generator cleans up the bot session and prints the resulting URL, e.g.:

```text
URL: telegram://110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw@telegram?chats=@mychannel
```
