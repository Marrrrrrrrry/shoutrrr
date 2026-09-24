# IFTTT

:::caution Not available on the free tier
Since February 2024, the IFTTT Webhooks service is no longer available on free accounts;
using it requires a paid Pro/Pro+ plan.
:::

## URL Format

:::info
ifttt://__`key`__/?events=__`event1`__[,__`event2`__,...]&value1=__`value1`__&value2=__`value2`__&value3=__`value3`__
:::

### URL Fields

*  __WebHookID__ (**Required**)  
  URL part: <code class="service-url">ifttt://<strong>webhookid</strong>/</code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Events__ (**Required**)  

*  __Title__ - Notification title, optionally set by the sender  
  Default: *empty*  

*  __UseMessageAsValue__ - Sets the corresponding value field to the notification message  
  Default: `2`  

*  __UseTitleAsValue__ - Sets the corresponding value field to the notification title  
  Default: `0`  

*  __Value1__  
  Default: *empty*  

*  __Value2__  
  Default: *empty*  

*  __Value3__  
  Default: *empty*