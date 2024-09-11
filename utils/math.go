package utils

import (
	"math"

	"github.com/mavryk-network/mvgo/mavryk"
)

type FloatConstraint interface {
	float32 | float64
}

func getZPortion[T FloatConstraint](val mavryk.Z, portion T) mavryk.Z {
	// percentage with 4 decimals
	portionFromVal := int64(math.Floor(float64(portion) * 10000))
	return val.Mul64(portionFromVal).Div64(10000)
}

func GetZPortion[T FloatConstraint](val mavryk.Z, portion T) mavryk.Z {
	if portion <= 0 {
		return mavryk.Zero
	}
	if portion >= 1 {
		return val
	}
	return getZPortion(val, portion)
}

func IsPortionWithin0n1[T FloatConstraint](portion T) bool {
	total := mavryk.NewZ(1000000)
	zPortion := getZPortion(total, portion)
	totalSubZPortion := total.Sub(zPortion)
	return !zPortion.IsNeg() && !totalSubZPortion.IsNeg()
}

type NumberConstraint interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Max[T NumberConstraint](v1 T, v2 T) T {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Abs[T NumberConstraint](v T) T {
	if v < 0 {
		return -v
	}
	return v
}
