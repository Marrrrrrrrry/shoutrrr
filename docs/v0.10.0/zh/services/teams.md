# Teams

:::caution Office 365 Connector Webhook 已退役
旧版 `teams://group@tenant/altId/groupOwner` URL 所依赖的 Office 365 Connector 已被 Microsoft 退役：
2024 年 8 月起禁止新建 connector，2026 年 4 月起存量 connector 停止投递，2026 年 5 月全面禁用。

请改用 **Power Automate Workflows** 的请求 URL（见[设置 Webhook](#设置-webhook)）。
旧版 URL 格式仅保留用于报错诊断，不再能投递消息。
:::

## URL 格式

:::info Workflows Webhook（现行）
<pre><code class="service-url">teams://?webhook=<strong>&lt;workflows 请求 URL&gt;</strong></code></pre>

也可以直接在 workflow URL 前加 `teams+` 前缀：

<pre><code class="service-url">teams+https://<strong>&lt;region&gt;</strong>.logic.azure.com:443/workflows/<strong>&lt;workflow-id&gt;</strong>/triggers/manual/paths/invoke?api-version=2016-06-01&amp;sp=...&amp;sv=...&amp;sig=...</code></pre>
:::

### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` 等。

*  __Webhook__  
  完整的 Power Automate Workflows 请求 URL。不能与旧版 URL 字段同时使用。  
  默认值：*empty*

*  __Color__  
  默认值：*empty*  

*  __Host__  
  默认值：`outlook.office.com`  
  仅旧版 connector URL 格式使用。

*  __Title__  
  默认值：*empty*

## 设置 Webhook

1. 在 Teams 中打开你想发消息的频道，选择 **更多操作 (...)** → **Workflows**。
2. 选择模板 **Post to a channel when a webhook request is received** 并创建 workflow。
3. 从 workflow 详情页复制请求 URL，形如：
   `https://<region>.logic.azure.com:443/workflows/<workflow-id>/triggers/manual/paths/invoke?api-version=2016-06-01&sp=...&sv=...&sig=...`
4. 以 `teams://?webhook=<url>`（URL 转义后）或加 `teams+` 前缀的方式使用：
   `teams+https://<region>.logic.azure.com/...`

:::info 消息格式
服务发送的是经典 MessageCard 格式。Workflows webhook 可以渲染 MessageCard，
但交互元素（HttpPost 按钮动作）不受支持。消息大小上限 28 KB，发送者显示为 Workflows (Flow) bot。
:::

## 旧版 URL 格式（已退役）

基于 connector 的格式把 Webhook URL 拆分为四段：

<pre><code class="service-url">teams://<strong>group</strong>@<strong>tenant</strong>/<strong>altId</strong>/<strong>groupOwner</strong>?host=<strong>&lt;organization&gt;</strong>.webhook.office.com</code></pre>

*  __Group__  
  URL 位置：<code class="service-url">teams://<strong>group</strong>@tenant/altid/groupowner</code>  
*  __Tenant__  
  URL 位置：<code class="service-url">teams://group@<strong>tenant</strong>/altid/groupowner</code>  
*  __AltID__  
  URL 位置：<code class="service-url">teams://group@tenant/<strong>altid</strong>/groupowner</code>  
*  __GroupOwner__  
  URL 位置：<code class="service-url">teams://group@tenant/altid/<strong>groupowner</strong></code>  

令牌过去从 connector Webhook URL 中提取：

<pre><code>https://<b>&lt;organization&gt;</b>.webhook.office.com/webhookb2/<b>&lt;group&gt;</b>@<b>&lt;tenant&gt;</b>/IncomingWebhook/<b>&lt;altId&gt;</b>/<b>&lt;groupOwner&gt;</b></code></pre>

由于 connector Webhook 已无法投递消息，该格式现在只会产生错误。
