# Bark

Upstream docs: https://github.com/Finb/Bark

## URL Format

### URL Fields

*  __DeviceKey__ - The key for each device (**Required**)  
  URL part: <code class="service-url">bark://:<strong>devicekey</strong>@host/path</code>  
*  __Host__ - Server hostname and port (**Required**)  
  URL part: <code class="service-url">bark://:devicekey@<strong>host</strong>/path</code>  
*  __Path__ - Server path  
  Default: `/`  
  URL part: <code class="service-url">bark://:devicekey@host/<strong>path</strong></code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Badge__ - The number displayed next to App icon  
  Default: `0`  

*  __Category__ - Reserved field, no use yet  
  Default: *empty*  

*  __Copy__ - The value to be copied  
  Default: *empty*  

*  __Group__ - The group of the notification  
  Default: *empty*  

*  __Icon__ - An url to the icon, available only on iOS 15 or later  
  Default: *empty*  

*  __Scheme__ - Server protocol, http or https  
  Default: `https`  

*  __Sound__ - Value from https://github.com/Finb/Bark/tree/master/Sounds  
  Default: *empty*  

*  __Title__ - Notification title, optionally set by the sender  
  Default: *empty*  

*  __URL__ - Url that will jump when click notification  
  Default: *empty*