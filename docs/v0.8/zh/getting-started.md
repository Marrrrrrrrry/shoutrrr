# 快速开始

## 作为依赖包

使用 shoutrrr 非常简单！目前有两种作为依赖包使用的方式。

### 直接发送

最简单，但功能非常有限。

```go
url := "slack://token-a/token-b/token-c"
err := shoutrrr.Send(url, "Hello world (or slack channel) !")
```

### 使用 Sender

使用 Sender 可以预先配置多个通知服务，并用同一个 `Send(message, params)` 方法向它们全部发送。

```go
urlA := "slack://token-a/token-b/token-c"
urlB := "telegram://110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw@telegram?channels=@mychannel"
sender, err := shoutrrr.CreateSender(urlA, urlB)

// Send notifications instantly to all services
sender.Send("Hello world (or slack/telegram channel)!", map[string]string { "title": "He-hey~!"  })

// ...or bundle notifications... 
func doWork() error {
    // ...and send them when leaving the scope
    defer sender.Flush(map[string]string { "title": "Work Result" })
    
    sender.Enqueue("Started doing %v", stuff)
    
    // Maybe get creative...?
    defer func(start time.Time) { 
    	sender.Enqueue("Elapsed: %v", time.Now().Sub(start)) 
    }(time.Now())
    
    if err := doMoreWork(); err != nil {
        sender.Enqueue("Oh no! %v", err)
    	
        // This will send the currently queued up messages...
        return
    }   
    
    sender.Enqueue("Everything went very well!")
    
    // ...or this:
}

```


## 通过命令行

先运行 `build.sh` 脚本，然后就可以运行 shoutrrr 可执行文件：

```shell
$ ./shoutrrr

Usage:
./shoutrrr <ActionVerb> [...]
Possible actions: send, verify, generate
```

在装有 Go 的系统上，可以用以下命令安装最新的 Shoutrrr CLI：

```shell
go install github.com/marrrrrrrrry/shoutrrr/shoutrrr@latest
```

### 命令

#### Send

使用给定的通知服务 URL 发送一条通知。

```shell
go install github.com/marrrrrrrrry/shoutrrr/shoutrrr@latest
```

#### Verify

校验通知服务 URL 的有效性。

```bash
$ shoutrrr send \
    --url "<SERVICE_URL>" \
    --message "<MESSAGE BODY>"
```

#### Generate

生成并显示通知服务 URL 的配置。

```bash
$ shoutrrr verify \
    --url "<SERVICE_URL>"
```

| 参数                         | 说明                                           |
| ---------------------------- | ---------------------------------------------- |
| `-g, --generator string`     | 使用的生成器（默认 "basic"）                   |
| `-p, --property stringArray` | key=value 格式的配置项                         |
| `-s, --service string`       | 要生成 URL 的通知服务                          |

**注意**：服务既可以作为第一个参数传入，也可以用 `-s` 参数指定。

关于生成器的更多信息，参见[生成器](/zh/generators/overview)。

### 选项

#### Debug

开启 CLI 的调试输出。

| 参数            | 环境变量         | 默认值  | 必填 |
| --------------- | ---------------- | ------- | ---- |
| `--debug`, `-d` | `SHOUTRRR_DEBUG` | `false` |      |

#### URL

通知的目标 URL，参见[服务概览](/zh/services/overview)。

| 参数          | 环境变量       | 默认值 | 必填 |
| ------------- | -------------- | ------ | ---- |
| `--url`, `-u` | `SHOUTRRR_URL` | 无     | ✅   |

## 在 GitHub Actions 工作流中使用

你也可以在 GitHub Actions 工作流中使用 Shoutrrr。

参考下面的示例以及 [GitHub Marketplace 上的 action](https://github.com/marketplace/actions/shoutrrr-action)：

```bash
$ shoutrrr generate [OPTIONS] <SERVICE>
```
