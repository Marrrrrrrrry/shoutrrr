# 代理

要在 shoutrrr 中使用代理，既可以设置环境变量 `HTTP_PROXY` 中的代理 URL，也可以像下面这样覆盖默认的 HTTP 客户端：

```go
proxyurl, err := url.Parse("socks5://localhost:1337")
if err != nil {
	log.Fatalf("Error parsing proxy URL: %q", err)
}

http.DefaultClient.Transport = &http.Transport{
	Proxy: http.ProxyURL(proxyurl),
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
```
