# Pushbullet

:::caution 上游基本停止维护
Pushbullet API 仍可使用，但最后一次实质性 API 更新是 2020 年 6 月，服务几乎没有维护。
新项目建议考虑 ntfy、Gotify 等替代品。
:::

## URL 格式

:::info
pushbullet://__`api-token`__[/__`device`__/#__`channel`__/__`email`__]
:::

### URL 字段

*  __Token__ （**必填**）  
  URL 位置：<code class="service-url">pushbullet://<strong>token</strong>/targets</code>  
*  __Targets__ （**必填**）  
  URL 位置：<code class="service-url">pushbullet://token/<strong>targets</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Title__  
  默认值：`Shoutrrr notification`
