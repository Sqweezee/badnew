package util

import (
	"fmt"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func ValidateDecimal(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid type: %T", i)
	}

	if v.IsNil() || v.IsNegative() {
		return fmt.Errorf("invalid decimal: %s", v)
	}

	return nil
}

func ValidateNumber(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid type: %T", i)
	}

	if math.NewIntFromUint64(v).IsNil() || math.NewIntFromUint64(v).IsNegative() {
		return fmt.Errorf("invalid number: %d", v)
	}

	return nil
}

func ValidatePositiveNumber(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid type: %T", i)
	}

	if math.NewIntFromUint64(v).IsNil() ||
		math.NewIntFromUint64(v).IsNegative() ||
		math.NewIntFromUint64(v).IsZero() {
		return fmt.Errorf("invalid number: %d", v)
	}

	return nil
}

func ValidatePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid type: %T", i)
	}

	if v.IsNil() || v.IsNegative() || v.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid percentage: %s", v)
	}

	return nil
}
