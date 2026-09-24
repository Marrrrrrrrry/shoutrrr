# MatterMost

## URL 格式

:::info
mattermost://[__`username`__@]__`mattermost-host`__/__`token`__[/__`channel`__][?icon=__`smiley`__]
:::

### URL 字段

*  __UserName__ - Override webhook user  
  默认值：*empty*  
  URL 位置：<code class="service-url">mattermost://<strong>username</strong>@host:port/token/channel</code>  
*  __Host__ - Mattermost server host （**必填**）  
  URL 位置：<code class="service-url">mattermost://username@<strong>host</strong>:<strong>port</strong>/token/channel</code>  
*  __Token__ - Webhook token （**必填**）  
  URL 位置：<code class="service-url">mattermost://username@host:port/<strong>token</strong>/channel</code>  
*  __Channel__ - Override webhook channel  
  默认值：*empty*  
  URL 位置：<code class="service-url">mattermost://username@host:port/token/<strong>channel</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Icon__ - Use emoji or URL as icon (based on presence of http(s):// prefix)  
  默认值：*empty*  
  别名：`icon_emoji`, `icon_url`  

*  __Title__ - Notification title, optionally set by the sender (not used)  
  默认值：*empty*

## 在 MatterMost 中创建 Webhook

1. 点击菜单中的 *Integrations*，打开集成页面
![Screenshot 1](./mattermost/1.PNG)

2. 点击 *Incoming Webhooks*
![Screenshot 2](./mattermost/2.PNG)

3. 点击 *Add Incoming Webhook*
![Screenshot 3](./mattermost/3.PNG)

4. 填写 Webhook 信息并点击 *Save*
![Screenshot 4](./mattermost/4.PNG)

5. 如果操作正确，MatterMost 会给你新建 Webhook 的 *URL*
![Screenshot 5](./mattermost/5.PNG)

6. 拼装服务 URL
```
https://your-domain.com/hooks/bywsw8zt5jgpte3nm65qjiru6h
                              └────────────────────────┘
                                        token
mattermost://your-domain.com/bywsw8zt5jgpte3nm65qjiru6h
                             └────────────────────────┘
                                       token
```
