package teams

import (
	"errors"
	"log"
	"net/url"
	"testing"

	"github.com/jarcoal/httpmock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	legacyWebhookURL = "https://outlook.office.com/webhook/11111111-4444-4444-8444-cccccccccccc@22222222-4444-4444-8444-cccccccccccc/IncomingWebhook/33333333012222222222333333333344/44444444-4444-4444-8444-cccccccccccc"
	scopedWebhookURL = "https://test.webhook.office.com/webhookb2/11111111-4444-4444-8444-cccccccccccc@22222222-4444-4444-8444-cccccccccccc/IncomingWebhook/33333333012222222222333333333344/44444444-4444-4444-8444-cccccccccccc"
	scopedDomainHost = "test.webhook.office.com"
	testURLBase      = "teams://11111111-4444-4444-8444-cccccccccccc@22222222-4444-4444-8444-cccccccccccc/33333333012222222222333333333344/44444444-4444-4444-8444-cccccccccccc"
	scopedURLBase    = testURLBase + `?host=` + scopedDomainHost
	// Workflows request URLs replaced the retired Office 365 Connector webhooks
	workflowsWebhookURL = "https://prod-25.northeurope.logic.azure.com:443/workflows/111111114444444484444cccccccccccc/triggers/manual/paths/invoke?api-version=2016-06-01&sp=%2Ftriggers%2Fmanual%2Frun&sv=1.0&sig=AbCdEf123456"
)

var logger = log.New(GinkgoWriter, "Test", log.LstdFlags)

func TestTeams(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Shoutrrr Teams Suite")
}

