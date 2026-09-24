# Teams

:::caution Webhook URL scheme changed
Microsoft has changed the URL scheme for Teams webhooks. You will now have to specify the hostname using:
```text
?host=example.webhook.office.com
```
Where `example` is your organization short name
:::

## URL Format

:::info
teams://__`group`__@__`tenant`__/__`altId`__/__`groupOwner`__?host=__`organization`__.webhook.office.com
:::

### URL Fields

*  __Group__  
  Default: *empty*  
  URL part: <code class="service-url">teams://<strong>group</strong>@tenant/altid/groupowner</code>  
*  __Tenant__  
  Default: *empty*  
  URL part: <code class="service-url">teams://group@<strong>tenant</strong>/altid/groupowner</code>  
*  __AltID__  
  Default: *empty*  
  URL part: <code class="service-url">teams://group@tenant/<strong>altid</strong>/groupowner</code>  
*  __GroupOwner__  
  Default: *empty*  
  URL part: <code class="service-url">teams://group@tenant/altid/<strong>groupowner</strong></code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Color__  
  Default: *empty*  

*  __Host__  
  Default: `outlook.office.com`  

*  __Title__  
  Default: *empty*

## Setting up a webhook

To be able to use the Microsoft Teams notification service, you first need to set up a custom webhook.
Instructions on how to do this can be found in [this guide](https://docs.microsoft.com/en-us/microsoftteams/platform/webhooks-and-connectors/how-to/connectors-using#setting-up-a-custom-incoming-webhook)

## Extracting the token

The token is extracted from your webhook URL:

<pre><code>https://<b>&lt;organization&gt;</b>.webhook.office.com/webhookb2/<b>&lt;group&gt;</b>@<b>&lt;tenant&gt;</b>/IncomingWebhook/<b>&lt;altId&gt;</b>/<b>&lt;groupOwner&gt;</b></code></pre>
