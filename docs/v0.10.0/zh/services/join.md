# Join

## URL 格式

:::info
join://shoutrrr:__`api-key`__@join/?devices=__`device1`__[,__`device2`__, ...][&icon=__`icon`__][&title=__`title`__]
:::

### URL 字段

*  __APIKey__ （**必填**）  
  URL 位置：<code class="service-url">join://:<strong>apikey</strong>@join/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Devices__ - Comma separated list of device IDs （**必填**）  

*  __Icon__ - Icon URL  
  默认值：*empty*  

*  __Title__ - If set creates a notification  
  默认值：*empty*

## 操作指引

1.  打开 [Join Webapp](https://joinjoaomgcd.appspot.com/)
2.  选择你的设备
3.  点击 **Join API**
4.  你的 `deviceId` 会显示在页面顶部
5.  点击 `API Key` 旁边的 **Show** 查看你的密钥
6.  你的 Shoutrrr URL 将是：
    `join://shoutrrr:`__`api-key`__`@join/?devices=`__`deviceId`__

:::info
多个 `deviceId` 可以用 `,` 组合（重复步骤 2-4）。
:::
