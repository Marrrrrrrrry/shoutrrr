# Matrix

:::info `title` 参数的使用
请注意，Matrix 会丢弃 `title` 参数中的所有信息，因为该服务没有与标题对应的概念。如需标题，请使用自定义消息格式，把想要的标题作为消息的一部分提供。
:::

## URL 格式

*matrix://__`user`__:__`password`__@__`host`__:__`port`__/[?rooms=__`!roomID1`__[,__`roomAlias2`__]][&disableTLS=yes]*

### URL 字段

*  __User__ - Username or empty when using access token  
  默认值：*empty*  
  URL 位置：<code class="service-url">matrix://<strong>user</strong>:password@host/</code>  
*  __Password__ - Password or access token （**必填**）  
  URL 位置：<code class="service-url">matrix://user:<strong>password</strong>@host/</code>  
*  __Host__ （**必填**）  
  URL 位置：<code class="service-url">matrix://user:password@<strong>host</strong>/</code>  
### 查询参数

这些参数既可以通过 params 参数传入，也可以直接通过 URL 传入：
`?key=value&key=value` etc.

*  __DeviceID__ - Device ID for password login; keeps Matrix homeservers from creating a new device for each login  
  默认值：`shoutrrr`  

*  __DisableTLS__  
  默认值：❌ `No`  

*  __Rooms__ - Room aliases, or with ! prefix, room IDs  
  默认值：*empty*  
  别名：`room`  

*  __Title__  
  默认值：*empty*

## 身份认证

如果没有指定 `user`，`password` 会被当作认证令牌（access token）。这意味着无论你的服务器使用哪种登录流程，只要你能手动获取一个令牌，Shoutrrr 就能使用它。

### 密码登录流程

如果同时提供了 `user` 和 `password`，在服务器支持的情况下会尝试 `m.login.password` 登录流程。Shoutrrr 在密码登录期间会发送一个固定的 Matrix `device_id`，避免每次初始化都创建一个新的 Matrix 设备。默认设备 ID 为 `shoutrrr`；如果同一账号的多个 Shoutrrr 实例需要各自独立的设备，可以用 `deviceID=...` 覆盖。

## 房间

如果*未*指定 `rooms`，服务会向该用户当前已加入的所有房间发送消息。

否则，服务只会向指定房间发送消息。如果用户*不在*其中任何房间，但收到了邀请，则会自动接受邀请。

**注意**：除非在 `rooms` 中显式指定，否则该服务**不会**加入任何房间。如果你需要用户加入这些房间，可以先在 `rooms` 中显式设置并发送一次通知。
