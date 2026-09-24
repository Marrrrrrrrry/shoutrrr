# Teams

:::caution Webhook URL 格式已变更
Microsoft 已经更改了 Teams Webhook 的 URL 格式。现在你必须通过以下方式指定主机名：
```text
?host=example.webhook.office.com
```
其中 `example` 是你的组织短名称。
:::

## URL 格式

:::info
teams://__`group`__@__`tenant`__/__`altId`__/__`groupOwner`__?host=__`organization`__.webhook.office.com
:::

### URL 字段

*  __Group__  
  默认值：*empty*  
  URL 位置：<code class="service-url">teams://<strong>group</strong>@tenant/altid/groupowner</code>  
*  __Tenant__  
  默认值：*empty*  
  URL 位置：<code class="service-url">teams://group@<strong>tenant</strong>/altid/groupowner</code>  
*  __AltID__  
  默认值：*empty*  
  URL 位置：<code class="service-url">teams://group@tenant/<strong>altid</strong>/groupowner</code>  
*  __GroupOwner__  
  默认值：*empty*  
  URL 位置：<code class="service-url">teams://group@tenant/altid/<strong>groupowner</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Color__  
  默认值：*empty*  

*  __Host__  
  默认值：`outlook.office.com`  

*  __Title__  
  默认值：*empty*

## 设置 Webhook

要使用 Microsoft Teams 通知服务，首先需要设置一个自定义 Webhook。
具体操作方法见[这份指南](https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/connectors-using#setting-up-a-custom-incoming-webhook)。

## 提取令牌

令牌从你的 Webhook URL 中提取：

<pre><code>https://<b>&lt;organization&gt;</b>.webhook.office.com/webhookb2/<b>&lt;group&gt;</b>@<b>&lt;tenant&gt;</b>/IncomingWebhook/<b>&lt;altId&gt;</b>/<b>&lt;groupOwner&gt;</b></code></pre>
