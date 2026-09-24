# OpsGenie

## URL 格式

### URL 字段

*  __Host__ - The OpsGenie API host. Use 'api.eu.opsgenie.com' for EU instances  
  默认值：`api.opsgenie.com`  
  URL 位置：<code class="service-url">opsgenie://<strong>host</strong>:port/apikey</code>  
*  __Port__ - The OpsGenie API port.  
  默认值：`443`  
  URL 位置：<code class="service-url">opsgenie://host:<strong>port</strong>/apikey</code>  
*  __APIKey__ - The OpsGenie API key （**必填**）  
  URL 位置：<code class="service-url">opsgenie://host:port/<strong>apikey</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Actions__ - Custom actions that will be available for the alert  
  默认值：*empty*  

*  __Alias__ - Client-defined identifier of the alert  
  默认值：*empty*  

*  __Description__ - Description field of the alert  
  默认值：*empty*  

*  __Details__ - Map of key-value pairs to use as custom properties of the alert  
  默认值：*empty*  

*  __Entity__ - Entity field of the alert that is generally used to specify which domain the Source field of the alert  
  默认值：*empty*  

*  __Note__ - Additional note that will be added while creating the alert  
  默认值：*empty*  

*  __Priority__ - Priority level of the alert. Possible values are P1, P2, P3, P4 and P5  
  默认值：*empty*  

*  __Responders__ - Teams, users, escalations and schedules that the alert will be routed to send notifications  
  默认值：*empty*  

*  __Source__ - Source field of the alert  
  默认值：*empty*  

*  __Tags__ - Tags of the alert  
  默认值：*empty*  

*  __Title__ - notification title, optionally set by the sender  
  默认值：*empty*  

*  __User__ - Display name of the request owner  
  默认值：*empty*  

*  __VisibleTo__ - Teams and users that the alert will become visible to without sending any notification  
  默认值：*empty*

## 在 OpsGenie 中创建 REST API 端点

1. 在菜单中点击 *Settings => Integration List*，打开集成列表页面
![Screenshot 1](./opsgenie/1.png)

2. 点击 *API => Add*

3. 确保勾选 *Create and Update Access* 和 *Enabled*，然后点击 *Save Integration*
![Screenshot 2](./opsgenie/2.png)

4. 复制 *API Key*

5. 拼装服务 URL

主机可以是 api.opsgenie.com 或 api.eu.opsgenie.com，取决于你的实例所在位置。详情见
[OpsGenie 文档](https://docs.opsgenie.com/docs/alert-api)。

```
opsgenie://api.opsgenie.com/eb243592-faa2-4ba2-a551q-1afdf565c889
                            └───────────────────────────────────┘
                                           token
```

## 通过代码传参

如果需要，你还可以向 `send` 函数传入额外参数。
<br/>
下面的示例包含当前支持的全部参数。

```go
service.Send("An example alert message", &types.Params{
    "alias":       "Life is too short for no alias",
    "description": "Every alert needs a description",
    "responders":  `[{"id":"4513b7ea-3b91-438f-b7e4-e3e54af9147c","type":"team"},{"name":"NOC","type":"team"}]`,
    "visibleTo":   `[{"id":"4513b7ea-3b91-438f-b7e4-e3e54af9147c","type":"team"},{"name":"rocket_team","type":"team"}]`,
    "actions":     "An action",
    "tags":        "tag1 tag2",
    "details":     `{"key1": "value1", "key2": "value2"}`,
    "entity":      "An example entity",
    "source":      "The source",
    "priority":    "P1",
    "user":        "Dracula",
    "note":        "Here is a note",
})
```

## 可选参数

你也可以在 URL 中指定这些参数：

:::info
opsgenie://api.opsgenie.com/eb243592-faa2-4ba2-a551q-1afdf565c889?alias=Life+is+too+short+for+no+alias&description=Every+alert+needs+a+description&actions=An+action&tags=["tag1","tag2"]&entity=An+example+entity&source=The+source&priority=P1&user=Dracula&note=Here+is+a+note
:::

命令行示例：

```shell
shoutrrr send -u 'opsgenie://api.eu.opsgenie.com/token?tags=["tag1","tag2"]&description=testing&responders=[{"username":"superuser", "type": "user"}]&entity=Example Entity&source=Example Source&actions=["asdf", "bcde"]' -m "Hello World6"
```

