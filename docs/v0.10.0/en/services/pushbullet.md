# Pushbullet

:::caution Upstream is barely maintained
The Pushbullet API keeps working, but the last significant API update was in June 2020 and the
service receives minimal upkeep. For new setups consider alternatives such as ntfy or Gotify.
:::

## URL Format

:::info
pushbullet://__`api-token`__[/__`device`__/#__`channel`__/__`email`__]
:::

### URL Fields

*  __Token__ (**Required**)  
  URL part: <code class="service-url">pushbullet://<strong>token</strong>/targets</code>  
*  __Targets__ (**Required**)  
  URL part: <code class="service-url">pushbullet://token/<strong>targets</strong></code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Title__  
  Default: `Shoutrrr notification`