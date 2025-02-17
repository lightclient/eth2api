package nodeapi

import (
	"context"
	"github.com/protolambda/eth2api"
)

// Retrieves data about the node's network presence.
func Identity(ctx context.Context, cli eth2api.Client, dest *eth2api.NetworkIdentity) error {
	return eth2api.MinimalRequest(ctx, cli, eth2api.PlainGET("/eth/v1/node/identity"), eth2api.Wrap(dest))
}
