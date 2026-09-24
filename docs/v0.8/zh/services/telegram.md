# Telegram

## URL 格式

:::info
telegram://__`token`__@telegram?chats=__`channel-1`__[,__`chat-id-1`__,...]
:::

### URL 字段

*  __Token__ （**必填**）  
  URL 位置：<code class="service-url">telegram://<strong>token</strong>@telegram/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Chats__ - Chat IDs or Channel names (using @channel-name) （**必填**）  
  别名：`channels`  

*  __Notification__ - If disabled, sends Message silently  
  默认值：✔ `Yes`  

*  __ParseMode__ - How the text Message should be parsed  
  默认值：`None`  
  可选值：`None`, `Markdown`, `HTML`, `MarkdownV2`  

*  __Preview__ - If disabled, no web page preview will be displayed for URLs  
  默认值：✔ `Yes`  

*  __Title__ - Notification title, optionally set by the sender  
  默认值：*empty*

## 获取 Telegram 令牌

找 [the botfather](https://core.telegram.org/bots#6-botfather) 申请。

## 确定目标聊天/频道

`chats` 参数由一个或多个 `Chat ID` 或 `channel name` 组成。

### 公开频道
公开频道的频道名可以在 Telegram 客户端的 `Channel info` 区域找到。
把链接中的 `t.me/` 前缀替换为 `@` 即可。

:::info
频道名必须以 `@` 开头，以表明它是频道名。
:::

:::info
如果你的频道只有邀请链接（以 `t.me/+` 开头），则必须使用它的 Chat ID（见下文）。
:::

:::info
可以添加 `message_thread_id` 参数（[参考文档](https://core.telegram.org/bots/api#sendmessage)），格式为 `$chat_id:$message_thread_id`。关于如何获取 `message_thread_id` 的[更多信息](https://stackoverflow.com/questions/74773675/how-to-get-topic-id-for-telegram-group-chat/75178418#75178418)。
:::

### 聊天
私有频道、群组聊天和私聊通过 `Chat ID` 标识。遗憾的是，它们通常在 Telegram 客户端里看不到。
最简单的获取方式是使用 `shoutrrr generate telegram` 命令，它会引导你一步步创建包含目标聊天的 URL。

:::tip
你可以在 Docker 中使用 `marrrrrrrrry/shoutrrr` 镜像来运行它，无需下载/安装 `shoutrrr` CLI：
```
docker run --rm -it marrrrrrrrry/shoutrrr generate telegram
```
:::

### 询问 @shoutrrrbot
另一种获取 Chat ID 的方法，是把目标聊天中的一条消息转发给 [@shoutrrrbot](https://t.me/shoutrrrbot)。
它会回复这条转发消息 originally 所在聊天的 Chat ID。
注意，它对群组聊天不太好用，因为那些消息只会被看作由某个用户发出，而不是发在特定聊天里。
作为替代，你可以使用第二种方法：把 @shoutrrrbot 拉进你的群聊，并向它发送一条消息（消息以 @shoutrrrbot 开头）。之后你就可以放心把它踢出群聊。

这个机器人需要保持在线，除非它的用量超过了 GCP 的免费额度。它的源码在 [github.com/marrrrrrrrry/shoutrrrbot](https://github.com/marrrrrrrrry/shoutrrrbot)。



## 可选参数

你可以在 URL 中可选地指定 __`notification`__、__`parseMode`__ 和 __`preview`__ 参数：

:::info
<pre>telegram://__`token`__@__`telegram`__/?channels=__`channel`__&notification=no&preview=false&parseMode=html</pre>
:::

更多信息参见 [Telegram 官方文档](https://core.telegram.org/bots/api#sendmessage)。

:::info
`preview` 和 `notification` 与 Telegram API 中的对应参数（`disable_web_page_preview` 和 `disable_notification`）语义相反。
:::

### 解析模式与标题

如果指定了解析模式，消息必须按照
[Formatting options](https://core.telegram.org/bots/api#formatting-options) 中对应章节的规则进行转义。

指定了标题时，标题会被加在消息前面，但这只在 `HTML` 解析模式下受支持。注意，如果没有指定解析模式，消息会按 `HTML` 转义并发送。

由于 markdown 模式很难正确转义，建议始终使用 `HTML` 解析模式。
