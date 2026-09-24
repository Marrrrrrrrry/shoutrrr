# OpsGenie

## URL Format

### URL Fields

*  __Host__ - The OpsGenie API host. Use 'api.eu.opsgenie.com' for EU instances  
  Default: `api.opsgenie.com`  
  URL part: <code class="service-url">opsgenie://<strong>host</strong>:port/apikey</code>  
*  __Port__ - The OpsGenie API port.  
  Default: `443`  
  URL part: <code class="service-url">opsgenie://host:<strong>port</strong>/apikey</code>  
*  __APIKey__ - The OpsGenie API key (**Required**)  
  URL part: <code class="service-url">opsgenie://host:port/<strong>apikey</strong></code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Actions__ - Custom actions that will be available for the alert  
  Default: *empty*  

*  __Alias__ - Client-defined identifier of the alert  
  Default: *empty*  

*  __Description__ - Description field of the alert  
  Default: *empty*  

*  __Details__ - Map of key-value pairs to use as custom properties of the alert  
  Default: *empty*  

*  __Entity__ - Entity field of the alert that is generally used to specify which domain the Source field of the alert  
  Default: *empty*  

*  __Note__ - Additional note that will be added while creating the alert  
  Default: *empty*  

*  __Priority__ - Priority level of the alert. Possible values are P1, P2, P3, P4 and P5  
  Default: *empty*  

*  __Responders__ - Teams, users, escalations and schedules that the alert will be routed to send notifications  
  Default: *empty*  

*  __Source__ - Source field of the alert  
  Default: *empty*  

*  __Tags__ - Tags of the alert  
  Default: *empty*  

*  __Title__ - notification title, optionally set by the sender  
  Default: *empty*  

*  __User__ - Display name of the request owner  
  Default: *empty*  

*  __VisibleTo__ - Teams and users that the alert will become visible to without sending any notification  
  Default: *empty*

## Creating a REST API endpoint in OpsGenie

1. Open up the Integration List page by clicking on *Settings => Integration List* within the menu
![Screenshot 1](./opsgenie/1.png)

2. Click *API => Add*

3. Make sure *Create and Update Access* and *Enabled* are checked and click *Save Integration*
![Screenshot 2](./opsgenie/2.png)

4. Copy the *API Key*

5. Format the service URL

The host can be either api.opsgenie.com or api.eu.opsgenie.com depending on the location of your instance. See
the [OpsGenie documentation](https://docs.opsgenie.com/docs/alert-api) for details.

```
opsgenie://api.opsgenie.com/eb243592-faa2-4ba2-a551q-1afdf565c889
                            └───────────────────────────────────┘
                                           token
```

## Passing parameters via code

If you want to, you can pass additional parameters to the `send` function.
<br/>
The following example contains all parameters that are currently supported.

```go
service.Send("An example alert message", &types.Params{
    "alias":       "Life is too short for no alias",
    "description": "Every alert needs a description",
    "responders":  `[{"id":"4513b7ea-3b91-438f-b7e4-e3e54af9147c","type":"team"},{"name":"NOC","type":"team"}]`,
    "visibleTo":   `[{"id":"4513b7ea-3b91-438f-b7e4-e3e54af9147c","type":"team"},{"name":"rocket_team","type":"team"}]`,
    "actions":     "An action",
    "tags":        "tag1 tag2",
    "details":     `{"key1": "value1", "key2": "value2"}`,
    "entity":      "An example entity",
    "source":      "The source",
    "priority":    "P1",
    "user":        "Dracula",
    "note":        "Here is a note",
})
```

## Optional parameters

You can optionally specify the parameters in the URL:

:::info
opsgenie://api.opsgenie.com/eb243592-faa2-4ba2-a551q-1afdf565c889?alias=Life+is+too+short+for+no+alias&description=Every+alert+needs+a+description&actions=An+action&tags=["tag1","tag2"]&entity=An+example+entity&source=The+source&priority=P1&user=Dracula&note=Here+is+a+note
:::

Example using the command line:

```shell
shoutrrr send -u 'opsgenie://api.eu.opsgenie.com/token?tags=["tag1","tag2"]&description=testing&responders=[{"username":"superuser", "type": "user"}]&entity=Example Entity&source=Example Source&actions=["asdf", "bcde"]' -m "Hello World6"
```