var _ = Describe("the teams service", func() {
	When("creating the webhook URL", func() {
		It("should match the expected output for legacy URLs", func() {
			config := Config{}
			config.setFromWebhookParts([4]string{
				"11111111-4444-4444-8444-cccccccccccc",
				"22222222-4444-4444-8444-cccccccccccc",
				"33333333012222222222333333333344",
				"44444444-4444-4444-8444-cccccccccccc",
			})
			apiURL := buildWebhookURL(LegacyHost, config.Group, config.Tenant, config.AltID, config.GroupOwner)
			Expect(apiURL).To(Equal(legacyWebhookURL))

			parts, err := parseAndVerifyWebhookURL(apiURL)
			Expect(err).ToNot(HaveOccurred())
			Expect(parts).To(Equal(config.webhookParts()))
		})
		It("should match the expected output for custom URLs", func() {
			config := Config{}
			config.setFromWebhookParts([4]string{
				"11111111-4444-4444-8444-cccccccccccc",
				"22222222-4444-4444-8444-cccccccccccc",
				"33333333012222222222333333333344",
				"44444444-4444-4444-8444-cccccccccccc",
			})
			apiURL := buildWebhookURL(scopedDomainHost, config.Group, config.Tenant, config.AltID, config.GroupOwner)
			Expect(apiURL).To(Equal(scopedWebhookURL))

			parts, err := parseAndVerifyWebhookURL(apiURL)
			Expect(err).ToNot(HaveOccurred())
			Expect(parts).To(Equal(config.webhookParts()))
		})
	})

	Describe("creating a config", func() {
		When("parsing the configuration URL", func() {
			It("should be identical after de-/serialization", func() {
				testURL := testURLBase + "?color=aabbcc&host=test.outlook.office.com&title=Test+title"

				url, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				config := &Config{Host: LegacyHost}
				err = config.SetURL(url)
				Expect(err).NotTo(HaveOccurred(), "verifying")

				outputURL := config.GetURL()
				Expect(outputURL.String()).To(Equal(testURL))
			})
		})
	})

	Describe("converting custom URL to service URL", func() {
		When("an invalid custom URL is provided", func() {
			It("should return an error", func() {
				service := Service{}
				testURL := "teams+https://google.com/search?q=what+is+love"

				customURL, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				_, err = service.GetConfigURLFromCustom(customURL)
				Expect(err).To(HaveOccurred(), "converting")
			})
		})
		When("a valid custom URL is provided", func() {
			It("should set the host field from the custom URL", func() {
				service := Service{}
				testURL := `teams+` + scopedWebhookURL

				customURL, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				serviceURL, err := service.GetConfigURLFromCustom(customURL)
				Expect(err).NotTo(HaveOccurred(), "converting")

				Expect(serviceURL.String()).To(Equal(scopedURLBase))
			})
			It("should preserve the query params in the generated service URL", func() {
				service := Service{}
				testURL := "teams+" + legacyWebhookURL + "?color=f008c1&title=TheTitle"

				customURL, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				serviceURL, err := service.GetConfigURLFromCustom(customURL)
				Expect(err).NotTo(HaveOccurred(), "converting")

				Expect(serviceURL.String()).To(Equal(testURLBase + "?color=f008c1&title=TheTitle"))
			})
		})
	})

	Describe("converting Workflows custom URL to service URL", func() {
		When("a Workflows request URL is provided", func() {
			It("should keep the full URL in the webhook query param", func() {
				service := Service{}
				testURL := `teams+` + workflowsWebhookURL

				customURL, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				serviceURL, err := service.GetConfigURLFromCustom(customURL)
				Expect(err).NotTo(HaveOccurred(), "converting")

				expectedQuery := url.Values{"webhook": []string{workflowsWebhookURL}}
				Expect(serviceURL.Scheme).To(Equal(Scheme))
				Expect(serviceURL.Query()).To(Equal(expectedQuery))
			})
		})
		When("a service URL with a webhook query param is parsed", func() {
			It("should be identical after de-/serialization", func() {
				testURL := "teams:///?webhook=" + url.QueryEscape(workflowsWebhookURL)

				parsedURL, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				config := &Config{Host: LegacyHost}
				err = config.SetURL(parsedURL)
				Expect(err).NotTo(HaveOccurred(), "verifying")
				Expect(config.WebhookURL).To(Equal(workflowsWebhookURL))

				outputURL := config.GetURL()
				Expect(outputURL.String()).To(Equal(testURL))
			})
			It("should reject an empty webhook value", func() {
				parsedURL, err := url.Parse("teams://?webhook=")
				Expect(err).NotTo(HaveOccurred(), "parsing")

				config := &Config{}
				err = config.SetURL(parsedURL)
				Expect(err).To(HaveOccurred(), "empty webhook must not be accepted")
			})
			It("should accept the webhook key in any case", func() {
				parsedURL, err := url.Parse("teams://?Webhook=" + url.QueryEscape(workflowsWebhookURL))
				Expect(err).NotTo(HaveOccurred(), "parsing")

				config := &Config{}
				err = config.SetURL(parsedURL)
				Expect(err).NotTo(HaveOccurred(), "verifying")
				Expect(config.WebhookURL).To(Equal(workflowsWebhookURL))
			})
			It("should prefer the webhook URL over connector URL parts", func() {
				parsedURL, err := url.Parse(testURLBase + "?webhook=" + url.QueryEscape(workflowsWebhookURL))
				Expect(err).NotTo(HaveOccurred(), "parsing")

				config := &Config{Host: LegacyHost}
				err = config.SetURL(parsedURL)
				Expect(err).NotTo(HaveOccurred(), "verifying")
				Expect(config.WebhookURL).To(Equal(workflowsWebhookURL))
				Expect(config.Group).To(Equal(""))
			})
		})
		When("a legacy connector custom URL has an unknown query param", func() {
			It("should still return an error", func() {
				service := Service{}
				testURL := "teams+" + scopedWebhookURL + "?nosuchkey=value"

				customURL, err := url.Parse(testURL)
				Expect(err).NotTo(HaveOccurred(), "parsing")

				_, err = service.GetConfigURLFromCustom(customURL)
				Expect(err).To(HaveOccurred(), "unknown keys must not be silently dropped")
			})
		})
	})

	Describe("sending the payload", func() {
		var err error
		var service Service
		BeforeEach(func() {
			httpmock.Activate()
		})
		AfterEach(func() {
			httpmock.DeactivateAndReset()
		})
		It("should not report an error if the server accepts the payload", func() {
			serviceURL, _ := url.Parse(scopedURLBase)
			err = service.Initialize(serviceURL, logger)
			Expect(err).NotTo(HaveOccurred())

			httpmock.RegisterResponder("POST", scopedWebhookURL, httpmock.NewStringResponder(200, ""))

			err = service.Send("Message", nil)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should not panic if an error occurs when sending the payload", func() {
			serviceURL, _ := url.Parse(testURLBase)
			err = service.Initialize(serviceURL, logger)
			Expect(err).NotTo(HaveOccurred())

			httpmock.RegisterResponder("POST", legacyWebhookURL, httpmock.NewErrorResponder(errors.New("dummy error")))

			err = service.Send("Message", nil)
			Expect(err).To(HaveOccurred())
		})
		It("should accept a Workflows 202 Accepted response", func() {
			serviceURL, _ := url.Parse("teams://?webhook=" + url.QueryEscape(workflowsWebhookURL))
			err = service.Initialize(serviceURL, logger)
			Expect(err).NotTo(HaveOccurred())

			httpmock.RegisterResponder("POST", workflowsWebhookURL, httpmock.NewStringResponder(202, ""))

			err = service.Send("Message", nil)
			Expect(err).NotTo(HaveOccurred())
		})
		It("should report an error when the Workflows endpoint rejects the payload", func() {
			serviceURL, _ := url.Parse("teams://?webhook=" + url.QueryEscape(workflowsWebhookURL))
			err = service.Initialize(serviceURL, logger)
			Expect(err).NotTo(HaveOccurred())

			httpmock.RegisterResponder("POST", workflowsWebhookURL, httpmock.NewStringResponder(400, ""))

			err = service.Send("Message", nil)
			Expect(err).To(HaveOccurred())
		})

	})

})
