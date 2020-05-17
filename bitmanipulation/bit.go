package bitmanipulation

type BitValue int

const (
	Falsy BitValue = iota
	Truthy
)

// GetBit returns the bit value at bitposition
// This method shifts the relevant bit to the zeroth position.
// Then we perform AND operation with one which has bit pattern like 0001.
// This clears all bits from the original number except the relevant one.
// If the relevant bit is one, the result is 1, otherwise the result is 0.
func GetBit(number, bitposition int) int {
	return (number >> bitposition) & 1
}

// SetBit sets the bit at bitposition
func SetBit(number, bitposition int) int {
	return number | (1 << bitposition)
}

// countSetBits,
// bitsDiff,
// bitLength

// ClearBit clears a bit if it is set
func ClearBit(number, bitposition int) int {
	mask := ^(1 << bitposition)
	return number & mask
}

func UpdateBit(number, bitposition int, bitValue BitValue) int {
	// clear bit
	mask := ^(1 << bitposition)
	number = number & mask
	return number | (int(bitValue) << bitposition)
}
