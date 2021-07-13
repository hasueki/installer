package ibmcloud

import (
	"context"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

// Client makes calls to the IBM Cloud API.
type Client struct {
	vpcAPI *vpcv1.VpcV1
}

// NewClient initializes a client with a session.
func NewClient() (*Client, error) {
	apiKey := os.Getenv("IC_API_KEY")
	authenticator := &core.IamAuthenticator{
		ApiKey: apiKey,
	}

	vpcService, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
		Authenticator: authenticator,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		vpcAPI: vpcService,
	}, nil
}

func (c *Client) GetVPCCountForRegion(ctx context.Context, region string) (int64, error) {
	vpcs, _, err := c.vpcAPI.ListVpcsWithContext(ctx, &vpcv1.ListVpcsOptions{})
	if err != nil {
		return 0, err
	}

	totalCount := 0
	for _, vpc := range vpcs.Vpcs {

	}

	return *vpcs.TotalCount, nil
}
