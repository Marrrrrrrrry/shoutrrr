# Telegram 生成器

Telegram 生成器通过与你的机器人对话来生成 Telegram 服务 URL：它会校验 bot token、获取机器人信息，
然后监听新消息——你只需给机器人发消息（或把它拉进群组/频道）即可选择要包含的聊天。

## 用法

```shell
$ shoutrrr generate telegram
```

1. 输入你的 bot token（还没有机器人的话，生成器会给出 [@BotFather](https://t.me/botfather?start) 的链接）。
2. 生成器获取机器人信息后开始等待消息。
3. 在每个想添加的聊天（私聊、群组或频道）里给机器人发一条消息。
4. 添加完所有聊天后，在"是否继续添加"的提示处回答 **No**。
5. 生成器清理 bot 会话并输出结果 URL，例如：

```text
URL: telegram://110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw@telegram?chats=@mychannel
```
