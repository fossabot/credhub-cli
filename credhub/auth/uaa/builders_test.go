package uaa

import (
	"errors"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/credhub-cli/credhub/auth"
)

type DummyServerConfig struct {
	Error error
}

func (d *DummyServerConfig) AuthURL() (string, error) {
	return "http://example.com/auth/url", d.Error
}

func (d *DummyServerConfig) Client() *http.Client {
	return http.DefaultClient
}

var _ = Describe("Constructors", func() {
	Describe("PasswordGrant()", func() {
		It("constructs a OAuthStrategy auth using password grant", func() {
			config := DummyServerConfig{}
			builder := PasswordGrantBuilder("some-client-id", "some-client-secret", "some-username", "some-password")
			strategy, _ := builder(&config)
			auth := strategy.(*auth.OAuthStrategy)
			Expect(auth.ClientId).To(Equal("some-client-id"))
			Expect(auth.ClientSecret).To(Equal("some-client-secret"))
			Expect(auth.Username).To(Equal("some-username"))
			Expect(auth.Password).To(Equal("some-password"))
			Expect(auth.OAuthClient.(*Client).AuthURL).To(Equal("http://example.com/auth/url"))
			client := config.Client()
			Expect(auth.OAuthClient.(*Client).Client).To(BeIdenticalTo(client))
			Expect(auth.ApiClient).To(BeIdenticalTo(client))
		})

		Context("when fetching an Auth URL fails", func() {
			It("returns an error", func() {
				config := DummyServerConfig{
					Error: errors.New("Failed to fetch Auth URL"),
				}
				builder := PasswordGrantBuilder("some-client-id", "some-client-secret", "some-username", "some-password")
				_, err := builder(&config)

				Expect(err).To(MatchError("Failed to fetch Auth URL"))
			})

		})
	})

	Describe("ClientCredentialsGrant()", func() {
		It("constructs a OAuthStrategy auth using client credentials grant", func() {
			config := DummyServerConfig{}
			builder := ClientCredentialsGrantBuilder("some-client-id", "some-client-secret")
			strategy, _ := builder(&config)
			auth := strategy.(*auth.OAuthStrategy)
			Expect(auth.ClientId).To(Equal("some-client-id"))
			Expect(auth.ClientSecret).To(Equal("some-client-secret"))
			Expect(auth.Username).To(BeEmpty())
			Expect(auth.Password).To(BeEmpty())
			Expect(auth.OAuthClient.(*Client).AuthURL).To(Equal("http://example.com/auth/url"))
			client := config.Client()
			Expect(auth.OAuthClient.(*Client).Client).To(BeIdenticalTo(client))
			Expect(auth.ApiClient).To(BeIdenticalTo(client))
		})

		Context("when fetching an Auth URL fails", func() {
			It("returns an error", func() {
				config := DummyServerConfig{
					Error: errors.New("Failed to fetch Auth URL"),
				}
				builder := ClientCredentialsGrantBuilder("some-client-id", "some-client-secret")
				_, err := builder(&config)

				Expect(err).To(MatchError("Failed to fetch Auth URL"))
			})

		})
	})

	Describe("AuthBuilder()", func() {
		It("constructs a OAuthStrategy auth using existing tokens", func() {
			config := DummyServerConfig{}
			builder := AuthBuilder("some-client-id",
				"some-client-secret",
				"some-username",
				"some-password",
				"some-access-token",
				"some-refresh-token")
			strategy, _ := builder(&config)
			auth := strategy.(*auth.OAuthStrategy)
			Expect(auth.ClientId).To(Equal("some-client-id"))
			Expect(auth.ClientSecret).To(Equal("some-client-secret"))
			Expect(auth.Username).To(Equal("some-username"))
			Expect(auth.Password).To(Equal("some-password"))
			Expect(auth.AccessToken()).To(Equal("some-access-token"))
			Expect(auth.RefreshToken()).To(Equal("some-refresh-token"))
			Expect(auth.OAuthClient.(*Client).AuthURL).To(Equal("http://example.com/auth/url"))
			client := config.Client()
			Expect(auth.OAuthClient.(*Client).Client).To(BeIdenticalTo(client))
			Expect(auth.ApiClient).To(BeIdenticalTo(client))
		})

		Context("when fetching an Auth URL fails", func() {
			It("returns an error", func() {
				config := DummyServerConfig{
					Error: errors.New("Failed to fetch Auth URL"),
				}
				builder := AuthBuilder("some-client-id",
					"some-client-secret",
					"some-username",
					"some-password",
					"some-access-token",
					"some-refresh-token")
				_, err := builder(&config)

				Expect(err).To(MatchError("Failed to fetch Auth URL"))
			})
		})
	})
})
