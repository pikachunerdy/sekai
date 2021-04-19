package types

import (
	"fmt"
	"time"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default parameter namespace
const (
	DefaultSignedBlocksWindow       = int64(100)
	DefaultMaxMischance             = int64(110)
	DefaultDowntimeInactiveDuration = 60 * 10 * time.Second
	DefaultMischanceConfidence      = int64(10)
)

// Parameter store keys
var (
	KeySignedBlocksWindow       = []byte("SignedBlocksWindow")
	KeyMaxMischance             = []byte("MaxMischance")
	KeyDowntimeInactiveDuration = []byte("DowntimeInactiveDuration")
	KeyMischanceConfidence      = []byte("MischanceConfidence")
)

// ParamKeyTable for slashing module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params object
func NewParams(
	signedBlocksWindow int64, maxMischance int64, downtimeInactiveDuration time.Duration, mischanceConfidence int64,
) Params {

	return Params{
		SignedBlocksWindow:       signedBlocksWindow,
		MaxMischance:             maxMischance,
		DowntimeInactiveDuration: downtimeInactiveDuration,
		MischanceConfidence:      mischanceConfidence,
	}
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeySignedBlocksWindow, &p.SignedBlocksWindow, validateSignedBlocksWindow),
		paramtypes.NewParamSetPair(KeyMaxMischance, &p.MaxMischance, validateMaxMischance),
		paramtypes.NewParamSetPair(KeyDowntimeInactiveDuration, &p.DowntimeInactiveDuration, validateDowntimeInactiveDuration),
		paramtypes.NewParamSetPair(KeyMischanceConfidence, &p.MischanceConfidence, validateMischanceConfidence),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultSignedBlocksWindow, DefaultMaxMischance, DefaultDowntimeInactiveDuration, DefaultMischanceConfidence,
	)
}

func validateSignedBlocksWindow(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid signed_blocks_window parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("signed blocks window must be positive: %d", v)
	}

	return nil
}

func validateMaxMischance(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid max_mischance parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("max mischance cannot be negative or equal to zero: %d", v)
	}

	return nil
}

func validateDowntimeInactiveDuration(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid downtime_inactive_duration parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("downtime inactive duration must be positive: %s", v)
	}

	return nil
}

func validateMischanceConfidence(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid mischance_confidence parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("mischance confidence cannot be negative: %d", v)
	}

	return nil
}
