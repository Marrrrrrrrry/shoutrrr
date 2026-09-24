# Join

## URL Format

:::info
join://shoutrrr:__`api-key`__@join/?devices=__`device1`__[,__`device2`__, ...][&icon=__`icon`__][&title=__`title`__]
:::

### URL Fields

*  __APIKey__ (**Required**)  
  URL part: <code class="service-url">join://:<strong>apikey</strong>@join/</code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Devices__ - Comma separated list of device IDs (**Required**)  

*  __Icon__ - Icon URL  
  Default: *empty*  

*  __Title__ - If set creates a notification  
  Default: *empty*

## Guide

1.  Go to the [Join Webapp](https://joinjoaomgcd.appspot.com/) 
2.  Select your device
3.  Click **Join API** 
4.  Your `deviceId` is shown in the top
5.  Click **Show** next to `API Key` to see your key 
6.  Your Shoutrrr URL will then be:
    `join://shoutrrr:`__`api-key`__`@join/?devices=`__`deviceId`__

:::info
Multiple `deviceId`s can be combined with a `,` (repeat steps 2-4).
:::