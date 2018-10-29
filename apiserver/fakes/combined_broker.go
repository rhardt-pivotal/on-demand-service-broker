// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"log"
	"sync"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/on-demand-service-broker/apiserver"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/cf"
	"github.com/pivotal-cf/on-demand-service-broker/service"
)

type FakeCombinedBroker struct {
	InstancesStub        func(logger *log.Logger) ([]service.Instance, error)
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct {
		logger *log.Logger
	}
	instancesReturns struct {
		result1 []service.Instance
		result2 error
	}
	instancesReturnsOnCall map[int]struct {
		result1 []service.Instance
		result2 error
	}
	FilteredInstancesStub        func(orgName, spaceName string, logger *log.Logger) ([]service.Instance, error)
	filteredInstancesMutex       sync.RWMutex
	filteredInstancesArgsForCall []struct {
		orgName   string
		spaceName string
		logger    *log.Logger
	}
	filteredInstancesReturns struct {
		result1 []service.Instance
		result2 error
	}
	filteredInstancesReturnsOnCall map[int]struct {
		result1 []service.Instance
		result2 error
	}
	OrphanDeploymentsStub        func(logger *log.Logger) ([]string, error)
	orphanDeploymentsMutex       sync.RWMutex
	orphanDeploymentsArgsForCall []struct {
		logger *log.Logger
	}
	orphanDeploymentsReturns struct {
		result1 []string
		result2 error
	}
	orphanDeploymentsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	UpgradeStub        func(ctx context.Context, instanceID string, updateDetails brokerapi.UpdateDetails, logger *log.Logger) (broker.OperationData, error)
	upgradeMutex       sync.RWMutex
	upgradeArgsForCall []struct {
		ctx           context.Context
		instanceID    string
		updateDetails brokerapi.UpdateDetails
		logger        *log.Logger
	}
	upgradeReturns struct {
		result1 broker.OperationData
		result2 error
	}
	upgradeReturnsOnCall map[int]struct {
		result1 broker.OperationData
		result2 error
	}
	CountInstancesOfPlansStub        func(logger *log.Logger) (map[cf.ServicePlan]int, error)
	countInstancesOfPlansMutex       sync.RWMutex
	countInstancesOfPlansArgsForCall []struct {
		logger *log.Logger
	}
	countInstancesOfPlansReturns struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}
	countInstancesOfPlansReturnsOnCall map[int]struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}
	ServicesStub        func(ctx context.Context) ([]brokerapi.Service, error)
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
		ctx context.Context
	}
	servicesReturns struct {
		result1 []brokerapi.Service
		result2 error
	}
	servicesReturnsOnCall map[int]struct {
		result1 []brokerapi.Service
		result2 error
	}
	ProvisionStub        func(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error)
	provisionMutex       sync.RWMutex
	provisionArgsForCall []struct {
		ctx          context.Context
		instanceID   string
		details      brokerapi.ProvisionDetails
		asyncAllowed bool
	}
	provisionReturns struct {
		result1 brokerapi.ProvisionedServiceSpec
		result2 error
	}
	provisionReturnsOnCall map[int]struct {
		result1 brokerapi.ProvisionedServiceSpec
		result2 error
	}
	DeprovisionStub        func(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error)
	deprovisionMutex       sync.RWMutex
	deprovisionArgsForCall []struct {
		ctx          context.Context
		instanceID   string
		details      brokerapi.DeprovisionDetails
		asyncAllowed bool
	}
	deprovisionReturns struct {
		result1 brokerapi.DeprovisionServiceSpec
		result2 error
	}
	deprovisionReturnsOnCall map[int]struct {
		result1 brokerapi.DeprovisionServiceSpec
		result2 error
	}
	GetInstanceStub        func(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error)
	getInstanceMutex       sync.RWMutex
	getInstanceArgsForCall []struct {
		ctx        context.Context
		instanceID string
	}
	getInstanceReturns struct {
		result1 brokerapi.GetInstanceDetailsSpec
		result2 error
	}
	getInstanceReturnsOnCall map[int]struct {
		result1 brokerapi.GetInstanceDetailsSpec
		result2 error
	}
	BindStub        func(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		ctx          context.Context
		instanceID   string
		bindingID    string
		details      brokerapi.BindDetails
		asyncAllowed bool
	}
	bindReturns struct {
		result1 brokerapi.Binding
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 brokerapi.Binding
		result2 error
	}
	UnbindStub        func(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error)
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		ctx          context.Context
		instanceID   string
		bindingID    string
		details      brokerapi.UnbindDetails
		asyncAllowed bool
	}
	unbindReturns struct {
		result1 brokerapi.UnbindSpec
		result2 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 brokerapi.UnbindSpec
		result2 error
	}
	GetBindingStub        func(ctx context.Context, instanceID, bindingID string) (brokerapi.GetBindingSpec, error)
	getBindingMutex       sync.RWMutex
	getBindingArgsForCall []struct {
		ctx        context.Context
		instanceID string
		bindingID  string
	}
	getBindingReturns struct {
		result1 brokerapi.GetBindingSpec
		result2 error
	}
	getBindingReturnsOnCall map[int]struct {
		result1 brokerapi.GetBindingSpec
		result2 error
	}
	UpdateStub        func(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		ctx          context.Context
		instanceID   string
		details      brokerapi.UpdateDetails
		asyncAllowed bool
	}
	updateReturns struct {
		result1 brokerapi.UpdateServiceSpec
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 brokerapi.UpdateServiceSpec
		result2 error
	}
	LastOperationStub        func(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		ctx        context.Context
		instanceID string
		details    brokerapi.PollDetails
	}
	lastOperationReturns struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	LastBindingOperationStub        func(ctx context.Context, instanceID, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error)
	lastBindingOperationMutex       sync.RWMutex
	lastBindingOperationArgsForCall []struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.PollDetails
	}
	lastBindingOperationReturns struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	lastBindingOperationReturnsOnCall map[int]struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCombinedBroker) Instances(logger *log.Logger) ([]service.Instance, error) {
	fake.instancesMutex.Lock()
	ret, specificReturn := fake.instancesReturnsOnCall[len(fake.instancesArgsForCall)]
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("Instances", []interface{}{logger})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.instancesReturns.result1, fake.instancesReturns.result2
}

