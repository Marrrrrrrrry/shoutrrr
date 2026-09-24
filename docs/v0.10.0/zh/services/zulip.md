# Zulip Chat

## URL 格式

shoutrrr 服务 URL 应如下所示：

:::info
zulip://__`botmail`__:__`botkey`__@__`host`__/?stream=__`stream`__&topic=__`topic`__
:::

### URL 字段

*  __BotMail__ - Bot e-mail address （**必填**）  
  URL 位置：<code class="service-url">zulip://<strong>botmail</strong>:botkey@host:port/</code>  
*  __BotKey__ - API Key （**必填**）  
  URL 位置：<code class="service-url">zulip://botmail:<strong>botkey</strong>@host:port/</code>  
*  __Host__ - API server hostname （**必填**）  
  URL 位置：<code class="service-url">zulip://botmail:botkey@<strong>host</strong>:<strong>port</strong>/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Stream__  
  默认值：*empty*  

*  __Topic__  
  默认值：*empty*  
  别名：`title`


:::info
由于 __`botmail`__ 是一个邮件地址，你需要把其中的 `@` URL 转义为 `%40`。
:::

### 示例

stream 和 topic 都是可选的，也可以作为参数传给 Send 方法：

```go
  sender, _ := shoutrrr.CreateSender(url)

  params := make(types.Params)
  params["stream"] = "mystream"
  params["topic"] = "This is my topic"

  sender.Send(message, &params)
```

:::tip 服务 URL 示例
zulip://my-bot%40zulipchat.com:correcthorsebatterystable@example.zulipchat.com?stream=foo&topic=bar
:::
