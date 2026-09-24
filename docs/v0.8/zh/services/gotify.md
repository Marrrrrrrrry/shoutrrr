# Gotify

## URL 格式

### URL 字段

*  __Host__ - Server hostname (and optionally port) （**必填**）  
  URL 位置：<code class="service-url">gotify://<strong>host</strong>:<strong>port</strong>/path/token</code>  
*  __Path__ - Server subpath  
  默认值：*empty*  
  URL 位置：<code class="service-url">gotify://host:port/<strong>path</strong>/token</code>  
*  __Token__ - Application token （**必填**）  
  URL 位置：<code class="service-url">gotify://host:port/path/<strong>token</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __DisableTLS__  
  默认值：❌ `No`  

*  __Priority__  
  默认值：`0`  

*  __Title__  
  默认值：`Shoutrrr notification`

## 示例

:::tip 常见用法

```text
gotify://gotify.example.com:443/AzyoeNS.D4iJLVa/?title=Great+News&priority=1
```
:::

:::tip 带子路径
```text
gotify://example.com:443/path/to/gotify/AzyoeNS.D4iJLVa/?title=Great+News&priority=1
```
:::
