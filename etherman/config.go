package etherman

import "github.com/0xPolygonHermez/zkevm-ethtx-manager/etherman"

// Config represents the configuration of the etherman
type Config struct {
	EthermanConfig etherman.Config

	// IsValidiumMode is a flag to indicate if the sequence sender is running in validium mode
	IsValidiumMode bool
	// ForkIDChunkSize is the max interval for each call to L1 provider to get the forkIDs
	ForkIDChunkSize uint64 `mapstructure:"ForkIDChunkSize"`
}
