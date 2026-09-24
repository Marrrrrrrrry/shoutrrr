# Ntfy

上游文档：https://docs.ntfy.sh/publish/

## URL 格式

### URL 字段

*  __Username__ - Auth username  
  默认值：*empty*  
  URL 位置：<code class="service-url">ntfy://<strong>username</strong>:password@host/topic</code>  
*  __Password__ - Auth password  
  默认值：*empty*  
  URL 位置：<code class="service-url">ntfy://username:<strong>password</strong>@host/topic</code>  
*  __Host__ - Server hostname and port  
  默认值：`ntfy.sh`  
  URL 位置：<code class="service-url">ntfy://username:password@<strong>host</strong>/topic</code>  
*  __Topic__ - Target topic name （**必填**）  
  URL 位置：<code class="service-url">ntfy://username:password@host/<strong>topic</strong></code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __Actions__ - Custom user action buttons for notifications, see https://docs.ntfy.sh/publish/#action-buttons  
  默认值：*empty*  

*  __Attach__ - URL of an attachment, see attach via URL  
  默认值：*empty*  

*  __Cache__ - Cache messages  
  默认值：✔ `yes`  

*  __Click__ - Website opened when notification is clicked  
  默认值：*empty*  

*  __Delay__ - Timestamp or duration for delayed delivery, see https://docs.ntfy.sh/publish/#scheduled-delivery  
  默认值：*empty*  
  别名：`at`, `in`  

*  __Email__ - E-mail address for e-mail notifications  
  默认值：*empty*  

*  __Filename__ - File name of the attachment  
  默认值：*empty*  

*  __Firebase__ - Send to firebase  
  默认值：✔ `yes`  

*  __Icon__ - URL to use as notification icon  
  默认值：*empty*  

*  __Markdown__ - Enable markdown formatting  
  默认值：❌ `no`  

*  __Priority__ - Message priority with 1=min, 3=default and 5=max  
  默认值：`default`  
  可选值：`Min`, `Low`, `Default`, `High`, `Max`  

*  __Scheme__ - Server protocol, http or https  
  默认值：`https`  

*  __Tags__ - List of tags that may or not map to emojis  
  默认值：*empty*  

*  __Title__ - Message title  
  默认值：*empty*
