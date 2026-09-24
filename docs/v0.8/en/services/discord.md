# Discord

## URL Format

Your Discord Webhook-URL will look like this:

:::info
<span>https://</span>discord.com/api/webhooks/__`webhookid`__/__`token`__  
:::

The shoutrrr service URL should look like this:  

:::info
discord://__`token`__@__`webhookid`__
:::

### URL Fields

*  __Token__ (**Required**)  
  URL part: <code class="service-url">discord://<strong>token</strong>@webhookid/</code>  
*  __WebhookID__ (**Required**)  
  URL part: <code class="service-url">discord://token@<strong>webhookid</strong>/</code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Avatar__ - Override the webhook default avatar with specified URL  
  Default: *empty*  
  Aliases: `avatarurl`  

*  __Color__ - The color of the left border for plain messages  
  Default: `0x50D9ff`  

*  __ColorDebug__ - The color of the left border for debug messages  
  Default: `0x7b00ab`  

*  __ColorError__ - The color of the left border for error messages  
  Default: `0xd60510`  

*  __ColorInfo__ - The color of the left border for info messages  
  Default: `0x2488ff`  

*  __ColorWarn__ - The color of the left border for warning messages  
  Default: `0xffc441`  

*  __JSON__ - Whether to send the whole message as the JSON payload instead of using it as the 'content' field  
  Default: ❌ `No`  

*  __SplitLines__ - Whether to send each line as a separate embedded item  
  Default: ✔ `Yes`  

*  __ThreadID__ - Optional thread ID for posting into a channel thread  
  Default: *empty*  

*  __Title__  
  Default: *empty*  

*  __Username__ - Override the webhook default username  
  Default: *empty*

## Optional: Sending messages to a specific thread

Discord supports sending messages via webhook to threads. You can target a specific thread by appending
`?thread_id=<id>` to the end of the url.

:::info
### Example
:::

```
discord://<token>@<webhook_id>?thread_id=1234567890123456789
```

To get the thread ID:
 - Open a Discord Thread
 - Right click -> copy link
 - OR Right click -> copy thread ID (Developer mode turned on in discord settings)

:::warning
A valid thread_id must is 19 digits long. If extracting from a link don't confuse the channel ID and thread ID.
:::

## Creating a webhook in Discord

1. Open your channel settings by first clicking on the gear icon next to the name of the channel.
![Screenshot 1](./discord/sc-1.png)

2. In the menu on the left, click on *Integrations*.
![Screenshot 2](./discord/sc-2.png)

3. In the menu on the right, click on *Create Webhook*.
![Screenshot 3](./discord/sc-3.png)

4. Set the name, channel and icon to your liking and click the *Copy Webhook URL* button.
![Screenshot 4](./discord/sc-4.png)

5. Press the *Save Changes* button.
![Screenshot 5](./discord/sc-5.png)

6. Format the service URL:
```
https://discord.com/api/webhooks/693853386302554172/W3dE2OZz4C13_4z_uHfDOoC7BqTW288s-z1ykqI0iJnY_HjRqMGO8Sc7YDqvf_KVKjhJ
                                 └────────────────┘ └──────────────────────────────────────────────────────────────────┘
                                     webhook id                                    token

discord://W3dE2OZz4C13_4z_uHfDOoC7BqTW288s-z1ykqI0iJnY_HjRqMGO8Sc7YDqvf_KVKjhJ@693853386302554172
          └──────────────────────────────────────────────────────────────────┘ └────────────────┘
                                          token                                    webhook id
```
