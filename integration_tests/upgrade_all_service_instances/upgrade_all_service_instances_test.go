// Copyright (C) 2016-Present Pivotal Software, Inc. All rights reserved.
// This program and the accompanying materials are made available under the terms of the under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

package upgrade_all_service_instances_test

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/ghttp"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	. "github.com/pivotal-cf/on-demand-service-broker/integration_tests/helpers"
	"github.com/pivotal-cf/on-demand-service-broker/loggerfactory"
	"gopkg.in/yaml.v2"
)

var _ = Describe("running the tool to upgrade all service instances", func() {
	const (
		brokerUsername = "broker username"
		brokerPassword = "broker password"
	)

	var errandConfig config.InstanceIteratorConfig

	startUpgradeAllInstanceBinary := func(errandConfig config.InstanceIteratorConfig) *gexec.Session {
		b, err := yaml.Marshal(errandConfig)
		Expect(err).ToNot(HaveOccurred())
		configPath := writeConfigFile(string(b))
		return StartBinaryWithParams(binaryPath, []string{"-configPath", configPath})
	}

	Describe("upgrading via BOSH", func() {

		var broker *ghttp.Server

		Describe("HTTP Broker", func() {
			var (
				serviceInstances string
				operationData    string
				instanceID       string

				lastOperationHandler    *FakeHandler
				serviceInstancesHandler *FakeHandler
				upgradeHandler          *FakeHandler
			)

			BeforeEach(func() {
				broker = ghttp.NewServer()

				lastOperationHandler = new(FakeHandler)
				serviceInstancesHandler = new(FakeHandler)
				upgradeHandler = new(FakeHandler)

				errandConfig = config.InstanceIteratorConfig{
					PollingInterval: 1,
					AttemptLimit:    2,
					AttemptInterval: 2,
					MaxInFlight:     1,
					BrokerAPI: config.BrokerAPI{
						URL: broker.URL(),
						Authentication: config.Authentication{
							Basic: config.UserCredentials{
								Username: brokerUsername,
								Password: brokerPassword,
							},
						},
					},
				}

				broker.RouteToHandler(http.MethodGet, "/mgmt/service_instances", ghttp.CombineHandlers(
					ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
					serviceInstancesHandler.Handle,
				))

				broker.RouteToHandler(http.MethodPatch, regexp.MustCompile(`/mgmt/service_instances/.*`), ghttp.CombineHandlers(
					ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
					ghttp.VerifyRequest(http.MethodPatch, ContainSubstring("/mgmt/service_instances/"), "operation_type=upgrade"),
					upgradeHandler.Handle,
				))

				broker.RouteToHandler(http.MethodGet, regexp.MustCompile(`/v2/service_instances/.*/last_operation`), ghttp.CombineHandlers(
					ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
					lastOperationHandler.Handle,
				))

				operationData = `{"BoshTaskID":1,"OperationType":"upgrade","PostDeployErrand":{},"PreDeleteErrand":{}}`
				instanceID = "service-instance-id"
				serviceInstances = fmt.Sprintf(`[{"plan_id": "service-plan-id", "service_instance_id": "%s"}]`, instanceID)

				serviceInstancesHandler.RespondsWith(http.StatusOK, serviceInstances)
				upgradeHandler.RespondsWith(http.StatusAccepted, operationData)
				lastOperationHandler.RespondsOnCall(0, http.StatusOK, `{"state":"in progress"}`)
				lastOperationHandler.RespondsOnCall(1, http.StatusOK, `{"state":"succeeded"}`)
			})

			AfterEach(func() {
				broker.Close()
			})

			It("exits successfully and upgrades the instance", func() {
				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
				Expect(runningTool).To(gbytes.Say("Sleep interval until next attempt: 2s"))
				Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
				Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))
			})

			It("exits successfully when all instances are already up-to-date", func() {
				runningTool := startUpgradeAllInstanceBinary(errandConfig)
				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
				Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))

				By("running upgrade all again")
				upgradeHandler.RespondsWith(http.StatusNoContent, "")
				runningTool = startUpgradeAllInstanceBinary(errandConfig)
				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))

				Expect(runningTool).To(gbytes.Say(`Result: instance already up to date - operation skipped`))
				Expect(runningTool).To(gbytes.Say("Sleep interval until next attempt: 2s"))
				Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
				Expect(runningTool).To(gbytes.Say("Number of skipped operations: 1"))
			})

			It("uses the canary_selection_params when querying canary instances", func() {
				instanceID := "my-instance-id"
				canaryInstanceID := "canary-instance-id"
				canariesList := fmt.Sprintf(`[{"plan_id": "service-plan-id", "service_instance_id": "%s"}]`, canaryInstanceID)
				serviceInstances := fmt.Sprintf(`[{"plan_id": "service-plan-id", "service_instance_id": "%s"}, {"plan_id": "service-plan-id", "service_instance_id": "%s"}]`, instanceID, canaryInstanceID)

				serviceInstancesHandler.WithQueryParams().RespondsWith(http.StatusOK, serviceInstances)
				serviceInstancesHandler.WithQueryParams("foo=bar").RespondsWith(http.StatusOK, canariesList)
				lastOperationHandler.RespondsWith(http.StatusOK, `{"state":"succeeded"}`)

				errandConfig.CanarySelectionParams = map[string]string{"foo": "bar"}
				errandConfig.Canaries = 1

				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
				Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] STARTING CANARIES: 1 canaries`))
				Expect(runningTool).To(gbytes.Say(`\[canary-instance-id] Starting to process service instance`))
				Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED CANARIES`))
				Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
				Expect(runningTool).To(gbytes.Say("Number of successful operations: 2"))
			})

			It("uses the canary_selection_params but returns an error if no instances found but instances exist", func() {
				canariesList := `[]`

				serviceInstancesHandler.WithQueryParams("cf_org=my-org", "cf_space=my-space").RespondsWith(http.StatusOK, canariesList)

				errandConfig.CanarySelectionParams = map[string]string{"cf_org": "my-org", "cf_space": "my-space"}
				errandConfig.Canaries = 1

				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(1))
				Expect(runningTool).To(gbytes.Say("Failed to find a match to the canary selection criteria"))
			})

			It("returns an error if service-instances api responds with a non-200", func() {
				serviceInstancesHandler.RespondsWith(http.StatusInternalServerError, `{"description": "a forced error"}`)

				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(1))
				Expect(runningTool).To(gbytes.Say("error listing service instances"))
				Expect(runningTool).To(gbytes.Say("500"))
			})

			It("exits with a failure and shows a summary message when the upgrade fails", func() {
				lastOperationHandler.RespondsOnCall(1, http.StatusOK, `{"state":"failed"}`)

				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(1))
				Expect(runningTool).To(gbytes.Say("Status: FAILED"))
				Expect(runningTool).To(gbytes.Say(fmt.Sprintf(`Number of service instances that failed to process: 1 \[%s\]`, instanceID)))
			})

			Context("when the attempt limit is reached", func() {
				It("exits with an error reporting the instances that were not upgraded", func() {
					upgradeHandler.RespondsWith(http.StatusConflict, operationData)

					runningTool := startUpgradeAllInstanceBinary(errandConfig)

					Eventually(runningTool, 5*time.Second).Should(gexec.Exit(1))
					Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] Processing all instances. Attempt 1/2`))
					Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] Processing all remaining instances. Attempt 2/2`))
					Expect(runningTool).To(gbytes.Say("Number of busy instances which could not be processed: 1"))
					Expect(runningTool).To(gbytes.Say(fmt.Sprintf("The following instances could not be processed: %s", instanceID)))
				})
			})

			Context("when a service instance plan is updated after upgrade-all starts but before instance upgrade", func() {
				It("uses the new plan for the upgrade", func() {
					serviceInstancesInitialResponse := fmt.Sprintf(`[{"plan_id": "service-plan-id", "service_instance_id": "%s"}]`, instanceID)
					serviceInstancesResponseAfterPlanUpdate := fmt.Sprintf(`[{"plan_id": "service-plan-id-2", "service_instance_id": "%s"}]`, instanceID)

					serviceInstancesHandler.RespondsOnCall(0, http.StatusOK, serviceInstancesInitialResponse)
					serviceInstancesHandler.RespondsOnCall(1, http.StatusOK, serviceInstancesResponseAfterPlanUpdate)

					runningTool := startUpgradeAllInstanceBinary(errandConfig)

					Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
					Expect(runningTool).To(gbytes.Say("Sleep interval until next attempt: 2s"))
					Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
					Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))

					Expect(upgradeHandler.GetRequestForCall(0).Body).To(Equal(`{"plan_id": "service-plan-id-2"}`))
				})
			})

			Context("when a service instance is deleted after upgrade-all starts but before the instance upgrade", func() {
				It("Fetches the latest service instances info and reports a deleted service", func() {
					serviceInstancesHandler.RespondsOnCall(0, http.StatusOK, serviceInstances)
					serviceInstancesHandler.RespondsOnCall(1, http.StatusOK, "[]")

					runningTool := startUpgradeAllInstanceBinary(errandConfig)

					Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
					Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
					Expect(runningTool).To(gbytes.Say("Number of successful operations: 0"))
					Expect(runningTool).To(gbytes.Say("Number of deleted instances before operation could happen: 1"))
				})
			})

			Context("when a service instance refresh fails prior to instance upgrade", func() {
				It("we log failure and carry on with previous data", func() {
					serviceInstancesHandler.RespondsOnCall(0, http.StatusOK, serviceInstances)
					serviceInstancesHandler.RespondsOnCall(1, http.StatusInternalServerError, "oops")

					runningTool := startUpgradeAllInstanceBinary(errandConfig)

					Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))

					Expect(runningTool).To(gbytes.Say("Failed to get refreshed list of instances. Continuing with previously fetched info"))
					Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
					Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))
				})
			})
		})

		Describe("HTTPS Broker", func() {
			var pemCert string

			BeforeEach(func() {
				broker = ghttp.NewTLSServer()
				rawPem := broker.HTTPTestServer.Certificate().Raw
				pemCert = string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: rawPem}))

				errandConfig = config.InstanceIteratorConfig{
					PollingInterval: 1,
					AttemptLimit:    5,
					AttemptInterval: 2,
					MaxInFlight:     1,
					BrokerAPI: config.BrokerAPI{
						URL: broker.URL(),
						Authentication: config.Authentication{
							Basic: config.UserCredentials{
								Username: brokerUsername,
								Password: brokerPassword,
							},
						},
					},
				}

				broker.RouteToHandler(http.MethodGet, "/mgmt/service_instances", ghttp.CombineHandlers(
					ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
					ghttp.RespondWith(http.StatusOK, `[{"service_instance_id":"foo"}]`),
				))

				broker.RouteToHandler(http.MethodPatch, regexp.MustCompile(`/mgmt/service_instances/.*`), ghttp.CombineHandlers(
					ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
					ghttp.RespondWith(http.StatusAccepted, `{"operation":{"task_id":7}}`),
					ghttp.VerifyRequest(http.MethodPatch, ContainSubstring(`/mgmt/service_instances/`), "operation_type=upgrade"),
				))

				broker.RouteToHandler(http.MethodGet, regexp.MustCompile(`/v2/service_instances/.*/last_operation`), ghttp.CombineHandlers(
					ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
					ghttp.RespondWith(http.StatusOK, `{"state":"succeeded"}`),
				))

				broker.HTTPTestServer.Config.ErrorLog = loggerfactory.New(GinkgoWriter, "server", loggerfactory.Flags).New()
			})

			AfterEach(func() {
				broker.Close()
			})

			It("upgrades all instances", func() {
				errandConfig.BrokerAPI.TLS.CACert = pemCert

				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
				Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))
			})

			It("skips ssl cert verification when disabled", func() {
				errandConfig.BrokerAPI.TLS.DisableSSLCertVerification = true

				runningTool := startUpgradeAllInstanceBinary(errandConfig)

				Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
				Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))
			})
		})
	})

	FDescribe("Upgrading via CF", func() {
		var (
			serviceInstances string
			instanceID       string
			broker           *ghttp.Server
			cfApi            *ghttp.Server
			uaaApi           *ghttp.Server
			cfUpgradeHandler *FakeHandler
		)

		const (
			cfUsername = "cf-username"
			cfPassword = "cf-password"
		)

		BeforeEach(func() {
			broker = ghttp.NewServer()
			cfApi = ghttp.NewServer()
			uaaApi = ghttp.NewServer()

			cfInfoHandler := new(FakeHandler)
			cfUpgradeHandler = new(FakeHandler)
			cfServiceHandler := new(FakeHandler)
			serviceInstancesHandler := new(FakeHandler)
			servicePlanHandler := new(FakeHandler)
			uaaAuthenticationHandler := new(FakeHandler)

			errandConfig = config.InstanceIteratorConfig{
				BrokerAPI: config.BrokerAPI{
					URL: broker.URL(),
					Authentication: config.Authentication{
						Basic: config.UserCredentials{
							Username: brokerUsername,
							Password: brokerPassword,
						},
					},
				},
				PollingInterval: 1,
				AttemptInterval: 2,
				AttemptLimit:    2,
				MaxInFlight:     1,
				CF: config.CF{
					URL: cfApi.URL(),
					Authentication: config.Authentication{
						UAA: config.UAAAuthentication{
							URL: uaaApi.URL(),
							UserCredentials: config.UserCredentials{
								Username: cfUsername,
								Password: cfPassword,
							},
						},
					},
					DisableSSLCertVerification: true,
				},
				MaintenanceInfoPresent: true,
			}

			broker.RouteToHandler(http.MethodGet, "/mgmt/service_instances", ghttp.CombineHandlers(
				ghttp.VerifyBasicAuth(brokerUsername, brokerPassword),
				serviceInstancesHandler.Handle,
			))

			instanceID = "service-instance-id"
			serviceInstances = fmt.Sprintf(`[{"plan_id": "service-plan-id", "service_instance_id": "%s"}]`, instanceID)
			serviceInstancesHandler.RespondsWith(http.StatusOK, serviceInstances)

			uaaApi.RouteToHandler(http.MethodPost, regexp.MustCompile(`/oauth/token`), ghttp.CombineHandlers(
				uaaAuthenticationHandler.Handle,
			))
			authenticationResponse := `{ "access_token": "some-random-token", "expires_in": 3600}`
			uaaAuthenticationHandler.RespondsWith(http.StatusOK, authenticationResponse)

			cfApi.RouteToHandler(http.MethodGet, "/v2/info", ghttp.CombineHandlers(
				cfInfoHandler.Handle))

			cfInfoResponse := `{"api_version": "2.139.0","osbapi_version": "2.15"}`
			cfInfoHandler.RespondsWith(http.StatusOK, cfInfoResponse)

			cfApi.RouteToHandler(http.MethodGet, regexp.MustCompile(`/v2/service_plans/.*`), ghttp.CombineHandlers(
				servicePlanHandler.Handle,
			))
			servicePlanResponse := `{ "resources":[{ "entity": { "maintenance_info": { "version": "0.31.0" }}}]}`
			servicePlanHandler.RespondsWith(http.StatusOK, servicePlanResponse)

			cfApi.RouteToHandler(http.MethodPut, regexp.MustCompile(`/v2/service_instances/.*`), ghttp.CombineHandlers(
				ghttp.VerifyRequest(http.MethodPut, ContainSubstring("/v2/service_instances/"), "accepts_incomplete=true"),
				cfUpgradeHandler.Handle,
			))
			upgradeResponse := `{ "entity": { "last_operation": { "type": "update", "state": "in progress" }}}`
			cfUpgradeHandler.RespondsWith(http.StatusAccepted, upgradeResponse)

			cfApi.RouteToHandler(http.MethodGet, regexp.MustCompile(`/v2/service_instances/.*`), ghttp.CombineHandlers(
				cfServiceHandler.Handle,
			))
			serviceResponse := `{ "entity": { "last_operation": { "type": "update", "state": "succeeded" }}}`
			cfServiceHandler.RespondsWith(http.StatusOK, serviceResponse)
		})

		AfterEach(func() {
			broker.Close()
			cfApi.Close()
		})

		It("exits successfully and upgrades the instance", func() {
			runningTool := startUpgradeAllInstanceBinary(errandConfig)

			Eventually(runningTool, 5*time.Second).Should(gexec.Exit(0))
			Expect(runningTool).To(gbytes.Say("Sleep interval until next attempt: 2s"))
			Expect(runningTool).To(gbytes.Say(`\[upgrade\-all\] FINISHED PROCESSING Status: SUCCESS`))
			Expect(runningTool).To(gbytes.Say("Number of successful operations: 1"))

			Expect(cfUpgradeHandler.GetRequestForCall(0).Body).To(MatchJSON(`{"maintenance_info": {"version": "0.31.1"}`))
		})
	})
})

func writeConfigFile(configContent string) string {
	file, err := ioutil.TempFile("", "config")
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	_, err = file.Write([]byte(configContent))
	Expect(err).NotTo(HaveOccurred())

	return file.Name()
}
