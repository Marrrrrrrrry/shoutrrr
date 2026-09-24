# Bark

上游文档：https://github.com/Finb/Bark

## URL 格式

### URL 字段

*  __DeviceKey__ - The key for each device （**必填**）  
  URL 位置：<code class="service-url">bark://:<strong>devicekey</strong>@host/path</code>  
*  __Host__ - Server hostname and port （**必填**）  
  URL 位置：<code class="service-url">bark://:devicekey@<strong>host</strong>/path</code>  
*  __Path__ - Server path  
  默认值：`/`  
  URL 位置：<code class="service-url">bark://:devicekey@host/<strong>path</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Badge__ - The number displayed next to App icon  
  默认值：`0`  

*  __Category__ - Reserved field, no use yet  
  默认值：*empty*  

*  __Copy__ - The value to be copied  
  默认值：*empty*  

*  __Group__ - The group of the notification  
  默认值：*empty*  

*  __Icon__ - An url to the icon, available only on iOS 15 or later  
  默认值：*empty*  

*  __Scheme__ - Server protocol, http or https  
  默认值：`https`  

*  __Sound__ - Value from https://github.com/Finb/Bark/tree/master/Sounds  
  默认值：*empty*  

*  __Title__ - Notification title, optionally set by the sender  
  默认值：*empty*  

*  __URL__ - Url that will jump when click notification  
  默认值：*empty*
