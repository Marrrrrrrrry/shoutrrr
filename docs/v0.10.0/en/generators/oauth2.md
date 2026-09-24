# OAuth2 generator

The OAuth2 generator creates an Email (`smtp`) service URL with OAuth2 authentication by walking you
through the [XOAUTH2](https://developers.google.com/workspace/gmail/imap/xoauth2-protocol) authorization flow:
it prints an authorization URL, exchanges your verification code for an access token and embeds the
token in the resulting service URL.

## Usage

```shell
$ shoutrrr generate -s smtp -g oauth2 [OPTIONS] [CREDENTIALS.json]
```

Three modes are available, depending on the options:

### Gmail (credentials file from Google Cloud Console)

Use the `provider=gmail` property with a Google OAuth2 client credentials JSON file.
The scope (`https://mail.google.com/`) and the SMTP host (`smtp.gmail.com`) are filled in automatically:

```shell
$ shoutrrr generate -s smtp -g oauth2 -p provider=gmail credentials.json
```

### Generic provider (credentials file)

Without the `provider` property, the generator reads a JSON file with the following fields:
`client_id`, `client_secret`, `redirect_url`, `auth_url`, `token_url`, `smtp_hostname` and `scopes`:

```shell
$ shoutrrr generate -s smtp -g oauth2 oauth2.json
```

### Interactive

Without any file argument, the generator prompts for all parameters
(ClientID, ClientSecret, AuthURL, TokenURL, RedirectURL, Scopes and SMTP Hostname):

```shell
$ shoutrrr generate -s smtp -g oauth2
```

## Notes

* The access token is embedded as the URL password. OAuth2 access tokens are short-lived
  (typically about one hour), so the generated URL is mainly useful for verifying the flow or
  together with a token refresh mechanism.
* The Gmail scope `https://mail.google.com/` is a restricted scope: for personal use,
  publishing the OAuth app in "testing" mode limits tokens to 7 days.
