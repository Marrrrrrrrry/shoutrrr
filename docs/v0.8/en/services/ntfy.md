# Ntfy

Upstream docs: https://docs.ntfy.sh/publish/

## URL Format

### URL Fields

*  __Username__ - Auth username  
  Default: *empty*  
  URL part: <code class="service-url">ntfy://<strong>username</strong>:password@host/topic</code>  
*  __Password__ - Auth password  
  Default: *empty*  
  URL part: <code class="service-url">ntfy://username:<strong>password</strong>@host/topic</code>  
*  __Host__ - Server hostname and port  
  Default: `ntfy.sh`  
  URL part: <code class="service-url">ntfy://username:password@<strong>host</strong>/topic</code>  
*  __Topic__ - Target topic name (**Required**)  
  URL part: <code class="service-url">ntfy://username:password@host/<strong>topic</strong></code>  
### Query/Param Props

Props can be either supplied using the params argument, or through the URL using  
`?key=value&key=value` etc.

*  __Actions__ - Custom user action buttons for notifications, see https://docs.ntfy.sh/publish/#action-buttons  
  Default: *empty*  

*  __Attach__ - URL of an attachment, see attach via URL  
  Default: *empty*  

*  __Cache__ - Cache messages  
  Default: ✔ `yes`  

*  __Click__ - Website opened when notification is clicked  
  Default: *empty*  

*  __Delay__ - Timestamp or duration for delayed delivery, see https://docs.ntfy.sh/publish/#scheduled-delivery  
  Default: *empty*  
  Aliases: `at`, `in`  

*  __Email__ - E-mail address for e-mail notifications  
  Default: *empty*  

*  __Filename__ - File name of the attachment  
  Default: *empty*  

*  __Firebase__ - Send to firebase  
  Default: ✔ `yes`  

*  __Icon__ - URL to use as notification icon  
  Default: *empty*  

*  __Markdown__ - Enable markdown formatting  
  Default: ❌ `no`  

*  __Priority__ - Message priority with 1=min, 3=default and 5=max  
  Default: `default`  
  Possible values: `Min`, `Low`, `Default`, `High`, `Max`  

*  __Scheme__ - Server protocol, http or https  
  Default: `https`  

*  __Tags__ - List of tags that may or not map to emojis  
  Default: *empty*  

*  __Title__ - Message title  
  Default: *empty*