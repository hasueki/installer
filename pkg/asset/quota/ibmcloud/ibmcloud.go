package ibmcloud

import (
	"sort"

	ibmcloudprovider "github.com/openshift/cluster-api-provider-ibmcloud/pkg/apis/ibmcloudprovider/v1beta1"
	machineapi "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"

	"github.com/openshift/installer/pkg/quota"
	"github.com/openshift/installer/pkg/types"
)

// Constraints returns a list of quota constraints based on the InstallConfig.
// These constraints can be used to check if there is enough quota for creating a cluster
// for the isntall config.
func Constraints(config *types.InstallConfig, controlPlanes []machineapi.Machine, computes []machineapi.MachineSet) []quota.Constraint {
	ctrplConfigs := make([]*ibmcloudprovider.IBMCloudMachineProviderSpec, len(controlPlanes))
	for i, m := range controlPlanes {
		ctrplConfigs[i] = m.Spec.ProviderSpec.Value.Object.(*ibmcloudprovider.IBMCloudMachineProviderSpec)
	}
	computeReplicas := make([]int64, len(computes))
	computeConfigs := make([]*ibmcloudprovider.IBMCloudMachineProviderSpec, len(computes))
	for i, w := range computes {
		computeReplicas[i] = int64(*w.Spec.Replicas)
		computeConfigs[i] = w.Spec.Template.Spec.ProviderSpec.Value.Object.(*ibmcloudprovider.IBMCloudMachineProviderSpec)
	}

	var ret []quota.Constraint
	for _, gen := range []constraintGenerator{
		network(config),
		// apiExternal(config),
		// apiInternal(config),
		// controlPlane(client, config, ctrplConfigs),
		// compute(client, config, computeReplicas, computeConfigs),
		// others,
	} {
		ret = append(ret, gen()...)
	}
	return aggregate(ret)
}

func aggregate(quotas []quota.Constraint) []quota.Constraint {
	sort.SliceStable(quotas, func(i, j int) bool {
		return quotas[i].Name < quotas[j].Name
	})

	i := 0
	for j := 1; j < len(quotas); j++ {
		if quotas[i].Name == quotas[j].Name && quotas[i].Region == quotas[j].Region {
			quotas[i].Count += quotas[j].Count
		} else {
			i++
			if i != j {
				quotas[i] = quotas[j]
			}
		}
	}
	return quotas[:i+1]
}

// constraintGenerator generates a list of constraints.
type constraintGenerator func() []quota.Constraint

func network(config *types.InstallConfig) func() []quota.Constraint {
	return func() []quota.Constraint {
		net := []quota.Constraint{
			{
				Name:   "is.vpc",
				Region: config.Platform.IBMCloud.Region,
				Count:  10,
			},
			{
				Name:   "is.load-balancer",
				Region: config.Platform.IBMCloud.Region,
				Count:  50,
			},
		}

		return net
	}
}
