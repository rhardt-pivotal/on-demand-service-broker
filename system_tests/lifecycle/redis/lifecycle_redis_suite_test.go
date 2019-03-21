package lifecycle_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	. "github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/bosh_helpers"
	"github.com/pivotal-cf/on-demand-service-broker/system_tests/test_helpers/env_helpers"
)

func TestRedisLifecycle(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Redis Lifecycle Suite")
}

var (
	brokerInfo        BrokerInfo
	dopplerAddress    string
	deploymentOptions BrokerDeploymentOptions
	metricsOpsFile string
)

var _ = BeforeSuite(func() {
	err := env_helpers.ValidateEnvVars("DOPPLER_ADDRESS")
	Expect(err).ToNot(HaveOccurred(), "Doppler address must be set")

	dopplerAddress = os.Getenv("DOPPLER_ADDRESS")
	legacyMetrics := os.Getenv("LEGACY_SERVICE_METRICS")
	metricsOpsFile = "service_metrics.yml"
	if legacyMetrics == "true" {
		metricsOpsFile = "service_metrics_with_metron_agent.yml"
	}

	deploymentOptions = bosh_helpers.BrokerDeploymentOptions{
		ServiceMetrics: true,
		BrokerTLS:      true,
	}
})

var _ = AfterSuite(func() {
	bosh_helpers.DeregisterAndDeleteBroker(brokerInfo.DeploymentName)
})
