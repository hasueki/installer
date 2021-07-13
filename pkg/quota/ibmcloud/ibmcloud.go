package ibmcloud

import (
	"context"
	"fmt"

	"github.com/openshift/installer/pkg/asset/quota/ibmcloud"
	"github.com/openshift/installer/pkg/quota"
	"github.com/pkg/errors"
)

// record stores the data from quota limits and usages.
type record struct {
	Service string
	Name    string

	// this can either be "account", "region", "zone", "vpc"
	Scope string

	Value int64
}

// Load load the quota information for a region. It provides information
// about the usage and limit for each resource quota.
func Load(ctx context.Context, client *ibmcloud.Client, services ...string) ([]quota.Quota, error) {
	limits, err := loadLimits()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load quota limits")
	}

	fmt.Printf("[WIP] Loaded limits: %+v\n", limits)

	usages, err := loadUsage(ctx, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load quota usages")
	}

	fmt.Printf("[WIP] Loaded usages: %+v\n", usages)

	return nil, nil
}

// IsUnauthorized checks if the error is unauthorized.
func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}

	return true
}
