package validatorapi

import (
	"context"
	"github.com/protolambda/eth2api"
	"github.com/protolambda/zrnt/eth2/beacon"
)

// Requests a beacon node to produce a valid block, which can then be signed by a validator.
func ProduceBlock(ctx context.Context, cli eth2api.Client,
	slot beacon.Slot, randaoReveal beacon.BLSSignature, graffiti *beacon.Root, dest *beacon.BeaconBlock) (syncing bool, err error) {
	q := eth2api.Query{
		"randao_reveal": randaoReveal,
	}
	if graffiti != nil {
		q["graffiti"] = graffiti
	}
	req := eth2api.FmtQueryGET(q, "eth/v1/validator/blocks/%d", slot)
	resp := cli.Request(ctx, req)
	if err := resp.Err(); err != nil {
		if err.Code() == 503 {
			return true, nil
		}
		return false, err
	}
	return false, resp.Decode(eth2api.Wrap(dest))
}
