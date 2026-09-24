# OAuth2 生成器

OAuth2 生成器通过引导你完成 [XOAUTH2](https://developers.google.com/workspace/gmail/imap/xoauth2-protocol)
授权流程，生成带 OAuth2 认证的 Email（`smtp`）服务 URL：它会打印授权链接，用你输入的验证码换取访问令牌，
并把令牌嵌入生成的服务 URL。

## 用法

```shell
$ shoutrrr generate -s smtp -g oauth2 [OPTIONS] [CREDENTIALS.json]
```

根据选项不同有三种模式：

### Gmail（Google Cloud Console 的凭据文件）

使用 `provider=gmail` 属性并传入 Google OAuth2 客户端凭据 JSON 文件。
scope（`https://mail.google.com/`）和 SMTP 主机（`smtp.gmail.com`）会自动填入：

```shell
$ shoutrrr generate -s smtp -g oauth2 -p provider=gmail credentials.json
```

### 通用提供商（凭据文件）

不带 `provider` 属性时，生成器读取包含以下字段的 JSON 文件：
`client_id`、`client_secret`、`redirect_url`、`auth_url`、`token_url`、`smtp_hostname` 和 `scopes`：

```shell
$ shoutrrr generate -s smtp -g oauth2 oauth2.json
```

### 交互模式

不带文件参数时，生成器逐项询问所有参数
（ClientID、ClientSecret、AuthURL、TokenURL、RedirectURL、Scopes 和 SMTP Hostname）：

```shell
$ shoutrrr generate -s smtp -g oauth2
```

## 注意

* 访问令牌会作为 URL 密码部分嵌入。OAuth2 访问令牌是短期的（通常约 1 小时），
  因此生成的 URL 主要用于验证授权流程，或配合令牌刷新机制使用。
* Gmail 的 `https://mail.google.com/` 属于受限 scope：个人使用时，OAuth 应用处于"测试"模式
  会让令牌 7 天过期。
