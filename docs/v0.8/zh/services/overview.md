# 服务概览

点击服务名查看更详细的说明。 <!-- @formatter:off -->

| 服务                              | URL 格式                                                                                                                                        |
| --------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| [Bark](/zh/services/bark)                 | *bark://__`devicekey`__@__`host`__*                                                                                                             |
| [Discord](/zh/services/discord)           | *discord://__`token`__@__`id`__*                                                                                                                |
| [Email](/zh/services/email)               | *smtp://__`username`__:__`password`__@__`host`__:__`port`__/?from=__`fromAddress`__&to=__`recipient1`__[,__`recipient2`__,...]*                 |
| [Gotify](/zh/services/gotify)             | *gotify://__`gotify-host`__/__`token`__*                                                                                                        |
| [Google Chat](/zh/services/googlechat)    | *googlechat://chat.googleapis.com/v1/spaces/FOO/messages?key=bar&token=baz*                                                                     |
| [IFTTT](/zh/services/ifttt)               | *ifttt://__`key`__/?events=__`event1`__[,__`event2`__,...]&value1=__`value1`__&value2=__`value2`__&value3=__`value3`__*                         |
| [Join](/zh/services/join)                 | *join://shoutrrr:__`api-key`__@join/?devices=__`device1`__[,__`device2`__, ...][&icon=__`icon`__][&title=__`title`__]*                          |
| [Mattermost](/zh/services/mattermost)     | *mattermost://[__`username`__@]__`mattermost-host`__/__`token`__[/__`channel`__]*                                                               |
| [Matrix](/zh/services/matrix)             | *matrix://__`username`__:__`password`__@__`host`__:__`port`__/[?rooms=__`!roomID1`__[,__`roomAlias2`__]]*                                       |
| [Ntfy](/zh/services/ntfy)                 | *ntfy://__`username`__:__`password`__@ntfy.sh/__`topic`__*                                                                                      |
| [OpsGenie](/zh/services/opsgenie)         | *opsgenie://__`host`__/token?responders=__`responder1`__[,__`responder2`__]*                                                                    |
| [Pushbullet](/zh/services/pushbullet)     | *pushbullet://__`api-token`__[/__`device`__/#__`channel`__/__`email`__]*                                                                        |
| [Pushover](/zh/services/pushover)         | *pushover://shoutrrr:__`apiToken`__@__`userKey`__/?devices=__`device1`__[,__`device2`__, ...]*                                                  |
| [Rocketchat](/zh/services/rocketchat)     | *rocketchat://[__`username`__@]__`rocketchat-host`__/__`token`__[/__`channel`&#124;`@recipient`__]*                                             |
| [Slack](/zh/services/slack)               | *slack://[__`botname`__@]__`token-a`__/__`token-b`__/__`token-c`__*                                                                             |
| [Teams](/zh/services/teams)               | *teams://__`group`__@__`tenant`__/__`altId`__/__`groupOwner`__?host=__`organization`__.webhook.office.com*                                      |
| [Telegram](/zh/services/telegram)         | *telegram://__`token`__@telegram?chats=__`@channel-1`__[,__`chat-id-1`__,...]*                                                                  |
| [Zulip Chat](/zh/services/zulip)          | *zulip://__`bot-mail`__:__`bot-key`__@__`zulip-domain`__/?stream=__`name-or-id`__&topic=__`name`__*                                             |

## 特殊服务

| 服务                              | 说明                                                                                                                                            |
| --------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| [Logger](/zh/services/logger)             | 将通知写入所配置的 Go `log.Logger`                                                                                                              |
| [Generic Webhook](/zh/services/generic)   | 直接向 Webhook 发送通知                                                                                                                         |

