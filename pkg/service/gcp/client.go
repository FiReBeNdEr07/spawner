package gcp

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	container "cloud.google.com/go/container/apiv1"
	"gitlab.com/netbook-devs/spawner-service/pkg/service/system"
	"google.golang.org/api/option"
)

func getClusterManagerClient(ctx context.Context, cred *system.GCPCredential) (*container.ClusterManagerClient, error) {

	sa_cred := []byte(cred.Certificate)
	opt := option.WithCredentialsJSON(sa_cred)
	c, err := container.NewClusterManagerClient(ctx, opt)
	return c, err
}

//getDiskClient
func getDiskClient(ctx context.Context, cred *system.GCPCredential) (*compute.DisksClient, error) {

	sa_cred := []byte(cred.Certificate)
	opt := option.WithCredentialsJSON(sa_cred)
	return compute.NewDisksRESTClient(ctx, opt)
}
