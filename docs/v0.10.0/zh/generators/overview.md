# 生成器

生成器用于通过命令行创建服务配置。
主生成器是基于反射的[基础生成器](./basic)，它通过一组简单的问答，力求为所有核心服务生成配置。
另外还有两个服务专属生成器：用于 Email 服务的 [OAuth2 生成器](./oauth2) 和 [Telegram 生成器](./telegram)。

## 用法

```bash
shoutrrr generate [OPTIONS] <SERVICE>
```
