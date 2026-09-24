# Email

## URL Format

:::info
smtp://__`username`__:__`password`__@__`host`__:__`port`__/?from=__`fromAddress`__&to=__`recipient1`__[,__`recipient2`__,...]
:::

### URL Fields

*  __Username__ - SMTP server username  
  Default: *empty*  
  URL part: <code class="service-url">smtp://<strong>username</strong>:password@host:port/</code>  
*  __Password__ - SMTP server password or hash (for OAuth2)  
  Default: *empty*  
  URL part: <code class="service-url">smtp://username:<strong>password</strong>@host:port/</code>  
*  __Host__ - SMTP server hostname or IP address (**Required**)  
  URL part: <code class="service-url">smtp://username:password@<strong>host</strong>:port/</code>  
*  __Port__ - SMTP server port, common ones are 25, 465, 587 or 2525  
  Default: `25`  
  URL part: <code class="service-url">smtp://username:password@host:<strong>port</strong>/</code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __FromAddress__ - E-mail address that the mail are sent from (**Required**)  
  Aliases: `from`  

*  __ToAddresses__ - List of recipient e-mails separated by "," (comma) (**Required**)  
  Aliases: `to`  

*  __Auth__ - SMTP authentication method  
  Default: `Unknown`  
  Possible values: `None`, `Plain`, `CRAMMD5`, `Unknown`, `OAuth2`  

*  __ClientHost__ - The client host name sent to the SMTP server during HELLO phase. If set to "auto" it will use the OS hostname  
  Default: `localhost`  

*  __Encryption__ - Encryption method  
  Default: `Auto`  
  Possible values: `None`, `ExplicitTLS`, `ImplicitTLS`, `Auto`  

*  __FromName__ - Name of the sender  
  Default: *empty*  

*  __Subject__ - The subject of the sent mail  
  Default: `Shoutrrr Notification`  
  Aliases: `title`  

*  __UseHTML__ - Whether the message being sent is in HTML  
  Default: ❌ `No`  

*  __UseStartTLS__ - Whether to use StartTLS encryption  
  Default: ✔ `Yes`  
  Aliases: `starttls`
