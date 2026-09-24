package teams

import (
	"fmt"
	"github.com/marrrrrrrrry/shoutrrr/pkg/format"
	"github.com/marrrrrrrrry/shoutrrr/pkg/types"
	"net/url"
	"regexp"
	"strings"

	"github.com/marrrrrrrrry/shoutrrr/pkg/services/standard"
)

// Config for use within the teams plugin
type Config struct {
	standard.EnumlessConfig
	Group      string `url:"user" optional:""`
	Tenant     string `url:"host" optional:""`
	AltID      string `url:"path1" optional:""`
	GroupOwner string `url:"path2" optional:""`

	Title string `key:"title" optional:""`
	Color string `key:"color" optional:""`
	Host  string `key:"host" optional:"" default:"outlook.office.com"`
	// WebhookURL is a full webhook URL, used for Workflows-based webhooks that
	// cannot be decomposed into the legacy connector URL parts
	WebhookURL string `key:"webhook" optional:"" desc:"Full webhook URL, e.g. a Teams Workflows request URL"`
}

func (config *Config) webhookParts() [4]string {
	return [4]string{config.Group, config.Tenant, config.AltID, config.GroupOwner}
}

// SetFromWebhookURL updates the config WebhookParts from a teams webhook URL
func (config *Config) SetFromWebhookURL(webhookURL string) error {
	parts, err := parseAndVerifyWebhookURL(webhookURL)
	if err != nil {
		return err
	}

	config.setFromWebhookParts(parts)
	return nil
}

// ConfigFromWebhookURL creates a new Config from a parsed Teams Webhook URL
func ConfigFromWebhookURL(webhookURL url.URL) (*Config, error) {
	config := &Config{
		Host: webhookURL.Host,
	}

	if err := config.SetFromWebhookURL(webhookURL.String()); err != nil {
		if !isWorkflowsWebhookURL(&webhookURL) {
			return nil, err
		}
		// Workflows request URLs cannot be decomposed into connector parts,
		// so the full URL is kept as-is (query parameters included)
		upstream := webhookURL
		upstream.Scheme = strings.TrimPrefix(webhookURL.Scheme, Scheme+"+")
		config = &Config{
			// keep the default host so that the query omits it
			Host:       LegacyHost,
			WebhookURL: upstream.String(),
		}
	}

	return config, nil
}

// GetURL returns a URL representation of it's current field values
func (config *Config) GetURL() *url.URL {
	resolver := format.NewPropKeyResolver(config)
	return config.getURL(&resolver)
}

// SetURL updates a ServiceConfig from a URL representation of it's field values
func (config *Config) SetURL(url *url.URL) error {
	resolver := format.NewPropKeyResolver(config)
	return config.setURL(&resolver, url)
}

func (config *Config) getURL(resolver types.ConfigQueryResolver) *url.URL {
	serviceURL := &url.URL{
		Scheme:     Scheme,
		ForceQuery: false,
		RawQuery:   format.BuildQuery(resolver),
	}

	// A full webhook URL cannot be expressed as connector URL parts;
	// Path "/" keeps the canonical "teams:///?webhook=..." form
	if config.WebhookURL != "" {
		serviceURL.Path = "/"
		return serviceURL
	}

	serviceURL.User = url.User(config.Group)
	serviceURL.Host = config.Tenant
	serviceURL.Path = "/" + config.AltID + "/" + config.GroupOwner
	return serviceURL
}

func (config *Config) setURL(resolver types.ConfigQueryResolver, url *url.URL) error {
	// A full webhook URL cannot coexist with the legacy connector URL parts,
	// so when present it takes precedence and part parsing is skipped
	for key := range url.Query() {
		if strings.EqualFold(key, "webhook") {
			for key, vals := range url.Query() {
				if err := resolver.Set(key, vals[0]); err != nil {
					return err
				}
			}
			if config.WebhookURL == "" {
				return fmt.Errorf("webhook URL cannot be empty")
			}
			return nil
		}
	}

	var webhookParts [4]string

	if pass, legacyFormat := url.User.Password(); legacyFormat {
		parts := strings.Split(url.User.Username(), "@")
		if len(parts) != 2 {
			return fmt.Errorf("invalid URL format")
		}
		webhookParts = [4]string{parts[0], parts[1], pass, url.Hostname()}
	} else {
		parts := strings.Split(url.Path, "/")
		if parts[0] == "" {
			parts = parts[1:]
		}
		webhookParts = [4]string{url.User.Username(), url.Hostname(), parts[0], parts[1]}
	}

	if err := verifyWebhookParts(webhookParts); err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	config.setFromWebhookParts(webhookParts)

	for key, vals := range url.Query() {
		if err := resolver.Set(key, vals[0]); err != nil {
			return err
		}
	}

	return nil
}

func (config *Config) setFromWebhookParts(parts [4]string) {
	config.Group = parts[0]
	config.Tenant = parts[1]
	config.AltID = parts[2]
	config.GroupOwner = parts[3]
}

// isWorkflowsWebhookURL checks whether the URL is a Power Automate Workflows
// request URL, which replaced the retired Office 365 Connector webhooks
func isWorkflowsWebhookURL(webhookURL *url.URL) bool {
	host := webhookURL.Hostname()
	return host == workflowsHost || strings.HasSuffix(host, "."+workflowsHost)
}

// hasConfigKey checks whether the resolver knows the given config key
func hasConfigKey(resolver types.ConfigQueryResolver, key string) bool {
	for _, known := range resolver.QueryFields() {
		if strings.EqualFold(known, key) {
			return true
		}
	}
	return false
}

// buildWebhookURL assembles the legacy connector webhook URL from its parts
func buildWebhookURL(host, group, tenant, altID, groupOwner string) string {
	// config.Group, config.Tenant, config.AltID, config.GroupOwner
	path := Path
	if host == LegacyHost {
		path = LegacyPath
	}
	return fmt.Sprintf(
		"https://%s/%s/%s@%s/%s/%s/%s",
		host,
		path,
		group,
		tenant,
		ProviderName,
		altID,
		groupOwner)
}

func parseAndVerifyWebhookURL(webhookURL string) (parts [4]string, err error) {
	pattern, err := regexp.Compile(`([0-9a-f-]{36})@([0-9a-f-]{36})/[^/]+/([0-9a-f]{32})/([0-9a-f-]{36})`)
	if err != nil {
		return parts, err
	}

	groups := pattern.FindStringSubmatch(webhookURL)
	if len(groups) != 5 {
		return parts, fmt.Errorf("invalid webhook URL format")
	}

	copy(parts[:], groups[1:])
	return parts, nil
}

const (
	// Scheme is the identifying part of this service's configuration URL
	Scheme = "teams"
	// LegacyHost is the default host for legacy webhook requests
	LegacyHost = "outlook.office.com"
	// LegacyPath is the initial path of the webhook URL for legacy webhook requests
	LegacyPath = "webhook"
	// Path is the initial path of the webhook URL for domain-scoped webhook requests
	Path = "webhookb2"
	// ProviderName is the name of the Teams integration provider
	ProviderName = "IncomingWebhook"
	// workflowsHost is the host of Power Automate Workflows request URLs,
	// which replaced the retired Office 365 Connector webhooks
	workflowsHost = "logic.azure.com"
)
