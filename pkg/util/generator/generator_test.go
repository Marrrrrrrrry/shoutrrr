package generator_test

import (
	"fmt"
	"io"
	re "regexp"
	"strings"
	"testing"

	"github.com/marrrrrrrrry/shoutrrr/pkg/util/generator"
	"github.com/mattn/go-colorable"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

func TestGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generator Suite")
}

var (
	client  *generator.UserDialog
	userOut *io.PipeWriter // typed input, consumed by the dialog; writes block until read
	typed   strings.Builder
	userIn  *gbytes.Buffer // dialog output (prompts and replies)
)

// mockTyped simulates keyboard input. Writing to the pipe blocks until the
// dialog's scanner consumes the data, which synchronizes the test with the
// dialog goroutine and avoids read-EOF races.
func mockTyped(a ...interface{}) {
	_, _ = fmt.Fprint(&typed, a...)
	_, _ = fmt.Fprint(&typed, "\n")
	_, _ = fmt.Fprint(userOut, a...)
	_, _ = fmt.Fprint(userOut, "\n")
}

func dumpBuffers() {
	for _, line := range strings.Split(string(userIn.Contents()), "\n") {
		println(">", line)
	}
	for _, line := range strings.Split(typed.String(), "\n") {
		println("<", line)
	}
}

var _ = Describe("GeneratorCommon", func() {
	BeforeEach(func() {
		typed.Reset()
		reader, writer := io.Pipe()
		userOut = writer
		userIn = gbytes.NewBuffer()
		userInMono := colorable.NewNonColorable(userIn)
		client = generator.NewUserDialog(reader, userInMono, map[string]string{"propKey": "propVal"})
	})

	It("reprompt upon invalid answers", func() {
		defer dumpBuffers()
		answer := make(chan string)
		go func() {
			answer <- client.QueryString("name:", generator.Required, "")
		}()

		Eventually(userIn, "5s").Should(gbytes.Say(`name: `))
		mockTyped("")
		Eventually(userIn, "5s").Should(gbytes.Say(`field is required`))
		Eventually(userIn, "5s").Should(gbytes.Say(`name: `))
		mockTyped("Normal Human Name")

		Eventually(answer, "5s").Should(Receive(Equal("Normal Human Name")))
	})

	It("should accept any input when validator is nil", func() {
		defer dumpBuffers()
		answer := make(chan string)
		go func() {
			answer <- client.QueryString("name:", nil, "")
		}()

		Eventually(userIn, "5s").Should(gbytes.Say(`name: `))
		mockTyped("")

		Eventually(answer, "5s").Should(Receive(BeEmpty()))
	})

	It("should use predefined prop value if key is present", func() {
		defer dumpBuffers()
		answer := make(chan string)
		go func() {
			answer <- client.QueryString("name:", generator.Required, "propKey")
		}()
		Eventually(answer, "5s").Should(Receive(Equal("propVal")))
	})

	Describe("Query", func() {
		It("should prompt until a valid answer is provided", func() {
			defer dumpBuffers()
			answer := make(chan []string)
			query := "pick foo or bar:"
			go func() {
				answer <- client.Query(query, re.MustCompile("(foo|bar)"), "")
			}()

			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("")
			Eventually(userIn, "5s").Should(gbytes.Say(`invalid format`))
			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("foo")

			Eventually(answer, "5s").Should(Receive(ContainElement("foo")))
		})
	})

	Describe("QueryAll", func() {
		It("should prompt until a valid answer is provided", func() {
			defer dumpBuffers()
			answer := make(chan [][]string)
			query := "pick foo or bar:"
			go func() {
				answer <- client.QueryAll(query, re.MustCompile(`foo(ba[rz])`), "", -1)
			}()

			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("foobar foobaz")

			var matches [][]string
			Eventually(answer, "5s").Should(Receive(&matches))
			Expect(matches).To(ContainElement([]string{"foobar", "bar"}))
			Expect(matches).To(ContainElement([]string{"foobaz", "baz"}))
		})
	})

	Describe("QueryStringPattern", func() {
		It("should prompt until a valid answer is provided", func() {
			defer dumpBuffers()
			answer := make(chan string)
			query := "type of bar:"
			go func() {
				answer <- client.QueryStringPattern(query, re.MustCompile(".*bar"), "")
			}()

			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("foo")
			Eventually(userIn, "5s").Should(gbytes.Say(`invalid format`))
			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("foobar")

			Eventually(answer, "5s").Should(Receive(Equal("foobar")))
		})
	})

	Describe("QueryInt", func() {
		It("should prompt until a valid answer is provided", func() {
			defer dumpBuffers()
			answer := make(chan int64)
			query := "number:"
			go func() {
				answer <- client.QueryInt(query, "", 64)
			}()

			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("x")
			Eventually(userIn, "5s").Should(gbytes.Say(`not a number`))
			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("0x20")

			Eventually(answer, "5s").Should(Receive(Equal(int64(32))))
		})
	})

	Describe("QueryBool", func() {
		It("should prompt until a valid answer is provided", func() {
			defer dumpBuffers()
			answer := make(chan bool)
			query := "cool?"
			go func() {
				answer <- client.QueryBool(query, "")
			}()

			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("maybe")
			Eventually(userIn, "5s").Should(gbytes.Say(`answer using yes or no`))
			Eventually(userIn, "5s").Should(gbytes.Say(query))
			mockTyped("y")

			Eventually(answer, "5s").Should(Receive(BeTrue()))
		})
	})
})
