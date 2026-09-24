# Zulip Chat

## URL Format

The shoutrrr service URL should look like this:

:::info
zulip://__`botmail`__:__`botkey`__@__`host`__/?stream=__`stream`__&topic=__`topic`__
:::

### URL Fields

*  __BotMail__ - Bot e-mail address (**Required**)  
  URL part: <code class="service-url">zulip://<strong>botmail</strong>:botkey@host:port/</code>  
*  __BotKey__ - API Key (**Required**)  
  URL part: <code class="service-url">zulip://botmail:<strong>botkey</strong>@host:port/</code>  
*  __Host__ - API server hostname (**Required**)  
  URL part: <code class="service-url">zulip://botmail:botkey@<strong>host</strong>:<strong>port</strong>/</code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Stream__  
  Default: *empty*  

*  __Topic__  
  Default: *empty*  
  Aliases: `title`

:::info
Since __`botmail`__  is a mail address you need to URL escape the `@` in it to `%40`.
:::

### Examples

Stream and topic are both optional and can be given as parameters to the Send method:

```go
  sender, _ := shoutrrr.CreateSender(url)

  params := make(types.Params)
  params["stream"] = "mystream"
  params["topic"] = "This is my topic"

  sender.Send(message, &params)
```

:::tip Example service URL
zulip://my-bot%40zulipchat.com:correcthorsebatterystable@example.zulipchat.com?stream=foo&topic=bar
:::
