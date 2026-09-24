# Email

:::caution Gmail 必须使用 OAuth
Google 已对 Google Workspace 账户关闭基于密码（Basic auth）的 IMAP/SMTP 访问，
应用专用密码也在逐步淘汰。Gmail 用户请使用 OAuth2 认证类型配合访问令牌
（通过 [`oauth2` 生成器](/zh/generators/oauth2) 获取）。SMTP 的 XOAUTH2 仍被官方支持；
参见 [Google XOAUTH2 文档](https://developers.google.com/workspace/gmail/imap/xoauth2-protocol)。
:::

## URL 格式

:::info
smtp://__`username`__:__`password`__@__`host`__:__`port`__/?from=__`fromAddress`__&to=__`recipient1`__[,__`recipient2`__,...]
:::

### URL 字段

*  __Username__ - SMTP server username  
  默认值：*empty*  
  URL 位置：<code class="service-url">smtp://<strong>username</strong>:password@host:port/</code>  
*  __Password__ - SMTP server password or hash (for OAuth2)  
  默认值：*empty*  
  URL 位置：<code class="service-url">smtp://username:<strong>password</strong>@host:port/</code>  
*  __Host__ - SMTP server hostname or IP address （**必填**）  
  URL 位置：<code class="service-url">smtp://username:password@<strong>host</strong>:port/</code>  
*  __Port__ - SMTP server port, common ones are 25, 465, 587 or 2525  
  默认值：`25`  
  URL 位置：<code class="service-url">smtp://username:password@host:<strong>port</strong>/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __FromAddress__ - E-mail address that the mail are sent from （**必填**）  
  别名：`from`  

*  __ToAddresses__ - List of recipient e-mails separated by "," (comma) （**必填**）  
  别名：`to`  

*  __Auth__ - SMTP authentication method  
  默认值：`Unknown`  
  可选值：`None`, `Plain`, `CRAMMD5`, `Unknown`, `OAuth2`  

*  __ClientHost__ - The client host name sent to the SMTP server during HELLO phase. If set to "auto" it will use the OS hostname  
  默认值：`localhost`  

*  __Encryption__ - Encryption method  
  默认值：`Auto`  
  可选值：`None`, `ExplicitTLS`, `ImplicitTLS`, `Auto`  

*  __FromName__ - Name of the sender  
  默认值：*empty*  

*  __Subject__ - The subject of the sent mail  
  默认值：`Shoutrrr Notification`  
  别名：`title`  

*  __UseHTML__ - Whether the message being sent is in HTML  
  默认值：❌ `No`  

*  __UseStartTLS__ - Whether to use StartTLS encryption  
  默认值：✔ `Yes`  
  别名：`starttls`
