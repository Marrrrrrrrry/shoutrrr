# Generic
The Generic service can be used for any target that is not explicitly supported by Shoutrrr, as long as it
supports receiving the message via a POST request.
Usually, this requires customization on the receiving end to interpret the payload that it receives, and might
not be a viable approach.

Common examples for use with service providers can be found under [examples](/examples/generic).

## Custom headers
You can add additional HTTP headers to your request by adding query variables prefixed with `@` (`@key=value`).

Using 
```text
generic://example.com?@acceptLanguage=tlh-Piqd
```
would result in the additional header being added:

```http
Accept-Language: tlh-Piqd
```

## JSON template
By using the built in `JSON` template (`template=json`) you can create a generic JSON payload. The keys used for `title` and `message` can be overriden
by supplying the params/query values `titleKey` and `messageKey`.

:::tip Example
```json
{
    "title": "Oh no!",
    "message": "The thing happened and now there is stuff all over the area!"
}
```
:::

### Custom data fields
When using the JSON template, you can add additional key/value pairs to the JSON object by adding query variables prefixed with `$` (`$key=value`).

:::tip Example
Using `generic://example.com?$projection=retroazimuthal` would yield: 

```json
{
    "title": "Amazing opportunities!",
    "message": "New map book available for purchase.",
    "projection": "retroazimuthal"
}
```
:::

## Shortcut URL
You can just add `generic+` as a prefix to your target URL to use it with the generic service, so
```text
https://example.com/api/v1/postStuff
```
would become
```text
generic+https://example.com/api/v1/postStuff
```

:::info
Any query variables added to the URL will be escaped so that they can be forwarded to the remote server. That means that you cannot use `?template=json` with the  `generic+https://`, just use `generic://` instead!
:::

## Forwarded query variables
All query variables that are not listed in the [Query/Param Props](#queryparam_props) section will be
forwarded to the target endpoint.
If you need to pass a query variable that _is_ reserved, you can prefix it with an underscore (`_`).

:::tip Example
The URL `generic+https://example.com/api/v1/postStuff?contenttype=text/plain` would send a POST message
to `https://example.com/api/v1/postStuff` using the `Content-Type: text/plain` header.

If instead escaped, `generic+https://example.com/api/v1/postStuff?_contenttype=text/plain` would send a POST message
to `https://example.com/api/v1/postStuff?contenttype=text/plain` using the `Content-Type: application/json` header (as it's the default).
:::


## URL Format

### URL Fields

### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __ContentType__ - The value of the Content-Type header  
  Default: `application/json`  

*  __DisableTLS__  
  Default: ❌ `No`  

*  __MessageKey__ - The key that will be used for the message value  
  Default: `message`  

*  __RequestMethod__  
  Default: `POST`  

*  __Template__ - The template used for creating the request payload  
  Default: *empty*  

*  __Title__  
  Default: *empty*  

*  __TitleKey__ - The key that will be used for the title value  
  Default: `title`