func (fake *FakeCombinedBroker) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeCombinedBroker) InstancesArgsForCall(i int) *log.Logger {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return fake.instancesArgsForCall[i].logger
}

func (fake *FakeCombinedBroker) InstancesReturns(result1 []service.Instance, result2 error) {
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) InstancesReturnsOnCall(i int, result1 []service.Instance, result2 error) {
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []service.Instance
			result2 error
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) FilteredInstances(orgName string, spaceName string, logger *log.Logger) ([]service.Instance, error) {
	fake.filteredInstancesMutex.Lock()
	ret, specificReturn := fake.filteredInstancesReturnsOnCall[len(fake.filteredInstancesArgsForCall)]
	fake.filteredInstancesArgsForCall = append(fake.filteredInstancesArgsForCall, struct {
		orgName   string
		spaceName string
		logger    *log.Logger
	}{orgName, spaceName, logger})
	fake.recordInvocation("FilteredInstances", []interface{}{orgName, spaceName, logger})
	fake.filteredInstancesMutex.Unlock()
	if fake.FilteredInstancesStub != nil {
		return fake.FilteredInstancesStub(orgName, spaceName, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.filteredInstancesReturns.result1, fake.filteredInstancesReturns.result2
}

func (fake *FakeCombinedBroker) FilteredInstancesCallCount() int {
	fake.filteredInstancesMutex.RLock()
	defer fake.filteredInstancesMutex.RUnlock()
	return len(fake.filteredInstancesArgsForCall)
}

func (fake *FakeCombinedBroker) FilteredInstancesArgsForCall(i int) (string, string, *log.Logger) {
	fake.filteredInstancesMutex.RLock()
	defer fake.filteredInstancesMutex.RUnlock()
	return fake.filteredInstancesArgsForCall[i].orgName, fake.filteredInstancesArgsForCall[i].spaceName, fake.filteredInstancesArgsForCall[i].logger
}

func (fake *FakeCombinedBroker) FilteredInstancesReturns(result1 []service.Instance, result2 error) {
	fake.FilteredInstancesStub = nil
	fake.filteredInstancesReturns = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) FilteredInstancesReturnsOnCall(i int, result1 []service.Instance, result2 error) {
	fake.FilteredInstancesStub = nil
	if fake.filteredInstancesReturnsOnCall == nil {
		fake.filteredInstancesReturnsOnCall = make(map[int]struct {
			result1 []service.Instance
			result2 error
		})
	}
	fake.filteredInstancesReturnsOnCall[i] = struct {
		result1 []service.Instance
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) OrphanDeployments(logger *log.Logger) ([]string, error) {
	fake.orphanDeploymentsMutex.Lock()
	ret, specificReturn := fake.orphanDeploymentsReturnsOnCall[len(fake.orphanDeploymentsArgsForCall)]
	fake.orphanDeploymentsArgsForCall = append(fake.orphanDeploymentsArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("OrphanDeployments", []interface{}{logger})
	fake.orphanDeploymentsMutex.Unlock()
	if fake.OrphanDeploymentsStub != nil {
		return fake.OrphanDeploymentsStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.orphanDeploymentsReturns.result1, fake.orphanDeploymentsReturns.result2
}

func (fake *FakeCombinedBroker) OrphanDeploymentsCallCount() int {
	fake.orphanDeploymentsMutex.RLock()
	defer fake.orphanDeploymentsMutex.RUnlock()
	return len(fake.orphanDeploymentsArgsForCall)
}

func (fake *FakeCombinedBroker) OrphanDeploymentsArgsForCall(i int) *log.Logger {
	fake.orphanDeploymentsMutex.RLock()
	defer fake.orphanDeploymentsMutex.RUnlock()
	return fake.orphanDeploymentsArgsForCall[i].logger
}

func (fake *FakeCombinedBroker) OrphanDeploymentsReturns(result1 []string, result2 error) {
	fake.OrphanDeploymentsStub = nil
	fake.orphanDeploymentsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) OrphanDeploymentsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.OrphanDeploymentsStub = nil
	if fake.orphanDeploymentsReturnsOnCall == nil {
		fake.orphanDeploymentsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.orphanDeploymentsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Upgrade(ctx context.Context, instanceID string, updateDetails brokerapi.UpdateDetails, logger *log.Logger) (broker.OperationData, error) {
	fake.upgradeMutex.Lock()
	ret, specificReturn := fake.upgradeReturnsOnCall[len(fake.upgradeArgsForCall)]
	fake.upgradeArgsForCall = append(fake.upgradeArgsForCall, struct {
		ctx           context.Context
		instanceID    string
		updateDetails brokerapi.UpdateDetails
		logger        *log.Logger
	}{ctx, instanceID, updateDetails, logger})
	fake.recordInvocation("Upgrade", []interface{}{ctx, instanceID, updateDetails, logger})
	fake.upgradeMutex.Unlock()
	if fake.UpgradeStub != nil {
		return fake.UpgradeStub(ctx, instanceID, updateDetails, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.upgradeReturns.result1, fake.upgradeReturns.result2
}

func (fake *FakeCombinedBroker) UpgradeCallCount() int {
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	return len(fake.upgradeArgsForCall)
}

func (fake *FakeCombinedBroker) UpgradeArgsForCall(i int) (context.Context, string, brokerapi.UpdateDetails, *log.Logger) {
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	return fake.upgradeArgsForCall[i].ctx, fake.upgradeArgsForCall[i].instanceID, fake.upgradeArgsForCall[i].updateDetails, fake.upgradeArgsForCall[i].logger
}

func (fake *FakeCombinedBroker) UpgradeReturns(result1 broker.OperationData, result2 error) {
	fake.UpgradeStub = nil
	fake.upgradeReturns = struct {
		result1 broker.OperationData
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) UpgradeReturnsOnCall(i int, result1 broker.OperationData, result2 error) {
	fake.UpgradeStub = nil
	if fake.upgradeReturnsOnCall == nil {
		fake.upgradeReturnsOnCall = make(map[int]struct {
			result1 broker.OperationData
			result2 error
		})
	}
	fake.upgradeReturnsOnCall[i] = struct {
		result1 broker.OperationData
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) CountInstancesOfPlans(logger *log.Logger) (map[cf.ServicePlan]int, error) {
	fake.countInstancesOfPlansMutex.Lock()
	ret, specificReturn := fake.countInstancesOfPlansReturnsOnCall[len(fake.countInstancesOfPlansArgsForCall)]
	fake.countInstancesOfPlansArgsForCall = append(fake.countInstancesOfPlansArgsForCall, struct {
		logger *log.Logger
	}{logger})
	fake.recordInvocation("CountInstancesOfPlans", []interface{}{logger})
	fake.countInstancesOfPlansMutex.Unlock()
	if fake.CountInstancesOfPlansStub != nil {
		return fake.CountInstancesOfPlansStub(logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.countInstancesOfPlansReturns.result1, fake.countInstancesOfPlansReturns.result2
}

func (fake *FakeCombinedBroker) CountInstancesOfPlansCallCount() int {
	fake.countInstancesOfPlansMutex.RLock()
	defer fake.countInstancesOfPlansMutex.RUnlock()
	return len(fake.countInstancesOfPlansArgsForCall)
}

func (fake *FakeCombinedBroker) CountInstancesOfPlansArgsForCall(i int) *log.Logger {
	fake.countInstancesOfPlansMutex.RLock()
	defer fake.countInstancesOfPlansMutex.RUnlock()
	return fake.countInstancesOfPlansArgsForCall[i].logger
}

func (fake *FakeCombinedBroker) CountInstancesOfPlansReturns(result1 map[cf.ServicePlan]int, result2 error) {
	fake.CountInstancesOfPlansStub = nil
	fake.countInstancesOfPlansReturns = struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) CountInstancesOfPlansReturnsOnCall(i int, result1 map[cf.ServicePlan]int, result2 error) {
	fake.CountInstancesOfPlansStub = nil
	if fake.countInstancesOfPlansReturnsOnCall == nil {
		fake.countInstancesOfPlansReturnsOnCall = make(map[int]struct {
			result1 map[cf.ServicePlan]int
			result2 error
		})
	}
	fake.countInstancesOfPlansReturnsOnCall[i] = struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Services(ctx context.Context) ([]brokerapi.Service, error) {
	fake.servicesMutex.Lock()
	ret, specificReturn := fake.servicesReturnsOnCall[len(fake.servicesArgsForCall)]
	fake.servicesArgsForCall = append(fake.servicesArgsForCall, struct {
		ctx context.Context
	}{ctx})
	fake.recordInvocation("Services", []interface{}{ctx})
	fake.servicesMutex.Unlock()
	if fake.ServicesStub != nil {
		return fake.ServicesStub(ctx)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.servicesReturns.result1, fake.servicesReturns.result2
}

func (fake *FakeCombinedBroker) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeCombinedBroker) ServicesArgsForCall(i int) context.Context {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return fake.servicesArgsForCall[i].ctx
}

func (fake *FakeCombinedBroker) ServicesReturns(result1 []brokerapi.Service, result2 error) {
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []brokerapi.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) ServicesReturnsOnCall(i int, result1 []brokerapi.Service, result2 error) {
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 []brokerapi.Service
			result2 error
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 []brokerapi.Service
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	fake.provisionMutex.Lock()
	ret, specificReturn := fake.provisionReturnsOnCall[len(fake.provisionArgsForCall)]
	fake.provisionArgsForCall = append(fake.provisionArgsForCall, struct {
		ctx          context.Context
		instanceID   string
		details      brokerapi.ProvisionDetails
		asyncAllowed bool
	}{ctx, instanceID, details, asyncAllowed})
	fake.recordInvocation("Provision", []interface{}{ctx, instanceID, details, asyncAllowed})
	fake.provisionMutex.Unlock()
	if fake.ProvisionStub != nil {
		return fake.ProvisionStub(ctx, instanceID, details, asyncAllowed)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.provisionReturns.result1, fake.provisionReturns.result2
}

func (fake *FakeCombinedBroker) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeCombinedBroker) ProvisionArgsForCall(i int) (context.Context, string, brokerapi.ProvisionDetails, bool) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].ctx, fake.provisionArgsForCall[i].instanceID, fake.provisionArgsForCall[i].details, fake.provisionArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBroker) ProvisionReturns(result1 brokerapi.ProvisionedServiceSpec, result2 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 brokerapi.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) ProvisionReturnsOnCall(i int, result1 brokerapi.ProvisionedServiceSpec, result2 error) {
	fake.ProvisionStub = nil
	if fake.provisionReturnsOnCall == nil {
		fake.provisionReturnsOnCall = make(map[int]struct {
			result1 brokerapi.ProvisionedServiceSpec
			result2 error
		})
	}
	fake.provisionReturnsOnCall[i] = struct {
		result1 brokerapi.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	fake.deprovisionMutex.Lock()
	ret, specificReturn := fake.deprovisionReturnsOnCall[len(fake.deprovisionArgsForCall)]
	fake.deprovisionArgsForCall = append(fake.deprovisionArgsForCall, struct {
		ctx          context.Context
		instanceID   string
		details      brokerapi.DeprovisionDetails
		asyncAllowed bool
	}{ctx, instanceID, details, asyncAllowed})
	fake.recordInvocation("Deprovision", []interface{}{ctx, instanceID, details, asyncAllowed})
	fake.deprovisionMutex.Unlock()
	if fake.DeprovisionStub != nil {
		return fake.DeprovisionStub(ctx, instanceID, details, asyncAllowed)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.deprovisionReturns.result1, fake.deprovisionReturns.result2
}

func (fake *FakeCombinedBroker) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeCombinedBroker) DeprovisionArgsForCall(i int) (context.Context, string, brokerapi.DeprovisionDetails, bool) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].ctx, fake.deprovisionArgsForCall[i].instanceID, fake.deprovisionArgsForCall[i].details, fake.deprovisionArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBroker) DeprovisionReturns(result1 brokerapi.DeprovisionServiceSpec, result2 error) {
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 brokerapi.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) DeprovisionReturnsOnCall(i int, result1 brokerapi.DeprovisionServiceSpec, result2 error) {
	fake.DeprovisionStub = nil
	if fake.deprovisionReturnsOnCall == nil {
		fake.deprovisionReturnsOnCall = make(map[int]struct {
			result1 brokerapi.DeprovisionServiceSpec
			result2 error
		})
	}
	fake.deprovisionReturnsOnCall[i] = struct {
		result1 brokerapi.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) GetInstance(ctx context.Context, instanceID string) (brokerapi.GetInstanceDetailsSpec, error) {
	fake.getInstanceMutex.Lock()
	ret, specificReturn := fake.getInstanceReturnsOnCall[len(fake.getInstanceArgsForCall)]
	fake.getInstanceArgsForCall = append(fake.getInstanceArgsForCall, struct {
		ctx        context.Context
		instanceID string
	}{ctx, instanceID})
	fake.recordInvocation("GetInstance", []interface{}{ctx, instanceID})
	fake.getInstanceMutex.Unlock()
	if fake.GetInstanceStub != nil {
		return fake.GetInstanceStub(ctx, instanceID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getInstanceReturns.result1, fake.getInstanceReturns.result2
}

func (fake *FakeCombinedBroker) GetInstanceCallCount() int {
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	return len(fake.getInstanceArgsForCall)
}

func (fake *FakeCombinedBroker) GetInstanceArgsForCall(i int) (context.Context, string) {
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	return fake.getInstanceArgsForCall[i].ctx, fake.getInstanceArgsForCall[i].instanceID
}

func (fake *FakeCombinedBroker) GetInstanceReturns(result1 brokerapi.GetInstanceDetailsSpec, result2 error) {
	fake.GetInstanceStub = nil
	fake.getInstanceReturns = struct {
		result1 brokerapi.GetInstanceDetailsSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) GetInstanceReturnsOnCall(i int, result1 brokerapi.GetInstanceDetailsSpec, result2 error) {
	fake.GetInstanceStub = nil
	if fake.getInstanceReturnsOnCall == nil {
		fake.getInstanceReturnsOnCall = make(map[int]struct {
			result1 brokerapi.GetInstanceDetailsSpec
			result2 error
		})
	}
	fake.getInstanceReturnsOnCall[i] = struct {
		result1 brokerapi.GetInstanceDetailsSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails, asyncAllowed bool) (brokerapi.Binding, error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		ctx          context.Context
		instanceID   string
		bindingID    string
		details      brokerapi.BindDetails
		asyncAllowed bool
	}{ctx, instanceID, bindingID, details, asyncAllowed})
	fake.recordInvocation("Bind", []interface{}{ctx, instanceID, bindingID, details, asyncAllowed})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(ctx, instanceID, bindingID, details, asyncAllowed)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindReturns.result1, fake.bindReturns.result2
}

func (fake *FakeCombinedBroker) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeCombinedBroker) BindArgsForCall(i int) (context.Context, string, string, brokerapi.BindDetails, bool) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].ctx, fake.bindArgsForCall[i].instanceID, fake.bindArgsForCall[i].bindingID, fake.bindArgsForCall[i].details, fake.bindArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBroker) BindReturns(result1 brokerapi.Binding, result2 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 brokerapi.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) BindReturnsOnCall(i int, result1 brokerapi.Binding, result2 error) {
	fake.BindStub = nil
	if fake.bindReturnsOnCall == nil {
		fake.bindReturnsOnCall = make(map[int]struct {
			result1 brokerapi.Binding
			result2 error
		})
	}
	fake.bindReturnsOnCall[i] = struct {
		result1 brokerapi.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Unbind(ctx context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails, asyncAllowed bool) (brokerapi.UnbindSpec, error) {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		ctx          context.Context
		instanceID   string
		bindingID    string
		details      brokerapi.UnbindDetails
		asyncAllowed bool
	}{ctx, instanceID, bindingID, details, asyncAllowed})
	fake.recordInvocation("Unbind", []interface{}{ctx, instanceID, bindingID, details, asyncAllowed})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(ctx, instanceID, bindingID, details, asyncAllowed)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.unbindReturns.result1, fake.unbindReturns.result2
}

func (fake *FakeCombinedBroker) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeCombinedBroker) UnbindArgsForCall(i int) (context.Context, string, string, brokerapi.UnbindDetails, bool) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].ctx, fake.unbindArgsForCall[i].instanceID, fake.unbindArgsForCall[i].bindingID, fake.unbindArgsForCall[i].details, fake.unbindArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBroker) UnbindReturns(result1 brokerapi.UnbindSpec, result2 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 brokerapi.UnbindSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) UnbindReturnsOnCall(i int, result1 brokerapi.UnbindSpec, result2 error) {
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 brokerapi.UnbindSpec
			result2 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 brokerapi.UnbindSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) GetBinding(ctx context.Context, instanceID string, bindingID string) (brokerapi.GetBindingSpec, error) {
	fake.getBindingMutex.Lock()
	ret, specificReturn := fake.getBindingReturnsOnCall[len(fake.getBindingArgsForCall)]
	fake.getBindingArgsForCall = append(fake.getBindingArgsForCall, struct {
		ctx        context.Context
		instanceID string
		bindingID  string
	}{ctx, instanceID, bindingID})
	fake.recordInvocation("GetBinding", []interface{}{ctx, instanceID, bindingID})
	fake.getBindingMutex.Unlock()
	if fake.GetBindingStub != nil {
		return fake.GetBindingStub(ctx, instanceID, bindingID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getBindingReturns.result1, fake.getBindingReturns.result2
}

func (fake *FakeCombinedBroker) GetBindingCallCount() int {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return len(fake.getBindingArgsForCall)
}

func (fake *FakeCombinedBroker) GetBindingArgsForCall(i int) (context.Context, string, string) {
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	return fake.getBindingArgsForCall[i].ctx, fake.getBindingArgsForCall[i].instanceID, fake.getBindingArgsForCall[i].bindingID
}

func (fake *FakeCombinedBroker) GetBindingReturns(result1 brokerapi.GetBindingSpec, result2 error) {
	fake.GetBindingStub = nil
	fake.getBindingReturns = struct {
		result1 brokerapi.GetBindingSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) GetBindingReturnsOnCall(i int, result1 brokerapi.GetBindingSpec, result2 error) {
	fake.GetBindingStub = nil
	if fake.getBindingReturnsOnCall == nil {
		fake.getBindingReturnsOnCall = make(map[int]struct {
			result1 brokerapi.GetBindingSpec
			result2 error
		})
	}
	fake.getBindingReturnsOnCall[i] = struct {
		result1 brokerapi.GetBindingSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		ctx          context.Context
		instanceID   string
		details      brokerapi.UpdateDetails
		asyncAllowed bool
	}{ctx, instanceID, details, asyncAllowed})
	fake.recordInvocation("Update", []interface{}{ctx, instanceID, details, asyncAllowed})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(ctx, instanceID, details, asyncAllowed)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.updateReturns.result1, fake.updateReturns.result2
}

func (fake *FakeCombinedBroker) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeCombinedBroker) UpdateArgsForCall(i int) (context.Context, string, brokerapi.UpdateDetails, bool) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].ctx, fake.updateArgsForCall[i].instanceID, fake.updateArgsForCall[i].details, fake.updateArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBroker) UpdateReturns(result1 brokerapi.UpdateServiceSpec, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 brokerapi.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) UpdateReturnsOnCall(i int, result1 brokerapi.UpdateServiceSpec, result2 error) {
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 brokerapi.UpdateServiceSpec
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 brokerapi.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) LastOperation(ctx context.Context, instanceID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		ctx        context.Context
		instanceID string
		details    brokerapi.PollDetails
	}{ctx, instanceID, details})
	fake.recordInvocation("LastOperation", []interface{}{ctx, instanceID, details})
	fake.lastOperationMutex.Unlock()
	if fake.LastOperationStub != nil {
		return fake.LastOperationStub(ctx, instanceID, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastOperationReturns.result1, fake.lastOperationReturns.result2
}

func (fake *FakeCombinedBroker) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *FakeCombinedBroker) LastOperationArgsForCall(i int) (context.Context, string, brokerapi.PollDetails) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return fake.lastOperationArgsForCall[i].ctx, fake.lastOperationArgsForCall[i].instanceID, fake.lastOperationArgsForCall[i].details
}

func (fake *FakeCombinedBroker) LastOperationReturns(result1 brokerapi.LastOperation, result2 error) {
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) LastOperationReturnsOnCall(i int, result1 brokerapi.LastOperation, result2 error) {
	fake.LastOperationStub = nil
	if fake.lastOperationReturnsOnCall == nil {
		fake.lastOperationReturnsOnCall = make(map[int]struct {
			result1 brokerapi.LastOperation
			result2 error
		})
	}
	fake.lastOperationReturnsOnCall[i] = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) LastBindingOperation(ctx context.Context, instanceID string, bindingID string, details brokerapi.PollDetails) (brokerapi.LastOperation, error) {
	fake.lastBindingOperationMutex.Lock()
	ret, specificReturn := fake.lastBindingOperationReturnsOnCall[len(fake.lastBindingOperationArgsForCall)]
	fake.lastBindingOperationArgsForCall = append(fake.lastBindingOperationArgsForCall, struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.PollDetails
	}{ctx, instanceID, bindingID, details})
	fake.recordInvocation("LastBindingOperation", []interface{}{ctx, instanceID, bindingID, details})
	fake.lastBindingOperationMutex.Unlock()
	if fake.LastBindingOperationStub != nil {
		return fake.LastBindingOperationStub(ctx, instanceID, bindingID, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastBindingOperationReturns.result1, fake.lastBindingOperationReturns.result2
}

func (fake *FakeCombinedBroker) LastBindingOperationCallCount() int {
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	return len(fake.lastBindingOperationArgsForCall)
}

func (fake *FakeCombinedBroker) LastBindingOperationArgsForCall(i int) (context.Context, string, string, brokerapi.PollDetails) {
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	return fake.lastBindingOperationArgsForCall[i].ctx, fake.lastBindingOperationArgsForCall[i].instanceID, fake.lastBindingOperationArgsForCall[i].bindingID, fake.lastBindingOperationArgsForCall[i].details
}

func (fake *FakeCombinedBroker) LastBindingOperationReturns(result1 brokerapi.LastOperation, result2 error) {
	fake.LastBindingOperationStub = nil
	fake.lastBindingOperationReturns = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) LastBindingOperationReturnsOnCall(i int, result1 brokerapi.LastOperation, result2 error) {
	fake.LastBindingOperationStub = nil
	if fake.lastBindingOperationReturnsOnCall == nil {
		fake.lastBindingOperationReturnsOnCall = make(map[int]struct {
			result1 brokerapi.LastOperation
			result2 error
		})
	}
	fake.lastBindingOperationReturnsOnCall[i] = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBroker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	fake.filteredInstancesMutex.RLock()
	defer fake.filteredInstancesMutex.RUnlock()
	fake.orphanDeploymentsMutex.RLock()
	defer fake.orphanDeploymentsMutex.RUnlock()
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	fake.countInstancesOfPlansMutex.RLock()
	defer fake.countInstancesOfPlansMutex.RUnlock()
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	fake.getInstanceMutex.RLock()
	defer fake.getInstanceMutex.RUnlock()
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.getBindingMutex.RLock()
	defer fake.getBindingMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	fake.lastBindingOperationMutex.RLock()
	defer fake.lastBindingOperationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCombinedBroker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ apiserver.CombinedBroker = new(FakeCombinedBroker)
