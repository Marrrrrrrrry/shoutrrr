# IFTTT

## URL 格式

:::info
ifttt://__`key`__/?events=__`event1`__[,__`event2`__,...]&value1=__`value1`__&value2=__`value2`__&value3=__`value3`__
:::

### URL 字段

*  __WebHookID__ （**必填**）  
  URL 位置：<code class="service-url">ifttt://<strong>webhookid</strong>/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Events__ （**必填**）  

*  __Title__ - Notification title, optionally set by the sender  
  默认值：*empty*  

*  __UseMessageAsValue__ - Sets the corresponding value field to the notification message  
  默认值：`2`  

*  __UseTitleAsValue__ - Sets the corresponding value field to the notification title  
  默认值：`0`  

*  __Value1__  
  默认值：*empty*  

*  __Value2__  
  默认值：*empty*  

*  __Value3__  
  默认值：*empty*
