package fireblocks

import (
	"time"

	"github.com/omni-network/omni/lib/errors"
)

type Environment int

const (
	TestNet Environment = iota + 1 // EnumIndex = 1
	MainNet                        // EnumIndex = 2
)

// options houses parameters for altering the behavior of a SimpleTxManager.
type options struct {
	// NetworkTimeout is the allowed duration for a single network request.
	// This is intended to be used for network requests that can be replayed.
	NetworkTimeout time.Duration

	// QueryInterval is the interval at which the FireBlocks client will
	// call the get transaction by id to check for confirmations after a txn
	// has been sent
	QueryInterval time.Duration

	// LogFreqFactor is the frequency at which the FireBlocks client will
	// log a warning message if the transaction has not been signed yet
	LogFreqFactor int

	// Network is the environment that we have deployed in, either testnet or mainnet
	Network Environment

	// VaultAccountID is the ID of the vault account to use for signing.
	VaultAccountID uint64

	// HostOverride overrides the network based host if populated.
	HostOverride string
}

func (c options) host() string {
	if c.HostOverride != "" {
		return c.HostOverride
	}

	switch c.Network {
	case MainNet:
		return hostProd
	default:
		return hostSandbox
	}
}

// String - Creating common behavior - give the type a String function.
func (e Environment) String() string {
	return [...]string{"testnet", "mainnet"}[e-1]
}

// defaultOptions returns a options with default values.
func defaultOptions() options {
	return options{
		NetworkTimeout: time.Duration(30) * time.Second,
		QueryInterval:  time.Duration(500) * time.Millisecond,
		LogFreqFactor:  10,
		Network:        TestNet,
		VaultAccountID: 0,
	}
}

func WithQueryInterval(interval time.Duration) func(*options) {
	return func(cfg *options) {
		cfg.QueryInterval = interval
	}
}

func WithLogFreqFactor(factor int) func(*options) {
	return func(cfg *options) {
		cfg.LogFreqFactor = factor
	}
}

func WithHost(host string) func(*options) {
	return func(cfg *options) {
		cfg.HostOverride = host
	}
}

func WithEnvironment(env Environment) func(*options) {
	return func(cfg *options) {
		cfg.Network = env
	}
}

func WithVaultAccountID(id uint64) func(*options) {
	return func(cfg *options) {
		cfg.VaultAccountID = id
	}
}

// check validates the options.
func (c options) check() error {
	if c.LogFreqFactor <= 0 {
		return errors.New("must provide LogFreqFactor")
	}

	if c.NetworkTimeout <= 0 {
		return errors.New("must provide NetworkTimeout")
	}

	if c.QueryInterval <= 0 {
		return errors.New("must provide QueryInterval")
	}

	if c.Network != TestNet && c.Network != MainNet {
		return errors.New("must provide valid Network")
	}

	return nil
}
