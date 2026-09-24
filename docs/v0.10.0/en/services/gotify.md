# Gotify

## URL Format

### URL Fields

*  __Host__ - Server hostname (and optionally port) (**Required**)  
  URL part: <code class="service-url">gotify://<strong>host</strong>:<strong>port</strong>/path/token</code>  
*  __Path__ - Server subpath  
  Default: *empty*  
  URL part: <code class="service-url">gotify://host:port/<strong>path</strong>/token</code>  
*  __Token__ - Application token (**Required**)  
  URL part: <code class="service-url">gotify://host:port/path/<strong>token</strong></code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __DisableTLS__  
  Default: ❌ `No`  

*  __Priority__  
  Default: `0`  

*  __Title__  
  Default: `Shoutrrr notification`

## Examples

:::tip Common usage

```text
gotify://gotify.example.com:443/AzyoeNS.D4iJLVa/?title=Great+News&priority=1
```
:::

:::tip With subpath
```text
gotify://example.com:443/path/to/gotify/AzyoeNS.D4iJLVa/?title=Great+News&priority=1
```
:::