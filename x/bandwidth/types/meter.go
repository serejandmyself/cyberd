package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccStakeProvider interface {
	GetAccStakePercentage(ctx sdk.Context, address sdk.AccAddress) float64
}

// General bw handler tx flow
// 1. Update bw to current for signer
// 2. Check if signer have enough bw
// 3. Run Tx, if tx succeed, than:
// 4. Consume bw and save with old max bw value
// 5. Load bw with new max value and save it
type BandwidthMeter interface {
	// Update acc max bandwidth for current stake. Also, performs recover.
	UpdateAccMaxBandwidth(ctx sdk.Context, address sdk.AccAddress)
	// Returns recovered to current block height acc bandwidth
	GetCurrentAccBandwidth(ctx sdk.Context, address sdk.AccAddress) AcсBandwidth
	// Returns acc max bandwidth
	GetAccMaxBandwidth(ctx sdk.Context, address sdk.AccAddress) int64
	// Returns tx bandwidth cost
	GetTxCost(ctx sdk.Context, price float64, tx sdk.Tx) int64
	//
	// Performs bw consumption for given acc
	// To get right number, should be called after tx delivery with bw state obtained prior delivery
	//
	// Pseudo code:
	// bw := getCurrentBw(addr)
	// bwCost := deliverTx(tx)
	// consumeBw(bw, bwCost)
	ConsumeAccBandwidth(ctx sdk.Context, bw AcсBandwidth, amt int64)
}
