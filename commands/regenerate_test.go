package commands_test

import (
	"net/http"

	"fmt"

	"github.com/cloudfoundry-incubator/credhub-cli/commands"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"
)

const REGENERATE_SECRET_REQUEST_JSON = `{"regenerate":true,"name":"my-password-stuffs"}`

var _ = Describe("Regenerate", func() {
	Describe("Regenerating password", func() {
		It("prints the regenerated password secret", func() {
			server.AppendHandlers(
				CombineHandlers(
					VerifyRequest("POST", "/api/v1/data"),
					VerifyJSON(REGENERATE_SECRET_REQUEST_JSON),
					RespondWith(http.StatusOK, fmt.Sprintf(STRING_SECRET_RESPONSE_JSON, "password", "my-password-stuffs", "nu-potatoes")),
				),
			)

			session := runCommand("regenerate", "--name", "my-password-stuffs")

			Eventually(session).Should(Exit(0))
			Expect(session.Out).To(Say(fmt.Sprintf(STRING_SECRET_RESPONSE_YAML, "password", "my-password-stuffs", "nu-potatoes")))
		})
	})

	Describe("help", func() {
		ItBehavesLikeHelp("regenerate", "r", func(session *Session) {
			Expect(session.Err).To(Say("regenerate"))
			Expect(session.Err).To(Say("name"))
		})

		It("has short flags", func() {
			Expect(commands.RegenerateCommand{}).To(SatisfyAll(
				commands.HaveFlag("name", "n"),
			))
		})
	})
})
