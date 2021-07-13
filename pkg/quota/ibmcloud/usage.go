package ibmcloud

import (
	"context"
	"fmt"

	"github.com/openshift/installer/pkg/asset/quota/ibmcloud"
	"github.com/pkg/errors"
)

func loadUsage(ctx context.Context, client *ibmcloud.Client) ([]record, error) {
	// Get VPC count
	vpcCount, err := client.GetVPCCountForRegion(ctx, "us-east")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get VPC count")
	}

	fmt.Printf("[WIP] loadUsage; VPC Count: %d\n", vpcCount)

	// Get ./

	return nil, nil
}
