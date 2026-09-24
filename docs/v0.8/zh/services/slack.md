# Slack

:::caution URL 格式已更新
Slack 的 URL 格式已更新，现在同时支持 API 令牌和 Webhook 令牌。
使用旧格式（`slack://xxxx/yyyy/zzzz`）仍然有效，并会在使用时自动升级为新格式。
:::

Slack 通知服务使用 [Slack Webhooks](https://api.slack.com/messaging/webhooks) 或
[Bot API](https://api.slack.com/methods/chat.postMessage) 发送消息。

获取 *token* 和 *channel* 的方法参见[操作指南](/zh/guides/slack/)。


## URL 格式

:::info
注意，令牌使用前缀来区分类型，通常是 `hook`（Webhook）或 `xoxb`（Bot API）。
:::

### URL 字段

*  __Token__ - API Bot token （**必填**）  
  URL 位置：<code class="service-url">slack://<strong>token</strong>:<strong>token</strong>@channel/</code>  
*  __Channel__ - Channel to send messages to in Cxxxxxxxxxx format （**必填**）  
  URL 位置：<code class="service-url">slack://token:token@<strong>channel</strong>/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __BotName__ - Bot name  
  默认值：*empty*  
  别名：`username`  

*  __Color__ - Message left-hand border color  
  默认值：*empty*  

*  __Icon__ - Use emoji or URL as icon (based on presence of http(s):// prefix)  
  默认值：*empty*  
  别名：`icon_emoji`, `icon_url`  

*  __ThreadTS__ - ts value of the parent message (to send message as reply in thread)  
  默认值：*empty*  

*  __Title__ - Prepended text above the message  
  默认值：*empty*

:::info Color format
The format for the `Color` prop follows the [slack docs](https://api.slack.com/reference/messaging/attachments#fields)
but `#` needs to be escaped as `%23` when passed in a URL.  
So <span style="background:#ff8000;width:.9em;height:.9em;display:inline-block;vertical-align:middle"></span><code>#ff8000</code> would be `%23ff8000` etc.
:::

## 示例

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
