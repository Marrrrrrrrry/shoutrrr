# 示例

以下是配合常见服务提供商、使用 [generic 服务](/zh/services/generic)的服务 URL 示例。

## Home Assistant

服务 URL 需要写成：
```
generic://HAIPAddress:HAPort/api/webhook/WebhookIDFromHA?template=json
```

如果需要 http://：
```
generic://HAIPAddress:HAPort/api/webhook/WebhookIDFromHA?template=json&disabletls=yes
```

然后，在 HA 中使用 `{{ trigger.json.message }}` 获取从 JSON 传来的消息。

_致谢 [@JeffCrum1](https://github.com/JeffCrum1)，来源：[https://github.com/containrrr/shoutrrr/issues/325#issuecomment-1460105065](https://github.com/containrrr/shoutrrr/issues/325#issuecomment-1460105065)_

## Apprise

服务 URL 需要写成：

```
generic://apprise-url/notify/devops?template=json&messagekey=body&title=title
```

如果需要 http://：
```
generic://apprise-url/notify/devops?template=json&messagekey=body&title=title&disabletls=yes
```
