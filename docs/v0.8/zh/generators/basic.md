# 基础生成器

基础生成器会查看服务配置结构体上的 `key:""`、`desc:""` 和 `default:""` 标签，并利用它们提示用户填写对应的值。

示例：
```shell
$ shoutrrr generate telegram
```
```yaml
Generating URL for telegram using basic generator
Enter the configuration values as prompted

Token: 110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw
Preview[Yes]: No
Notification[Yes]:
ParseMode[None]:
Channels: @mychannel

URL: telegram://110201543:AAHdqTcvCH1vGWJxfSeofSAs0K5PALDsaw@telegram?channels=@mychannel&notification=Yes&parsemode=None&preview=No
```
