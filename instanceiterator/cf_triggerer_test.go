package instanceiterator_test

import (
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-service-broker/broker/services"
	"github.com/pivotal-cf/on-demand-service-broker/cf"
	"github.com/pivotal-cf/on-demand-service-broker/instanceiterator"
	"github.com/pivotal-cf/on-demand-service-broker/instanceiterator/fakes"
	"github.com/pivotal-cf/on-demand-service-broker/service"
	"github.com/pkg/errors"
)

var _ = Describe("CfTriggerer", func() {
	Describe("TriggerOperation", func() {
		var (
			fakeCFClient            *fakes.FakeCFClient
			expectedMaintenanceInfo cf.MaintenanceInfo
		)
		BeforeEach(func() {
			fakeCFClient = new(fakes.FakeCFClient)
			expectedMaintenanceInfo = cf.MaintenanceInfo{
				Version: "2.1.3",
			}
			fakeCFClient.GetPlanByUniqueIDReturns(cf.ServicePlan{
				ServicePlanEntity: cf.ServicePlanEntity{
					MaintenanceInfo: expectedMaintenanceInfo,
				},
			}, nil)
			fakeCFClient.UpgradeServiceInstanceReturns(cf.LastOperation{
				Type:  cf.OperationType("update"),
				State: cf.OperationStateInProgress,
			}, nil)
		})

		It("should call CF to upgrade the service instance", func() {

			cfTriggerer := instanceiterator.NewCFTrigger(fakeCFClient, new(log.Logger))

			triggeredOperation, _ := cfTriggerer.TriggerOperation(service.Instance{
				GUID:         "service-instance-id",
				PlanUniqueID: "plan-id",
			})

			Expect(fakeCFClient.GetPlanByUniqueIDCallCount()).To(Equal(1), "expected to get CF plan")

			Expect(fakeCFClient.UpgradeServiceInstanceCallCount()).To(Equal(1), "expected to call CF upgrade service")
			_, actualMaintenanceInfo, _ := fakeCFClient.UpgradeServiceInstanceArgsForCall(0)
			Expect(actualMaintenanceInfo).To(Equal(expectedMaintenanceInfo))

			Expect(triggeredOperation.Type).To(Equal(services.OperationAccepted))
		})

		It("should call CF to upgrade the service instance", func() {
			fakeCFClient.UpgradeServiceInstanceReturns(cf.LastOperation{
				Type:  cf.OperationType("update"),
				State: cf.OperationStateFailed,
			}, nil)

			cfTriggerer := instanceiterator.NewCFTrigger(fakeCFClient, new(log.Logger))

			triggeredOperation, _ := cfTriggerer.TriggerOperation(service.Instance{
				GUID:         "service-instance-id",
				PlanUniqueID: "plan-id",
			})

			Expect(fakeCFClient.GetPlanByUniqueIDCallCount()).To(Equal(1), "expected to get CF plan")

			Expect(fakeCFClient.UpgradeServiceInstanceCallCount()).To(Equal(1), "expected to call CF upgrade service")
			_, actualMaintenanceInfo, _ := fakeCFClient.UpgradeServiceInstanceArgsForCall(0)
			Expect(actualMaintenanceInfo).To(Equal(expectedMaintenanceInfo))

			Expect(triggeredOperation.Type).To(Equal(services.OperationFailed))
		})

		It("return an error when the CF client cannot get plan by unique ID", func() {
			fakeCFClient.GetPlanByUniqueIDReturns(cf.ServicePlan{}, errors.New("failed to get plan"))
			cfTriggerer := instanceiterator.NewCFTrigger(fakeCFClient, new(log.Logger))

			_, err := cfTriggerer.TriggerOperation(service.Instance{
				GUID:         "service-instance-id",
				PlanUniqueID: "plan-id",
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to get plan"))
		})

		It("return an error when the CF client cannot upgrade service instance", func() {
			fakeCFClient.UpgradeServiceInstanceReturns(cf.LastOperation{}, errors.New("failed to upgrade instance"))
			cfTriggerer := instanceiterator.NewCFTrigger(fakeCFClient, new(log.Logger))

			_, err := cfTriggerer.TriggerOperation(service.Instance{
				GUID:         "service-instance-id",
				PlanUniqueID: "plan-id",
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to upgrade instance"))
		})
	})
})
