# Teams

:::caution Office 365 Connector webhooks have been retired
Microsoft has retired the Office 365 Connectors that the legacy `teams://group@tenant/altId/groupOwner`
URL format is based on: new connectors were blocked in August 2024, existing ones stopped delivering
in April 2026 and were fully disabled in May 2026.

Use a **Power Automate Workflows** request URL instead (see [Setting up a webhook](#setting-up-a-webhook)).
The legacy URL format is kept for error reporting only and no longer delivers messages.
:::

## URL Format

:::info Workflows webhook (current)
<pre><code class="service-url">teams://?webhook=<strong>&lt;workflows-request-url&gt;</strong></code></pre>

or paste the workflow URL directly with a `teams+` scheme prefix:

<pre><code class="service-url">teams+https://<strong>&lt;region&gt;</strong>.logic.azure.com:443/workflows/<strong>&lt;workflow-id&gt;</strong>/triggers/manual/paths/invoke?api-version=2016-06-01&amp;sp=...&amp;sv=...&amp;sig=...</code></pre>
:::

### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Webhook__  
  The full Power Automate Workflows request URL. Cannot be combined with the legacy URL fields.  
  Default: *empty*

*  __Color__  
  Default: *empty*  

*  __Host__  
  Default: `outlook.office.com`  
  Only used by the legacy connector URL format.

*  __Title__  
  Default: *empty*

## Setting up a webhook

1. In Teams, open the channel you want to post to, select **More options (...)** → **Workflows**.
2. Pick the template **Post to a channel when a webhook request is received** and create the workflow.
3. Copy the request URL from the workflow details page. It looks like:
   `https://<region>.logic.azure.com:443/workflows/<workflow-id>/triggers/manual/paths/invoke?api-version=2016-06-01&sp=...&sv=...&sig=...`
4. Use it either as `teams://?webhook=<url>` (URL-escaped) or with a `teams+` scheme prefix:
   `teams+https://<region>.logic.azure.com/...`

:::info Payload format
The service sends a legacy MessageCard payload. Workflows webhooks render MessageCards,
but interactive elements (HttpPost action buttons) are not supported. Messages are limited
to 28 KB and the sender is shown as the Workflows (Flow) bot.
:::

## Legacy URL format (retired)

The connector-based format decomposed the webhook URL into four parts:

<pre><code class="service-url">teams://<strong>group</strong>@<strong>tenant</strong>/<strong>altId</strong>/<strong>groupOwner</strong>?host=<strong>&lt;organization&gt;</strong>.webhook.office.com</code></pre>

*  __Group__  
  URL part: <code class="service-url">teams://<strong>group</strong>@tenant/altid/groupowner</code>  
*  __Tenant__  
  URL part: <code class="service-url">teams://group@<strong>tenant</strong>/altid/groupowner</code>  
*  __AltID__  
  URL part: <code class="service-url">teams://group@tenant/<strong>altid</strong>/groupowner</code>  
*  __GroupOwner__  
  URL part: <code class="service-url">teams://group@tenant/altid/<strong>groupowner</strong></code>  

The token used to be extracted from a connector webhook URL:

<pre><code>https://<b>&lt;organization&gt;</b>.webhook.office.com/webhookb2/<b>&lt;group&gt;</b>@<b>&lt;tenant&gt;</b>/IncomingWebhook/<b>&lt;altId&gt;</b>/<b>&lt;groupOwner&gt;</b></code></pre>

Since connector webhooks no longer deliver messages, this format only produces errors.
