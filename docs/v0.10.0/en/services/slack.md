# Slack

:::caution New URL format
The URL format for Slack has been changed to allow for API- as well as webhook tokens.  
Using the old format (`slack://xxxx/yyyy/zzzz`) will still work as before and will automatically be upgraded to
the new format when used.
:::

The Slack notification service uses either [Slack Webhooks](https://api.slack.com/messaging/webhooks) or the 
[Bot API](https://api.slack.com/methods/chat.postMessage) to send messages.  

:::tip Related guide
This page is the reference for the URL format and props.  
For step-by-step instructions on how to get your *token* and *channel*, see the
[Slack Token Guide](/guides/slack/).
:::

:::caution Legacy webhook tokens are deprecated
Slack has deprecated the legacy custom-integration incoming webhooks. They keep working for now
but have no removal date and may stop working at any time. Prefer a Bot API token
(`xoxb`, see the [Token Guide](/guides/slack/)) or an incoming webhook created through a Slack app.
:::


## URL Format

:::info
Note that the token uses a prefix to determine the type, usually either `hook` (for webhooks) or `xoxb` (for bot API).
:::

### URL Fields

*  __Token__ - API Bot token (**Required**)  
  URL part: <code class="service-url">slack://<strong>token</strong>:<strong>token</strong>@channel/</code>  
*  __Channel__ - Channel to send messages to in Cxxxxxxxxxx format (**Required**)  
  URL part: <code class="service-url">slack://token:token@<strong>channel</strong>/</code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __BotName__ - Bot name  
  Default: *empty*  
  Aliases: `username`  

*  __Color__ - Message left-hand border color  
  Default: *empty*  

*  __Icon__ - Use emoji or URL as icon (based on presence of http(s):// prefix)  
  Default: *empty*  
  Aliases: `icon_emoji`, `icon_url`  

*  __ThreadTS__ - ts value of the parent message (to send message as reply in thread)  
  Default: *empty*  

*  __Title__ - Prepended text above the message  
  Default: *empty*

:::info Color format
The format for the `Color` prop follows the [slack docs](https://api.slack.com/reference/messaging/attachments#fields)
but `#` needs to be escaped as `%23` when passed in a URL.  
So <span style="background:#ff8000;width:.9em;height:.9em;display:inline-block;vertical-align:middle"></span><code>#ff8000</code> would be `%23ff8000` etc.
:::

## Examples

:::tip Bot API
```text
slack://xoxb:123456789012-1234567890123-4mt0t4l1YL3g1T5L4cK70k3N@C001CH4NN3L?color=good&title=Great+News&icon=man-scientist&botname=Shoutrrrbot
```
:::

:::tip Webhook
```text
slack://hook:WNA3PBYV6-F20DUQND3RQ-Webc4MAvoacrpPakR8phF0zi@webhook?color=good&title=Great+News&icon=man-scientist&botname=Shoutrrrbot
```
:::