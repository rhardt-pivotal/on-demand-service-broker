package on_demand_service_broker_test

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/brokerapi"
	brokerConfig "github.com/pivotal-cf/on-demand-service-broker/config"
)

const (
	serverCertFile = "./assets/mybroker.crt"
	serverKeyFile  = "./assets/mybroker.key"
	caCertFile     = "./assets/bosh.ca.crt"
)

var acceptableCipherSuites = []uint16{
	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
}

var _ = Describe("Server Protocol", func() {
	Describe("with HTTPS configured", func() {
		BeforeEach(func() {
			conf := brokerConfig.Config{
				Broker: brokerConfig.Broker{
					Port: serverPort, Username: brokerUsername, Password: brokerPassword,
					TLS: brokerConfig.TLSConfig{
						CertFile: serverCertFile,
						KeyFile:  serverKeyFile,
					},
				},
				ServiceCatalog: brokerConfig.ServiceOffering{
					Name: serviceName,
				},
			}

			StartServer(conf)
		})

		It("serves HTTPS", func() {
			response, bodyContent := doHTTPSRequest(http.MethodGet, fmt.Sprintf("https://%s/v2/catalog", serverURL), nil, caCertFile, acceptableCipherSuites)

			Expect(response.StatusCode).To(Equal(http.StatusOK))
			catalog := make(map[string][]brokerapi.Service)
			Expect(json.Unmarshal(bodyContent, &catalog)).To(Succeed())
			Expect(catalog["services"][0].Name).To(Equal(serviceName))

			// Expect(loggerBuffer).Should(gbytes.Say("http: TLS handshake error"))
		})

		DescribeTable("can use the desired cipher suites",
			func(cipher uint16) {
				response, _ := doHTTPSRequest(http.MethodGet, fmt.Sprintf("https://%s/v2/catalog", serverURL), nil, caCertFile, []uint16{cipher})

				Expect(response.StatusCode).To(Equal(http.StatusOK))
				Expect(response.TLS.CipherSuite).To(Equal(cipher))
			},
			// The first two cipher suites that Pivotal recommends are not available in Golang
			// Entry("TLS_DHE_RSA_WITH_AES_128_GCM_SHA256", tls.TLS_DHE_RSA_WITH_AES_128_GCM_SHA256),
			// Entry("TLS_DHE_RSA_WITH_AES_256_GCM_SHA384", tls.TLS_DHE_RSA_WITH_AES_256_GCM_SHA384),
			Entry("TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256),
			Entry("TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384),
		)

		DescribeTable("does not serve when the client uses an unacceptable cipher",
			func(cipher uint16) {
				url := fmt.Sprintf("https://%s/v2/catalog", serverURL)
				tlsConfig := &tls.Config{
					CipherSuites:       []uint16{cipher},
					InsecureSkipVerify: true,
				}
				transport := &http.Transport{TLSClientConfig: tlsConfig}
				client := &http.Client{Transport: transport}

				req, err := http.NewRequest(http.MethodGet, url, nil)
				Expect(err).ToNot(HaveOccurred())

				req.SetBasicAuth(brokerUsername, brokerPassword)
				req.Header.Set("X-Broker-API-Version", "2.14")

				_, err = client.Do(req)
				Expect(err).To(MatchError(ContainSubstring("remote error: tls: handshake failure")))
			},
			Entry("TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305", tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305),
			Entry("TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305", tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA),
			Entry("TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305", tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256),
		)

		It("refuses to serve TLS 1.1", func() {
			url := fmt.Sprintf("https://%s/v2/catalog", serverURL)
			tlsConfig := &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         tls.VersionTLS11,
				MaxVersion:         tls.VersionTLS11,
			}
			transport := &http.Transport{TLSClientConfig: tlsConfig}
			client := &http.Client{Transport: transport}

			req, err := http.NewRequest(http.MethodGet, url, nil)
			Expect(err).ToNot(HaveOccurred())

			req.SetBasicAuth(brokerUsername, brokerPassword)
			req.Header.Set("X-Broker-API-Version", "2.14")
			_, err = client.Do(req)
			Expect(err).To(MatchError(ContainSubstring("remote error: tls: protocol version not supported")))
		})

		It("does not serve HTTP", func() {
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/v2/catalog", serverURL), nil)
			Expect(err).ToNot(HaveOccurred())

			req.SetBasicAuth(brokerUsername, brokerPassword)
			req.Header.Set("X-Broker-API-Version", "2.14")
			req.Close = true

			_, err = http.DefaultClient.Do(req)
			Expect(err).To(MatchError(ContainSubstring("HTTP/1.x transport connection broken: malformed HTTP response")))
		})
	})

	Describe("with HTTP configured", func() {
		BeforeEach(func() {
			conf := brokerConfig.Config{
				Broker: brokerConfig.Broker{
					Port: serverPort, Username: brokerUsername, Password: brokerPassword,
				},
				ServiceCatalog: brokerConfig.ServiceOffering{
					Name: serviceName,
				},
			}

			StartServer(conf)
		})
		It("serves HTTP", func() {
			response, bodyContent := doRequest(http.MethodGet, fmt.Sprintf("http://%s/v2/catalog", serverURL), nil)

			Expect(response.StatusCode).To(Equal(http.StatusOK))
			catalog := make(map[string][]brokerapi.Service)
			Expect(json.Unmarshal(bodyContent, &catalog)).To(Succeed())
		})
	})
})
