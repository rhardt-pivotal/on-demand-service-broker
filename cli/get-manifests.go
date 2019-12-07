// Copyright © 2017 Pivotal Software

package cli

import (
	"fmt"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/brokerinitiator"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/pivotal-cf/on-demand-service-broker/service"
	"github.com/pivotal-cf/on-demand-service-broker/task"
	"github.com/spf13/cobra"
	"os"
)

const InstancePrefix = "service-instance_"

func deploymentName(instanceID string) string {
	return InstancePrefix + instanceID
}


func getInstanceById(lister service.InstanceLister,  id string)(service.Instance, error) {
	instances, err := lister.Instances(map[string]string{})
	if err != nil {
		return service.Instance{}, err
	}
	for _, inst := range instances {
		if inst.GUID == id {
			return inst, nil
		}
	}
	return service.Instance{}, fmt.Errorf("No instance with ID: %s.", id)
}


func GetManifests(conf *config.Config, toolkit *brokerinitiator.ODBToolkit) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {

		logger := toolkit.LoggerFactory.New()
		manifest, found, err := toolkit.BoshClient.GetDeployment(deploymentName(args[0]), logger)
		if err != nil {
			return err
		}
		if !found {
			return broker.NewDeploymentNotFoundError(fmt.Errorf("bosh deployment '%s' not found", deploymentName))
		}

		instance, err := getInstanceById(toolkit.InstanceLister, args[0])
		if err != nil {
			return err
		}


		generateManifestProperties := task.GenerateManifestProperties{
			DeploymentName:  deploymentName(args[0]),
			PlanID:          instance.PlanUniqueID,
			RequestParams:   nil,
			OldManifest:     manifest,
			PreviousPlanID:  &instance.PlanUniqueID,
			SecretsMap:      nil,
			PreviousConfigs: nil,
		}

		newManifest, err := toolkit.ManifestGenerator.GenerateManifest(generateManifestProperties, logger)
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s", manifest)
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "*********************")
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintf(os.Stdout, "%s", newManifest.Manifest)

		return nil
	}
}

func createGetManifestsCommand(conf *config.Config, odbToolkit *brokerinitiator.ODBToolkit) *cobra.Command {
	command := &cobra.Command{
		Use:     "get-manifests [instance-id]",
		Short:   "get the current and to-be-deployed manifests for a PKS cluster",
		Long:    "get the current and to-be-deployed manifests for a PKS cluster",
		Example: "  on-demand-service-broker get-manifests2 8675309-jenny",
		RunE:    GetManifests(conf, odbToolkit),
		Aliases: []string{"get-manifests"},
		Args: 		cobra.ExactArgs(1),
	}

	return command
}