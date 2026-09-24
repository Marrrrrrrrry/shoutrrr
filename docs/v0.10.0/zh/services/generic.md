# Generic

Generic 服务可用于任何 Shoutrrr 未显式支持的目标，只要它支持通过 POST 请求接收消息。
通常这需要在接收端做一些定制来解析收到的负载，因此可能并不可行。

与常见服务提供商配合使用的示例见[示例页](/zh/examples/generic)。

## 自定义请求头
你可以在 URL 中添加以 `@` 为前缀的查询变量（`@key=value`）来附加自定义 HTTP 请求头。

使用
```text
generic://example.com?@acceptLanguage=tlh-Piqd
```
会额外添加如下请求头：

```http
Accept-Language: tlh-Piqd
```

## JSON 模板
通过内置的 `JSON` 模板（`template=json`）可以构造通用的 JSON 负载。`title` 和 `message` 所用的键可以通过 params/查询参数 `titleKey` 和 `messageKey` 覆盖。

:::tip 示例
```json
{
    "title": "Oh no!",
    "message": "The thing happened and now there is stuff all over the area!"
}
```
:::

### 自定义数据字段
使用 JSON 模板时，可以通过添加以 `$` 为前缀的查询变量（`$key=value`）向 JSON 对象附加额外的键值对。

:::tip 示例
使用 `generic://example.com?$projection=retroazimuthal` 将得到：

```json
{
    "title": "Amazing opportunities!",
    "message": "New map book available for purchase.",
    "projection": "retroazimuthal"
}
```
:::

## 快捷 URL
只需在目标 URL 前加 `generic+` 前缀即可使用 generic 服务，所以
```text
https://example.com/api/v1/postStuff
```
会变成
```text
generic+https://example.com/api/v1/postStuff
```

:::info
URL 上添加的任何查询变量都会被转义后原样转发给远端服务器。这意味着 `generic+https://` 无法使用 `?template=json`，请改用 `generic://`！
:::

## 转发的查询变量
所有未在[查询参数](#查询参数)一节中列出的查询变量都会被转发到目标端点。
如果你需要传递一个_被占用_的查询变量名，可以给它加下划线前缀（`_`）。

:::tip 示例
URL `generic+https://example.com/api/v1/postStuff?contenttype=text/plain` 会向 `https://example.com/api/v1/postStuff` 发送 POST 消息，并带上 `Content-Type: text/plain` 请求头。

而转义写法 `generic+https://example.com/api/v1/postStuff?_contenttype=text/plain` 会向 `https://example.com/api/v1/postStuff?contenttype=text/plain` 发送 POST 消息，并使用默认的 `Content-Type: application/json` 请求头。
:::


## URL 格式

### URL 字段

### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __ContentType__ - The value of the Content-Type header  
  默认值：`application/json`  

*  __DisableTLS__  
  默认值：❌ `No`  

*  __MessageKey__ - The key that will be used for the message value  
  默认值：`message`  

*  __RequestMethod__  
  默认值：`POST`  

*  __Template__ - The template used for creating the request payload  
  默认值：*empty*  

*  __Title__  
  默认值：*empty*  

*  __TitleKey__ - The key that will be used for the title value  
  默认值：`title`
