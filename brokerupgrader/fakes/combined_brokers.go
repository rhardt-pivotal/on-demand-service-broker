// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"log"
	"sync"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/brokerupgrader"
	"github.com/pivotal-cf/on-demand-service-broker/cf"
)

type FakeCombinedBrokers struct {
	InstancesStub        func(logger *log.Logger) ([]string, error)
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct {
		logger *log.Logger
	}
	instancesReturns struct {
		result1 []string
		result2 error
	}
	instancesReturnsOnCall map[int]struct {
		result1 []string
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
	UpgradeStub        func(ctx context.Context, instanceID string, logger *log.Logger) (broker.OperationData, error)
	upgradeMutex       sync.RWMutex
	upgradeArgsForCall []struct {
		ctx        context.Context
		instanceID string
		logger     *log.Logger
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
	ServicesStub        func(ctx context.Context) []brokerapi.Service
	servicesMutex       sync.RWMutex
	servicesArgsForCall []struct {
		ctx context.Context
	}
	servicesReturns struct {
		result1 []brokerapi.Service
	}
	servicesReturnsOnCall map[int]struct {
		result1 []brokerapi.Service
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
	BindStub        func(ctx context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error)
	bindMutex       sync.RWMutex
	bindArgsForCall []struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
	}
	bindReturns struct {
		result1 brokerapi.Binding
		result2 error
	}
	bindReturnsOnCall map[int]struct {
		result1 brokerapi.Binding
		result2 error
	}
	UnbindStub        func(ctx context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error
	unbindMutex       sync.RWMutex
	unbindArgsForCall []struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.UnbindDetails
	}
	unbindReturns struct {
		result1 error
	}
	unbindReturnsOnCall map[int]struct {
		result1 error
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
	LastOperationStub        func(ctx context.Context, instanceID, operationData string) (brokerapi.LastOperation, error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		ctx           context.Context
		instanceID    string
		operationData string
	}
	lastOperationReturns struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCombinedBrokers) Instances(logger *log.Logger) ([]string, error) {
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

func (fake *FakeCombinedBrokers) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeCombinedBrokers) InstancesArgsForCall(i int) *log.Logger {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return fake.instancesArgsForCall[i].logger
}

func (fake *FakeCombinedBrokers) InstancesReturns(result1 []string, result2 error) {
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) InstancesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) OrphanDeployments(logger *log.Logger) ([]string, error) {
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

func (fake *FakeCombinedBrokers) OrphanDeploymentsCallCount() int {
	fake.orphanDeploymentsMutex.RLock()
	defer fake.orphanDeploymentsMutex.RUnlock()
	return len(fake.orphanDeploymentsArgsForCall)
}

func (fake *FakeCombinedBrokers) OrphanDeploymentsArgsForCall(i int) *log.Logger {
	fake.orphanDeploymentsMutex.RLock()
	defer fake.orphanDeploymentsMutex.RUnlock()
	return fake.orphanDeploymentsArgsForCall[i].logger
}

func (fake *FakeCombinedBrokers) OrphanDeploymentsReturns(result1 []string, result2 error) {
	fake.OrphanDeploymentsStub = nil
	fake.orphanDeploymentsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) OrphanDeploymentsReturnsOnCall(i int, result1 []string, result2 error) {
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

func (fake *FakeCombinedBrokers) Upgrade(ctx context.Context, instanceID string, logger *log.Logger) (broker.OperationData, error) {
	fake.upgradeMutex.Lock()
	ret, specificReturn := fake.upgradeReturnsOnCall[len(fake.upgradeArgsForCall)]
	fake.upgradeArgsForCall = append(fake.upgradeArgsForCall, struct {
		ctx        context.Context
		instanceID string
		logger     *log.Logger
	}{ctx, instanceID, logger})
	fake.recordInvocation("Upgrade", []interface{}{ctx, instanceID, logger})
	fake.upgradeMutex.Unlock()
	if fake.UpgradeStub != nil {
		return fake.UpgradeStub(ctx, instanceID, logger)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.upgradeReturns.result1, fake.upgradeReturns.result2
}

func (fake *FakeCombinedBrokers) UpgradeCallCount() int {
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	return len(fake.upgradeArgsForCall)
}

func (fake *FakeCombinedBrokers) UpgradeArgsForCall(i int) (context.Context, string, *log.Logger) {
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	return fake.upgradeArgsForCall[i].ctx, fake.upgradeArgsForCall[i].instanceID, fake.upgradeArgsForCall[i].logger
}

func (fake *FakeCombinedBrokers) UpgradeReturns(result1 broker.OperationData, result2 error) {
	fake.UpgradeStub = nil
	fake.upgradeReturns = struct {
		result1 broker.OperationData
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) UpgradeReturnsOnCall(i int, result1 broker.OperationData, result2 error) {
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

func (fake *FakeCombinedBrokers) CountInstancesOfPlans(logger *log.Logger) (map[cf.ServicePlan]int, error) {
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

func (fake *FakeCombinedBrokers) CountInstancesOfPlansCallCount() int {
	fake.countInstancesOfPlansMutex.RLock()
	defer fake.countInstancesOfPlansMutex.RUnlock()
	return len(fake.countInstancesOfPlansArgsForCall)
}

func (fake *FakeCombinedBrokers) CountInstancesOfPlansArgsForCall(i int) *log.Logger {
	fake.countInstancesOfPlansMutex.RLock()
	defer fake.countInstancesOfPlansMutex.RUnlock()
	return fake.countInstancesOfPlansArgsForCall[i].logger
}

func (fake *FakeCombinedBrokers) CountInstancesOfPlansReturns(result1 map[cf.ServicePlan]int, result2 error) {
	fake.CountInstancesOfPlansStub = nil
	fake.countInstancesOfPlansReturns = struct {
		result1 map[cf.ServicePlan]int
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) CountInstancesOfPlansReturnsOnCall(i int, result1 map[cf.ServicePlan]int, result2 error) {
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

func (fake *FakeCombinedBrokers) Services(ctx context.Context) []brokerapi.Service {
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
		return ret.result1
	}
	return fake.servicesReturns.result1
}

func (fake *FakeCombinedBrokers) ServicesCallCount() int {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return len(fake.servicesArgsForCall)
}

func (fake *FakeCombinedBrokers) ServicesArgsForCall(i int) context.Context {
	fake.servicesMutex.RLock()
	defer fake.servicesMutex.RUnlock()
	return fake.servicesArgsForCall[i].ctx
}

func (fake *FakeCombinedBrokers) ServicesReturns(result1 []brokerapi.Service) {
	fake.ServicesStub = nil
	fake.servicesReturns = struct {
		result1 []brokerapi.Service
	}{result1}
}

func (fake *FakeCombinedBrokers) ServicesReturnsOnCall(i int, result1 []brokerapi.Service) {
	fake.ServicesStub = nil
	if fake.servicesReturnsOnCall == nil {
		fake.servicesReturnsOnCall = make(map[int]struct {
			result1 []brokerapi.Service
		})
	}
	fake.servicesReturnsOnCall[i] = struct {
		result1 []brokerapi.Service
	}{result1}
}

func (fake *FakeCombinedBrokers) Provision(ctx context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
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

func (fake *FakeCombinedBrokers) ProvisionCallCount() int {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return len(fake.provisionArgsForCall)
}

func (fake *FakeCombinedBrokers) ProvisionArgsForCall(i int) (context.Context, string, brokerapi.ProvisionDetails, bool) {
	fake.provisionMutex.RLock()
	defer fake.provisionMutex.RUnlock()
	return fake.provisionArgsForCall[i].ctx, fake.provisionArgsForCall[i].instanceID, fake.provisionArgsForCall[i].details, fake.provisionArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBrokers) ProvisionReturns(result1 brokerapi.ProvisionedServiceSpec, result2 error) {
	fake.ProvisionStub = nil
	fake.provisionReturns = struct {
		result1 brokerapi.ProvisionedServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) ProvisionReturnsOnCall(i int, result1 brokerapi.ProvisionedServiceSpec, result2 error) {
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

func (fake *FakeCombinedBrokers) Deprovision(ctx context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
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

func (fake *FakeCombinedBrokers) DeprovisionCallCount() int {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return len(fake.deprovisionArgsForCall)
}

func (fake *FakeCombinedBrokers) DeprovisionArgsForCall(i int) (context.Context, string, brokerapi.DeprovisionDetails, bool) {
	fake.deprovisionMutex.RLock()
	defer fake.deprovisionMutex.RUnlock()
	return fake.deprovisionArgsForCall[i].ctx, fake.deprovisionArgsForCall[i].instanceID, fake.deprovisionArgsForCall[i].details, fake.deprovisionArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBrokers) DeprovisionReturns(result1 brokerapi.DeprovisionServiceSpec, result2 error) {
	fake.DeprovisionStub = nil
	fake.deprovisionReturns = struct {
		result1 brokerapi.DeprovisionServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) DeprovisionReturnsOnCall(i int, result1 brokerapi.DeprovisionServiceSpec, result2 error) {
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

func (fake *FakeCombinedBrokers) Bind(ctx context.Context, instanceID string, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	fake.bindMutex.Lock()
	ret, specificReturn := fake.bindReturnsOnCall[len(fake.bindArgsForCall)]
	fake.bindArgsForCall = append(fake.bindArgsForCall, struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.BindDetails
	}{ctx, instanceID, bindingID, details})
	fake.recordInvocation("Bind", []interface{}{ctx, instanceID, bindingID, details})
	fake.bindMutex.Unlock()
	if fake.BindStub != nil {
		return fake.BindStub(ctx, instanceID, bindingID, details)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.bindReturns.result1, fake.bindReturns.result2
}

func (fake *FakeCombinedBrokers) BindCallCount() int {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return len(fake.bindArgsForCall)
}

func (fake *FakeCombinedBrokers) BindArgsForCall(i int) (context.Context, string, string, brokerapi.BindDetails) {
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	return fake.bindArgsForCall[i].ctx, fake.bindArgsForCall[i].instanceID, fake.bindArgsForCall[i].bindingID, fake.bindArgsForCall[i].details
}

func (fake *FakeCombinedBrokers) BindReturns(result1 brokerapi.Binding, result2 error) {
	fake.BindStub = nil
	fake.bindReturns = struct {
		result1 brokerapi.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) BindReturnsOnCall(i int, result1 brokerapi.Binding, result2 error) {
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

func (fake *FakeCombinedBrokers) Unbind(ctx context.Context, instanceID string, bindingID string, details brokerapi.UnbindDetails) error {
	fake.unbindMutex.Lock()
	ret, specificReturn := fake.unbindReturnsOnCall[len(fake.unbindArgsForCall)]
	fake.unbindArgsForCall = append(fake.unbindArgsForCall, struct {
		ctx        context.Context
		instanceID string
		bindingID  string
		details    brokerapi.UnbindDetails
	}{ctx, instanceID, bindingID, details})
	fake.recordInvocation("Unbind", []interface{}{ctx, instanceID, bindingID, details})
	fake.unbindMutex.Unlock()
	if fake.UnbindStub != nil {
		return fake.UnbindStub(ctx, instanceID, bindingID, details)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.unbindReturns.result1
}

func (fake *FakeCombinedBrokers) UnbindCallCount() int {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return len(fake.unbindArgsForCall)
}

func (fake *FakeCombinedBrokers) UnbindArgsForCall(i int) (context.Context, string, string, brokerapi.UnbindDetails) {
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	return fake.unbindArgsForCall[i].ctx, fake.unbindArgsForCall[i].instanceID, fake.unbindArgsForCall[i].bindingID, fake.unbindArgsForCall[i].details
}

func (fake *FakeCombinedBrokers) UnbindReturns(result1 error) {
	fake.UnbindStub = nil
	fake.unbindReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCombinedBrokers) UnbindReturnsOnCall(i int, result1 error) {
	fake.UnbindStub = nil
	if fake.unbindReturnsOnCall == nil {
		fake.unbindReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unbindReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCombinedBrokers) Update(ctx context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
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

func (fake *FakeCombinedBrokers) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeCombinedBrokers) UpdateArgsForCall(i int) (context.Context, string, brokerapi.UpdateDetails, bool) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].ctx, fake.updateArgsForCall[i].instanceID, fake.updateArgsForCall[i].details, fake.updateArgsForCall[i].asyncAllowed
}

func (fake *FakeCombinedBrokers) UpdateReturns(result1 brokerapi.UpdateServiceSpec, result2 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 brokerapi.UpdateServiceSpec
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) UpdateReturnsOnCall(i int, result1 brokerapi.UpdateServiceSpec, result2 error) {
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

func (fake *FakeCombinedBrokers) LastOperation(ctx context.Context, instanceID string, operationData string) (brokerapi.LastOperation, error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		ctx           context.Context
		instanceID    string
		operationData string
	}{ctx, instanceID, operationData})
	fake.recordInvocation("LastOperation", []interface{}{ctx, instanceID, operationData})
	fake.lastOperationMutex.Unlock()
	if fake.LastOperationStub != nil {
		return fake.LastOperationStub(ctx, instanceID, operationData)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastOperationReturns.result1, fake.lastOperationReturns.result2
}

func (fake *FakeCombinedBrokers) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *FakeCombinedBrokers) LastOperationArgsForCall(i int) (context.Context, string, string) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return fake.lastOperationArgsForCall[i].ctx, fake.lastOperationArgsForCall[i].instanceID, fake.lastOperationArgsForCall[i].operationData
}

func (fake *FakeCombinedBrokers) LastOperationReturns(result1 brokerapi.LastOperation, result2 error) {
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeCombinedBrokers) LastOperationReturnsOnCall(i int, result1 brokerapi.LastOperation, result2 error) {
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

func (fake *FakeCombinedBrokers) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
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
	fake.bindMutex.RLock()
	defer fake.bindMutex.RUnlock()
	fake.unbindMutex.RLock()
	defer fake.unbindMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCombinedBrokers) recordInvocation(key string, args []interface{}) {
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

var _ brokerupgrader.CombinedBrokers = new(FakeCombinedBrokers)
